import React from "react";
import { BrowserRouter, Route, Switch } from "react-router-dom";
import Login from "./pages/login/Login";
import Register from "./pages/register/Register";
import Home from "./pages/home/Home";
import Review from "./pages/review/Review";
import Modal from "./pages/review/Modal";
import Order from "./pages/order/Order";
import Navbar from "./asset/include/navbar/Navbar";
import { CookiesProvider } from "react-cookie";
import Detail from "./pages/home/Detailproduct";
import allProduct from "./pages/home/allProduct";
import Dropdown from "./pages/review/Dropdown";
import StarRating from "./pages/review/StarRating";
import { LoginUser } from "./models/User";
import { SuccessOrder } from "./pages/success/success";
import Profile from "./pages/profile/Profile";
export default function App() {
  return (
    <div className="wrap">
      <CookiesProvider>
        <BrowserRouter>
          <Switch>
            <Route path="/" exact>
              <Home />
            </Route>
            <Route path="/products/:id" exact>
              <Detail id = {allProduct}></Detail>
            </Route>
            <Route path="/order/:USER" exact>
              <Order USER={LoginUser.username}/>
            </Route>
            <Route path="/login" exact>
              <Login />
            </Route>
            <Route path="/register" exact>
              <Register />
            </Route>
            <Route path="/review" exact>
              <Review />
            </Route>
            <Route path="/profile" exact>
              <Profile />
            </Route>
            <Route path="/navbar" exact>
              <Navbar />
            </Route>
            <Route path="/modal" exact>
              <Modal />
            </Route>
            <Route path="/dropdown" exact>
              <Dropdown />
            </Route>
            <Route path="/starrating" exact>
              <StarRating />
            </Route>
            <Route path="/order/:USER/Success" exact>
              <SuccessOrder USER = {LoginUser.username}></SuccessOrder>
            </Route>
          </Switch>
        </BrowserRouter>
      </CookiesProvider>
    </div>
  );
}
