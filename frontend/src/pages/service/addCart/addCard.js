import axios from "axios";

export const addItemByID = async (id) => {
    let response;
    const config = {
            method:'post',
            url: `http://spaceship.trueddns.com:23720/api/v1/orders/products/${id}`,
            headers: { 
                'Authorization': `Bearer ${localStorage.getItem("token")}`,
            }
    }
    try {
        response = await axios(config)
    } catch (error) {
        console.error(error)
    }
    return response?.data ? response?.data : null // or set initial value
}