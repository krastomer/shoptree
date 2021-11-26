import React from 'react';
import './App.css';
import {  BrowserRouter, Route, Switch } from "react-router-dom";

import Home from './views/admin/Home';
import Product from './views/admin/Product';
import Category from './views/admin/Category';
import Confirm from './views/admin/Confirm';
import Delivery from './views/admin/Delivery';
import Customer from './views/admin/Customer';

export default function App() {
  return (
    <>
      <div className="relative md:ml-64 bg-blueGray-100 font-body">
        <div className="w-full px-4 mx-auto md:px-10">
          <BrowserRouter>
            <Switch>
              <Route path="/" exact>
                <Home />
              </Route>
              <Route path="/product" exact>
                <Product />
              </Route>
              <Route path="/category" exact>
                <Category />
              </Route>
              <Route path="/confirm" exact>
                <Confirm />
              </Route>
              <Route path="/delivery" exact>
                <Delivery />
              </Route>
              <Route path="/customer" exact>
                <Customer />
              </Route>
            </Switch>
          </BrowserRouter>
        </div>
      </div>
      </>

  );
}