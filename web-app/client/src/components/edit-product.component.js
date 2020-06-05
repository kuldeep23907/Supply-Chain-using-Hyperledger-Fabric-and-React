import React, { Component } from "react";
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";
import axios from "axios";

export class EditProduct extends Component {
  constructor(props) {
    super(props);

    this.state = {
      product_name: "",
      date: {
        manufacturerDate: "",
        sendToWholesalerDate: "",
        sendToDistributorDate: "",
        sendToRetailerDate: "",
        sellToConsumerDate: "",
        orderedDate: "",
        deliveredDate: "",
      },
      manufacturer_id: "",
      distributor_id: "",
      wholesaler_id: "",
      consumer_id: "",
      retailer_id: "",
      status: "",
      price: 0,
      manufacturers: [],
    };
  }

  componentDidMount() {
    axios.get("http://localhost:products/" + this.props.match.params.id)
    .then((response) => {
      this.setState({
        product_name: response.data.product_name,
        manufacturer_id: response.data.manufacturer_id,
        manufacturerDate: response.data.manufacturerDate,
        status: response.data.status,
        price: response.data.price,
      })
    })
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

  onChangeManufacturerId(e) {
    this.setState({
      manufacturer_id: e.target.value,
    });
  }

  onChangeManufacturerDate(date) {
    const newDate = { ...this.state.date, manufacturerDate: date };
    this.setState({ date: newDate });
  }

  render() {
    return (
      <div>
        <h3>Update Product</h3>
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
            <label>ManufacturerID: </label>
            <select
              ref="manufacturerInput"
              required
              className="form-control"
              value={this.state.manufacturer_id}
              onChange={this.onChangeManufacturerId}
            >
              {this.state.manufacturers.map(function (manufacturer) {
                return (
                  <option
                    key={manufacturer.user_id}
                    value={manufacturer.user_id}
                  >
                    {manufacturer.user_id}
                  </option>
                );
              })}
            </select>
          </div>
          <div className="form-group">
            <label>Manufacturer Date: </label>
            <div>
              <DatePicker
                selected={this.state.date.manufacturerDate}
                onChange={this.onChangeManufacturerDate}
              />
            </div>
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

export default EditProduct;
