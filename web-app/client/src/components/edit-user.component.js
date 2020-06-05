import React, { Component } from "react";
import axios from "axios";

export class EditUser extends Component {
  constructor(props) {
    super(props);

    this.onChangeName = this.onChangeName.bind(this);
    this.onChangeEmail = this.onChangeEmail.bind(this);
    this.onChangeUsertype = this.onChangeUsertype.bind(this);
    this.onChangeAddress = this.onChangeAddress.bind(this);
    this.onSubmit = this.onSubmit.bind(this);

    this.state = {
      name: "",
      email: "",
      usertype: "",
      address: "",
    };
  }

  componentDidMount() {
    axios
      .get("http://localhost:5000/users/" + this.props.match.params.id)
      .then((response) => {
        this.setState({
          name: response.data.name,
          email: response.data.email,
          usertype: response.data.usertype,
          address: response.data.address,
        });
      });
  }

  onChangeName(e) {
    this.setState({
      name: e.target.value,
    });
  }

  onChangeEmail(e) {
    this.setState({
      email: e.target.value,
    });
  }

  onChangeUsertype(e) {
    this.setState({
      usertype: e.target.value,
    });
  }

  onChangeAddress(e) {
    this.setState({
      address: e.target.value,
    });
  }

  onSubmit(e) {
    e.preventDefault();

    const user = {
      name: this.state.name,
      email: this.state.email,
      usertype: this.state.usertype,
      address: this.state.address,
    };
    console.log(this.props.match);
    console.log(user);

    axios
      .post(
        "http://localhost:5000/users/update/" + this.props.match.params.id,
        user
      )
      .then((res) => console.log(res.data));

    window.location = "/users";
  }

  render() {
    return (
      <div>
        <h3>Edit User</h3>
        <form onSubmit={this.onSubmit}>
          <div className="form-group">
            <label>Name: </label>
            <input
              type="text"
              className="form-control"
              value={this.state.name}
              onChange={this.onChangeName}
            />
          </div>
          <div className="form-group">
            <label>Email: </label>
            <input
              type="text"
              className="form-control"
              value={this.state.email}
              onChange={this.onChangeEmail}
            />
          </div>
          <div className="form-group">
            <label>Usertype: </label>
            <select
              ref="usertypeInput"
              required
              className="form-control"
              value={this.state.usertype}
              onChange={this.onChangeUsertype}
            >
              <option key="Manufacturer" value="Manufacturer">
                Manufacturer
              </option>
              <option key="Distributor" value="Distributor">
                Distributor
              </option>
              <option key="Wholesaler" value="Wholesaler">
                Wholesaler
              </option>
              <option key="Retailer" value="Retailer">
                Retailer
              </option>
              <option key="Consumer" value="Consumer">
                Consumer
              </option>
            </select>
          </div>
          <div className="form-group">
            <label>Address: </label>
            <input
              type="text"
              className="form-control"
              value={this.state.address}
              onChange={this.onChangeAddress}
            />
          </div>
          <div className="form-group">
            <input
              type="submit"
              value="Update User"
              className="btn btn-primary"
            />
          </div>
        </form>
      </div>
    );
  }
}

export default EditUser;
