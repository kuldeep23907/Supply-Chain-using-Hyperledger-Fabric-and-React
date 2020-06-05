import React from "react";
import { BrowserRouter as Router, Route } from "react-router-dom";
import "bootstrap/dist/css/bootstrap.min.css";
import Navbar from "./components/navbar.component";
import SignIn from "./components/signIn.component";
import CreateUser from "./components/create-user.component";
import CreateProduct from "./components/create-product.component";
import CreateOrder from "./components/create-order.component";
import EditUser from "./components/edit-user.component";
import EditProduct from "./components/edit-product.component";
import UsersList from "./components/users-list.component";
import ProductsList from "./components/products-list.component";
import OrdersList from "./components/orders-list.component";

function App() {
  const role = sessionStorage.getItem("role");
  console.log(role)
  return (
    <Router>
      <div className="container">
        <Navbar />
        <br />
        <Route path="/" exact component={SignIn} />
        <Route path="/products" component={ProductsList} />
        <Route path="/createUser" component={CreateUser} />
        <Route path="/createProduct" component={CreateProduct} />
        <Route path="/createOrder" component={CreateOrder} />
        <Route path="/updateUser/:id" component={EditUser} />
        <Route path="/updateProduct/:id" component={EditProduct} />
        <Route path="/users" component={UsersList} />
        <Route path="/orders" component={OrdersList} />
      </div>
    </Router>
  );
}

export default App;
