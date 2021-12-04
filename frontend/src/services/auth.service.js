import axios from "axios";
import jwt from 'jwt-decode' 
const API_URL = "http://spaceship.trueddns.com:23720";
const register_ = (data) => {
  console.log(data)
  // return axios.post(API_URL + "signup", {
  //   username,
  //   email,
  //   password,
  // });
};

const login = (username, password) => {
  return axios
    .post(API_URL + "/api/v1/auth/login", {
      username,
      password,
    })
    .then((response) => {
      if (response.data.token) {
        const token =response.data.token;
        const user = jwt(token); // decode your token here
        localStorage.setItem("token", token);
        localStorage.setItem("user", JSON.stringify(user));
      }

      return response.data;
    });
};
const logout = () => {
  localStorage.removeItem("token");
  localStorage.removeItem("user");
};

export default {
  register_,
  login,
  logout,
};
