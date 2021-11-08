import react from "react";
import { useParams } from "react-router";
export default function Detail(data){
    const {id} = useParams()
    return(
        <div>Product ={id}</div>
    )
}