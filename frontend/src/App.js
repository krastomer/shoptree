import React from "react";
import { BrowserRouter, Route, Switch } from "react-router-dom";
import Login from "./pages/login/Login";
import Register from "./pages/register/Register";

export default function Rounting() {
  return (
    <div className="wrap">
      <BrowserRouter>
        <Switch>
          <Route path="/login">
            <Login />
          </Route>
          <Route path="/register">
            <Register />
          </Route>
        </Switch>
      </BrowserRouter>
    </div>
  );
}
