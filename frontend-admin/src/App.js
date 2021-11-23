import React from 'react';
import './App.css';
import {  Redirect, Route, Switch } from "react-router-dom";
import AllItem from './views/admin/AllItem.js';
import AllTypes from './views/admin/AllTypes.js';
import AdminNavbar from './components/navbar/AdminNavbar.js';
import AdminSidebar from './components/sidebar/AdminSidebar.js';
export default function App() {
  return (
    <>
      <AdminSidebar />
      <div className="relative md:ml-64 bg-blueGray-100">
        <AdminNavbar />
        <div className="px-4 md:px-10 mx-auto w-full">
          <Switch>
            {/* add routes with layouts */}
            <Route path="/admin/AllItem" exact component={AllItem} />
            <Route path="admin/AllTypes" exact  component={AllTypes} />
            {/* <Route path="/Confirm" exact component={Confirm} />
            <Route path="/List" exact component={List} />
            <Route path="/ListCustomer" exact component={ListCustomer} />
            <Route path="/ReplyReviews" exact component={ReplyReviews} />
            <Route path="/Punish" exact component={Punish} /> */}
            {/* add redirect for first page */}
            <Redirect from="/admin" to="/admin/AllItem" />
          </Switch>
        </div>
      </div>
      </>

  );
}