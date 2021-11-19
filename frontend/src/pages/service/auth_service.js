import React from "react";
import { Api_Url } from "../constant/Link";
import { VeryfyToken } from "./verifytoken";

export function postLogin(User) {
  var myHeaders = new Headers();
  // myHeaders.append(
  //   "Cookie",
  //   "jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEsInVzZXIiOiJrcmFzdG9tZXJAZ21haWwuY29tIiwiYXVkIjoiQ3VzdG9tZXIiLCJleHAiOjE2MzU4MjgwODcsImlzcyI6InNob3B0cmVlIn0.rOIblmfOV0wAZtf7xKf6MpH8Od2fq4mkf_1miG45QXQ"
  // );
  var formdata = new FormData();
  formdata.append("Username", User.username);
  formdata.append("Password", User.password);
  var requestOptions = {
    method: "POST",
    body: formdata,
    redirect: "follow",
  };
  fetch(`${Api_Url}/api/v1/auth/login`, requestOptions)
    .then((response) => response.json())
    .then((data) => {
      User.auth.token = data.token;
      localStorage.setItem("user", JSON.stringify(VeryfyToken(data.token)));
    })
    .catch((error) => console.log("sasdas"));
  User.auth.loggedIn = true;
}

export function postRegister(data) {
  var axios = require("axios");
  var data = JSON.stringify({
    username: data.email,
    name: data.name,
    password: data.password,
    phone_number: data.mobile_number,
  });
  console.log(data)
  var config = {
    method: "post",
    url: `${Api_Url}/api/v1/auth/register`,
    headers: {
      "Content-Type": "application/json",
    },
    data: data,
  };

  axios(config)
    .then(function (response) {
      console.log(JSON.stringify(response.data));
    })
    .catch(function (error) {
      console.log(error.response.data);
    });
}
