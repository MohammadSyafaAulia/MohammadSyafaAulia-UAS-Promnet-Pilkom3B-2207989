import React, { Component } from "react";
import UserService from "../services/UserService";
import Swal from "sweetalert2";

class ListUserComponent extends Component {
  constructor(props) {
    super(props);

    this.state = {
      users: [],
      searchQuery: "",
    };

    this.addUser = this.addUser.bind(this);
    this.editUser = this.editUser.bind(this);
    this.deleteUser = this.deleteUser.bind(this);
    this.viewUser = this.viewUser.bind(this);
    this.handleSearchChange = this.handleSearchChange.bind(this);
  }

  showSuccessMessage = () => {
    Swal.fire({
      title: "Are you sure?",
      text: "You won't be able to revert this!",
      icon: "warning",
      showCancelButton: true,
      confirmButtonColor: "#3085d6",
      cancelButtonColor: "#d33",
      confirmButtonText: "Yes, delete it!",
    }).then((result) => {
      if (result.isConfirmed) {
        Swal.fire({
          title: "Deleted!",
          text: "Your file has been deleted.",
          icon: "success",
        });
      }
    });
  };

  showSuccessMessage2 = () => {
    Swal.fire({
      position: "center",
      icon: "success",
      title: "Your work has been saved",
      showConfirmButton: false,
      timer: 1500,
    });
  };

  deleteUser(id) {
    Swal.fire({
      title: "Are you sure?",
      text: "You won't be able to revert this!",
      icon: "warning",
      showCancelButton: true,
      confirmButtonColor: "#3085d6",
      cancelButtonColor: "#d33",
      confirmButtonText: "Yes, delete it!",
    }).then((result) => {
      if (result.isConfirmed) {
        UserService.deleteUser(id).then((res) => {
          this.setState({
            users: this.state.users.filter((user) => user.id !== id),
          });

          Swal.fire({
            title: "Deleted!",
            text: "Your file has been deleted.",
            icon: "success",
          });
        });
      }
    });
  }

  viewUser(id) {
    this.props.history.push(`/view-user/${id}`);
  }

  editUser(id) {
    this.props.history.push(`/add-user/${id}`);
  }

  handleSearchChange(event) {
    this.setState({ searchQuery: event.target.value });
  }

  componentDidMount() {
    this.fetchUsers();
  }

  fetchUsers() {
    UserService.getUsers().then((res) => {
      if (res.data == null) {
        this.props.history.push("/add-user/_add");
      }
      this.setState({ users: res.data });
    });
  }

  addUser() {
    this.props.history.push("/add-user/_add");
  }

  render() {
    const filteredUsers = this.state.users.filter((user) =>
      Object.values(user).some((value) =>
        value
          .toString()
          .toLowerCase()
          .includes(this.state.searchQuery.toLowerCase())
      )
    );

    return (
      <div>
        <h2 className="text-center">Data Peminjaman Buku</h2>
        <br></br>
        <br></br>
        <div className="row">
          <div className="col-md-9">
            <button className="btn btn-primary" onClick={this.addUser}>
              Tambahkan Data
            </button>
          </div>
          <div className="col-md-3">
            <div className="mb-3">
              <input
                type="text"
                className="form-control"
                id="searchInput"
                placeholder="Search users..."
                value={this.state.searchQuery}
                onChange={this.handleSearchChange}
              />
            </div>
          </div>
        </div>
        <div className="row">
          <table className="table table-striped">
            <thead>
              <tr style={{ margin: "0", border: "none" }}>
                <th style={{ border: "none" }}>Judul Buku</th>
                <th style={{ border: "none" }}>Jumlah Buku</th>
                <th style={{ border: "none" }}>Nama Peminjam</th>
                <th style={{ border: "none" }}>Alamat Peminjam</th>
                <th style={{ border: "none" }}>No Hp Peminjam</th>
                <th style={{ border: "none", width: "10%" }}>Tanggal Pinjam</th>
                <th style={{ border: "none" }}>Tanggal Pengembalian</th>
                <th style={{ border: "none" }}>Lama Pinjam</th>
                <th
                  style={{ border: "none", justifyContent: "center" }}
                  className="d-flex justify-content-center"
                >
                  Actions
                </th>
              </tr>
            </thead>
            <tbody>
              {filteredUsers.map((user) => (
                <tr key={user.id}>
                  <td>{user.judul_buku}</td>
                  <td>{user.jumlah}</td>
                  <td>{user.nama_peminjam}</td>
                  <td>{user.alamat_peminjam}</td>
                  <td>{user.nohp_peminjam}</td>
                  <td>{user.tanggal_pinjam}</td>
                  <td>{user.tanggal_pengembalian}</td>
                  <td>{user.lama_pinjam}</td>
                  <td>
                    <div
                      style={{
                        display: "flex",
                        justifyContent: "space-between",
                      }}
                    >
                      <button
                        style={{ flex: 1, width: "30%", marginRight: "5px" }}
                        onClick={() => this.editUser(user.id)}
                        className="btn btn-info"
                      >
                        Ubah
                      </button>

                      <button
                        style={{
                          flex: 1,
                          width: "30%",
                          padding: "10px",
                          marginRight: "5px",
                        }}
                        onClick={() => this.deleteUser(user.id)}
                        className="btn btn-danger"
                      >
                        Hapus
                      </button>

                      <button
                        style={{ flex: 1, width: "30%", marginLeft: "5px" }}
                        onClick={() => this.viewUser(user.id)}
                        className="btn btn-info"
                      >
                        Lihat
                      </button>
                    </div>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </div>
    );
  }
}

export default ListUserComponent;
