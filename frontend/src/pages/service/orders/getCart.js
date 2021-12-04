import axios from "axios";
export const getCart = async () => {
    let response;

    const config = {
            method: 'get',
            url: 'http://spaceship.trueddns.com:23720/api/v1/orders',
            headers: { 
                'Authorization': `Bearer ${localStorage.getItem("token")}`,
            }
    }
    try {
        response = await axios(config)
        //console.log(response.data);
    } catch (error) {
        console.error(error)
    }
    return response?.data ? response?.data : null // or set initial value
}