import axios from "axios";

const accesstoken = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEsInVzZXIiOiJrcmFzdG9tZXJAZ21haWwuY29tIiwiYXVkIjoiQ3VzdG9tZXIiLCJleHAiOjE2Mzg2MjA5MzQsImlzcyI6InNob3B0cmVlIn0.dzbBuHH0TkUzhe8goPdAkl9rXwuyLoB6HzAK4wDVMV8';
const appURL = 'http://spaceship.trueddns.com:23720/api';
const headers = {
    'Content-Type': 'application/json',
    'Authorization': accesstoken,
    'Cookie': accesstoken
}
  
export const getCart = async () => {
    let response;
    try {
        response = await axios.get(`${appURL}/v1/orders`,{
            headers: headers
        })
        console.log(response.data);
    } catch (error) {
        console.error(error)
    }
    return response?.data ? response?.data : null // or set initial value
}