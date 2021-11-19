import Order from "../pages/order/Order"

export const LoginUser ={
    username:"",
    password:"",
    auth:{
        loggedIn:false,
        token:""
    },
    level:{
        user:false,
        admin:false
    },
    basket:{
        order:"",
        state:1
    }
}
export const AuthUser =  {
    access_token : "",
    email : "",
    updated_timestamp : ""
}

export const RegisterUser ={
    name :"",
    username:"",
    password:"",
    phone:""
}