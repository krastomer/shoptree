import React from 'react';
import axios from "axios";
import Cookies from 'js-cookie';
import { letLogin } from "../auth/letLogin";
import { useState, useEffect } from 'react';

export const products = async () => {
    let response;
    try{
        let accessToken = Cookies.get('jwt');
        console.log("token2 : ", accessToken)
        const config = {
            method: 'get',
            url: 'http://spaceship.trueddns.com:23722/api/v1/products',
            headers: { 
                'Cookie': "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEsInVzZXIiOiJrYXNhbWEudHN3QHNob3B0cmVlLmNvbSIsImF1ZCI6IkFkbWluIiwiZXhwIjoxNjM4Njk0NzA5LCJpc3MiOiJzaG9wdHJlZSJ9.s0N6m4ciKtXzyQRzcVXjKHbqqG-NjH1l5zxERWnkMs0",
            }
        };
        response = await axios(config , {withCredentials: true});
        console.log('response: ', response);
    } catch (error) {
        console.error(error)
    }
    return response?.data ? response?.data : null
}  
