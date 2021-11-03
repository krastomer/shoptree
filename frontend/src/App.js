import React from "react";
import { BrowserRouter, Route, Switch } from "react-router-dom";
import Login from "./pages/login/Login";
import Register from "./pages/register/Register";
import Home from "./pages/home/Home";
import Review from "./pages/review/Review";
import Order from "./pages/order/Order";
import Navbar from "./asset/include/navbar/Navbar";
import { CookiesProvider } from "react-cookie";

export default function App() {
  return (
    <div className="wrap">
      <CookiesProvider>
        <BrowserRouter>
          <Switch>
            <Route path="/" exact>
              <Home />
            </Route>
            <Route path="/order" exact>
              <Order />
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
            <Route path="/navbar" exact>
              <Navbar />
            </Route>
          </Switch>
        </BrowserRouter>
      </CookiesProvider>
    </div>
  );
}
