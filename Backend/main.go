package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	Routers()
}

func Routers() {
	InitDB()
	defer db.Close()
	log.Println("Starting the HTTP server on port 8001")
	router := mux.NewRouter()
	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.HandleFunc("/users",CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}",GetUser).Methods("GET")
	router.HandleFunc("/users/{id}",UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}",DeleteUser).Methods("DELETE")
	
	http.ListenAndServe(":8001", &CORSRouterDecorator{R: router})
}

/***************************************************/

// Get all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []User

	result, err := db.Query("SELECT id, judul_buku, jumlah, nama_peminjam, alamat_peminjam, nohp_peminjam, tanggal_pinjam, tanggal_pengembalian, lama_pinjam FROM peminjamanbuku")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		var user User
		err := result.Scan(&user.ID, &user.JudulBuku, &user.Jumlah, &user.NamaPeminjam, &user.AlamatPeminjam, &user.NohpPeminjam, &user.TanggalPinjam, &user.TanggalPengembalian, &user.LamaPinjam)
		if err != nil {
			panic(err.Error())
		}
		users = append(users, user)
	}
	json.NewEncoder(w).Encode(users)
}

//Create user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	stmt, err := db.Prepare("INSERT INTO peminjamanbuku(judul_buku," +
		"jumlah,nama_peminjam,alamat_peminjam,nohp_peminjam,tanggal_pinjam,tanggal_pengembalian,lama_pinjam) VALUES(?,?,?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	judul_buku := keyVal["judul_buku"]
	jumlah := keyVal["jumlah"]
	nama_peminjam := keyVal["nama_peminjam"]
	alamat_peminjam := keyVal["alamat_peminjam"]
	nohp_peminjam := keyVal["nohp_peminjam"]
	tanggal_pinjam := keyVal["tanggal_pinjam"]
	tanggal_pengembalian := keyVal["tanggal_pengembalian"]
	lama_pinjam := keyVal["lama_pinjam"]
	_, err = stmt.Exec(judul_buku,jumlah,nama_peminjam,alamat_peminjam,nohp_peminjam,tanggal_pinjam,tanggal_pengembalian,lama_pinjam)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "New user was created")
}

//Get user by ID
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := db.Query("SELECT id, judul_buku, jumlah, nama_peminjam, alamat_peminjam, nohp_peminjam, tanggal_pinjam, tanggal_pengembalian, lama_pinjam FROM peminjamanbuku WHERE id = ?", params["id"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var user User
	for result.Next() {
		err := result.Scan(&user.ID, &user.JudulBuku, &user.Jumlah, &user.NamaPeminjam, &user.AlamatPeminjam, &user.NohpPeminjam, &user.TanggalPinjam, &user.TanggalPengembalian, &user.LamaPinjam)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(user)
}

//Update user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	stmt, err := db.Prepare("UPDATE peminjamanbuku SET judul_buku = ?," + 
		"jumlah= ?, nama_peminjam=?, alamat_peminjam=?, nohp_peminjam=?, tanggal_pinjam=?, tanggal_pengembalian=?, lama_pinjam=? WHERE id =?")

	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	judul_buku := keyVal["judul_buku"]
	jumlah := keyVal["jumlah"]
	nama_peminjam := keyVal["nama_peminjam"]
	alamat_peminjam := keyVal["alamat_peminjam"]
	nohp_peminjam := keyVal["nohp_peminjam"]
	tanggal_pinjam := keyVal["tanggal_pinjam"]
	tanggal_pengembalian := keyVal["tanggal_pengembalian"]
	lama_pinjam := keyVal["lama_pinjam"]
	_, err = stmt.Exec(judul_buku,jumlah,nama_peminjam,alamat_peminjam,nohp_peminjam,tanggal_pinjam,tanggal_pengembalian,lama_pinjam, params["id"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "User with ID = %s was updated",
		params["id"])
}


func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	stmt, err := db.Prepare("DELETE FROM peminjamanbuku WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(params["id"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "User with ID = %s was deleted",
		params["id"])
}


type User struct {
	ID                 int `json:"id"`
	JudulBuku          string `json:"judul_buku"`
	Jumlah             string `json:"jumlah"`
	NamaPeminjam       string `json:"nama_peminjam"`
	AlamatPeminjam     string `json:"alamat_peminjam"`
	NohpPeminjam       string `json:"nohp_peminjam"`
	TanggalPinjam      string `json:"tanggal_pinjam"`
	TanggalPengembalian string `json:"tanggal_pengembalian"`
	LamaPinjam         string `json:"lama_pinjam"`
}

var db *sql.DB
var err error

func InitDB() {
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_2207989_mohammadsyafaaulia_uas_pilkomb")
	if err != nil {
		panic(err.Error())
	}
}


type CORSRouterDecorator struct {
	R *mux.Router
}

func (c *CORSRouterDecorator) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods","POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Language, Content-Type, YourOwnHeader")
	}

	if req.Method == "OPTIONS" {
		return
	}

	c.R.ServeHTTP(rw, req)
}
