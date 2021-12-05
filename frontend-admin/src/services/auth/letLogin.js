import axios from 'axios';
import FormData from 'form-data';
import Cookies from 'js-cookie';

export const letLogin = async () => {
    let data = new FormData();
    data.append('Password', 'Pass1234');
    data.append('Username', 'kasama.tsw@shoptree.com');

    const config = {
        method: 'post',
        url: 'http://spaceship.trueddns.com:23722/api/v1/auth/login',
        headers: { 
        },
        data : data
    };

    let response;
    try {
        response = await axios(config, {withCredentials: true});
        localStorage.setItem("jwt", JSON.stringify(response.data.token));
        Cookies.set("jwt", (response.data.token), { path: '/' });
        console.log('token1 : ', response.data);
        // do something about response
      } catch (error) {
        console.error(error)
      }
    return response.data;
}