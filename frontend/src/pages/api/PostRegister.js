import React from "react";
export function getProduct(item){
    var requestOptions = {
        method: 'GET',
        redirect: 'follow'
      };
      
      fetch("http://spaceship.trueddns.com:23720/api/v1/products/1", requestOptions)
        .then(response => response.text())
        .then(result => console.log(result))
        .catch(error => console.log('error', error));
}
export function postRegister(item){
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

    fetch("http://spaceship.trueddns.com:23720/api/v1/auth/register", requestOptions)
      .then((response) => response.text())
      .then((result) => console.log(result))
      .catch((error) => console.log("error", error));
}