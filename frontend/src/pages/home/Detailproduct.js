import React from "react";
import { useParams } from "react-router";
export default function Detail(){
    const {id} = useParams()
    return(
        <div>Product ={id}</div>
    )
}