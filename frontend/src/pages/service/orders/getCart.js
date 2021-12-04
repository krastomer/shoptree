import axios from "axios";



export const getCart = async () => {
    let response;
    console.log("token :" ,localStorage.getItem("token"));
    const config = {
            method: 'get',
            url: 'http://spaceship.trueddns.com:23720/api/v1/orders',
            headers: { 
                'Cookie' : `jwt=${localStorage.getItem("token")}`
            }
    }
    try {
        response = await axios.get('http://spaceship.trueddns.com:23720/api/v1/orders', {headers: { 
            'Cookie' : `jwt=${localStorage.getItem("token")}`}
        })
        console.log(response.data);
    } catch (error) {
        console.error(error)
    }
    return response?.data ? response?.data : null // or set initial value
}