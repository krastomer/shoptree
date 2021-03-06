import axios from "axios";
export const getReviews = async () => {
    let response;

    const config = {
            method: 'get',
            url: 'http://spaceship.trueddns.com:23720/api/v1/reviews',
    }
    try {
        response = await axios(config)
        console.log("data :", response.data)
    } catch (error) {
        console.error(error)
    }
    return response?.data ? response?.data : null // or set initial value
}