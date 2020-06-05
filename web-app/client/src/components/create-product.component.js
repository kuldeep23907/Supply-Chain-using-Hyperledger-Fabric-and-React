import React, { Component } from "react";
import axios from "axios";

export class CreateProduct extends Component {
  constructor(props) {
    super(props);

    this.onChangeProductName = this.onChangeProductName.bind(this);
    this.onChangePrice = this.onChangePrice.bind(this);
    this.onSubmit = this.onSubmit.bind(this);

    this.state = {
      product_name: "",
      price: 0,
    };
  }

  onChangeProductName(e) {
    this.setState({
      product_name: e.target.value,
    });
  }

  onChangePrice(e) {
    this.setState({
      price: e.target.value,
    });
  }

  onSubmit(e) {
    e.preventDefault();

    const product = {
      name: this.state.product_name,
      price: this.state.price,
    };

    const headers = {
      "x-access-token": sessionStorage.getItem("jwtToken"),
    };

    console.log(product);

    axios
      .post("http://192.168.0.108:8090/product", product, { headers: headers })
      .then((res) => console.log(res));

    window.location = "/products";
  }

  render() {
    return (
      <div>
        <h3>Create New Product</h3>
        <form onSubmit={this.onSubmit}>
          <div className="form-group">
            <label>ProductName: </label>
            <input
              type="text"
              required
              className="form-control"
              value={this.state.product_name}
              onChange={this.onChangeProductName}
            />
          </div>
          <div className="form-group">
            <label>Price: </label>
            <input
              type="number"
              required
              className="form-control"
              value={this.state.price}
              onChange={this.onChangePrice}
            />
          </div>
          <div className="form-group">
            <input
              type="submit"
              value="Create Product"
              className="btn btn-primary"
            />
          </div>
        </form>
      </div>
    );
  }
}

export default CreateProduct;
