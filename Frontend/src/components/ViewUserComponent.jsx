import React, { Component } from "react";
import UserService from "../services/UserService";

class ViewUserComponent extends Component {
  constructor(props) {
    super(props);

    this.state = {
      id: this.props.match.params.id,
      user: {},
    };
  }

  componentDidMount() {
    UserService.getUserById(this.state.id).then((res) => {
      this.setState({ user: res.data });
    });
  }

  capitalizeFirstLetter = (str) => {
    return str.charAt(0).toUpperCase() + str.slice(1);
  };

  render() {
    return (
      <div>
        <br />
        <div className="card col-md-8 offset-md-2">
          <h3 className="text-center">View User Details</h3>
          <div className="card-body">
            <div className="table-responsive">
              <table className="table table-bordered">
                <tbody>
                  {Object.entries(this.state.user).map(
                    ([key, value], index) => (
                      <tr
                        key={index}
                        className={index % 2 === 0 ? "even-row" : "odd-row"}
                      >
                        <td>
                          <strong>
                            {this.capitalizeFirstLetter(key.replace(/_/g, " "))}
                          </strong>
                        </td>
                        <td className="text-center">{value}</td>
                      </tr>
                    )
                  )}
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </div>
    );
  }
}

export default ViewUserComponent;
