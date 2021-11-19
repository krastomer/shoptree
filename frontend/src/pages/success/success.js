import React ,{}from "react";
import { Route, useParams } from "react-router";
export function SuccessOrder(){
    const USER = useParams()
    return (
        <hr>Success your user{USER}</hr>
    );
}