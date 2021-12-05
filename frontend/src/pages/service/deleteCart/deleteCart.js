import axios from "axios";

export const deleteItemByID = async (id) => {
    let response;
    const config = {
            method:'delete',
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