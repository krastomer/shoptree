import axios from "axios"
export const products = async () => {
    let response;
    try {
        response = await axios.get("http://spaceship.trueddns.com:23720/api/v1/products/1")
    } catch (error) {
        console.error(error)
    }
    return response?.data ? response?.data : null // or set initial value
}  
