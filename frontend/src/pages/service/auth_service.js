import React from "react";
import { Api_Url } from "../constant/Link";
export function getProduct(item) {
  var requestOptions = {
    method: "GET",
    redirect: "follow",
  };

  fetch(`${Api_Url}/api/v1/products/1`, requestOptions)
    .then((response) => response.text())
    .then((result) => console.log(result))
    .catch((error) => console.log("error", error));
}
export function postLogin(User) {
  var myHeaders = new Headers();
  myHeaders.append(
    "Cookie",
    "jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEsInVzZXIiOiJrcmFzdG9tZXJAZ21haWwuY29tIiwiYXVkIjoiQ3VzdG9tZXIiLCJleHAiOjE2MzU4MjgwODcsImlzcyI6InNob3B0cmVlIn0.rOIblmfOV0wAZtf7xKf6MpH8Od2fq4mkf_1miG45QXQ"
  );
  var formdata = new FormData();
  formdata.append("Username", User.username);
  formdata.append("Password", User.password);
  var requestOptions = {
    method: "POST",
    headers: myHeaders,
    body: formdata,
    redirect: "follow",
  };
  const token = fetch(`${Api_Url}/api/v1/auth/login`, requestOptions)
  .then(response => 
    response.json())
  .then(data => {
    User.auth.token = data.token
  })
  .catch(error=> console.log(error)); 
  User.auth.loggedIn=true
}

export function postRegister(item) {
  var myHeaders = new Headers();
  myHeaders.append("Content-Type", "application/json");
  var raw = JSON.stringify({
    username: item[1],
    name: item[0],
    password: item[2],
    phone_number: item[3],
    level: "Customer",
  });

  var requestOptions = {
    method: "POST",
    headers: myHeaders,
    body: raw,
    redirect: "follow",
  };

  fetch(`${Api_Url}/api/v1/customers/`, requestOptions)
    .then((response) => response.text())
    .then((result) => console.log(result))
    .catch((error) => console.log("error", error));
}
