import React from "react";
import { BrowserRouter, Route, Switch } from "react-router-dom";
import Login from "./pages/login/Login";
import Register from "./pages/register/Register";
import Home from "./pages/Home/Home";
import Review from "./pages/review/Review";
import Navbar from "./asset/include/navbar/Navbar"

export default function App() {
  return (
    <div className="wrap">
      <BrowserRouter>
        <Switch>
          <Route path="/home">
            <Home />
          </Route>
          <Route path="/login">
            <Login />
          </Route>
          <Route path="/register">
            <Register />
          </Route>
          <Route path="/review">
            <Review />
          </Route>
          <Route path="/navbar">
            <Navbar />
          </Route>
        </Switch>
      </BrowserRouter>
    </div>
  );
}
