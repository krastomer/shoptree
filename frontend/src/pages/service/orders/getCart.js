import axios from "axios";

const accesstoken = localStorage.getItem("token")
const appURL = 'http://spaceship.trueddns.com:23720/api';
const headers = {
    'Cookie': `jwt=${accesstoken}`
}
export const login = (username, password) => {
    return axios
      .get(appURL + "/v1/orders",
      )
      .then((response) => {
        console.log("kuy")
  
        return response.data;
      });
  };
export const getCart = async () => {
    let response;
    try {
        response = await axios.get(`${appURL}/v1/orders`,{
        headers: headers
        })
    } catch (error) {
        console.error(error)
    }
    return response?.data ? response?.data : null // or set initial value
}