import React, { useState } from "react";
import StarPoint from "./starpoint";
import StarNone from "./starnone";

function StarRender(props){
    const point = props.point;
    if(point == 1){
        return (
            <div className="flex justify-between py-1">
                 <StarPoint />
                 <StarNone />
                 <StarNone />
                 <StarNone />
                 <StarNone />
            </div>
        )
    }else if(point == 2){
        return (
            <div className="flex justify-between py-1">
                 <StarPoint />
                 <StarPoint />
                 <StarNone />
                 <StarNone />
                 <StarNone />
            </div>
        )
    }else if(point == 3){
        return (
            <div className="flex justify-between py-1">
                 <StarPoint />
                 <StarPoint />
                 <StarPoint />
                 <StarNone />
                 <StarNone />
            </div>
        )
    }else if(point == 4){
        return (
            <div className="flex justify-between py-1">
                 <StarPoint />
                 <StarPoint />
                 <StarPoint />
                 <StarPoint />
                 <StarNone />
            </div>
        )
    }else if(point == 5){
        return (
            <div className="flex justify-between py-1">
                 <StarPoint />
                 <StarPoint />
                 <StarPoint />
                 <StarPoint />
                 <StarPoint />
            </div>
        )
    }else {
        return(
            <div className="flex justify-between py-1">
                 <StarNone />
                 <StarNone />
                 <StarNone />
                 <StarNone />
                 <StarNone />
            </div>
        )
    }
}

export default function ItemCard(props) {
    return (
        <div className = "rounded bg-theme w-64 p-2">
            <StarRender point ={props.point}/>
            <div className="mt-2">&nbsp;</div>
            <div className ="text-lg mt-3 text-white">
                {props.detail}
            </div>
            <div className="mt-2">&nbsp;</div>
        </div>
    )
}