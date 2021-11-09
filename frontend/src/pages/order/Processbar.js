import React, { useState } from "react";
import allOrder from "./allOrder";
import "./Order.css";
import NumberFormat from 'react-number-format';

const products = allOrder
const count = products.length
const sum = products.map(product => product.price).reduce((a, b) => a + b)

function GoBack(props){
  const {stateLocal} = props;

  if(stateLocal<=1){
    return(<div></div>)
  }
  else if(stateLocal==2){
    return(
      <a href="/order/1" className="text-xl font-semibold">
        <font className ="px-6 text-2xl font-bold text-green-500">ย้อนกลับ</font>
      </a>
    )
  }
  else if(stateLocal==3){
    return(
      <a href="/order/2" className="text-xl font-semibold">
        <font className ="px-6 text-2xl font-bold text-green-500">ย้อนกลับ</font>
      </a>
    )
  }
  else if(stateLocal==4){
    return(
      <a href="/order/3" className="text-xl font-semibold">
        <font className ="px-6 text-2xl font-bold text-green-500">ย้อนกลับ</font>
      </a>
    )
  }
  else if(stateLocal==5){
    return(
      <a href="/order/4" className="text-xl font-semibold">
        <font className ="px-6 text-2xl font-bold text-green-500">ย้อนกลับ</font>
      </a>
    )
  }
  else{
    return(<div></div>)
    }
}
function GoNext(props){
  const {stateLocal} = props;
  if(stateLocal==1){
    return(
      <a href="/order/2" className ="px-6 text-2xl font-bold text-green-500">เลือกที่จัดส่ง</a>
    )
  }
  else if(stateLocal==2){
    return(
      <a href="/order/3" className ="px-6 text-2xl font-bold text-green-500">ตรวจสอบคำสั่งซื้อ</a>
    )
  }
  else if(stateLocal==3){
    return(
      <a href="/order/4" className ="px-6 text-2xl font-bold text-green-500">ชำระเงิน</a>
    )
  }
  else if(stateLocal==4){
    return(
      <a href="/order/5" className ="px-6 text-2xl font-bold text-green-500">เสร็จสิ้น</a>
    )
  }
  else{
      return(
        <div></div>
      )
    }
}

export default function Processbar(props) {
  if(props.stateLocal ==5){
    return (
      <div></div>
    )
  }else{
    return (
      <>   
        <div className="p-4 mt-8">
        <hr></hr>
            <div className="flex px-2 py-4">
                <GoBack stateLocal = {props.stateLocal}/>
                <p className="text-xl font-semibold ">
                    จำนวนสินค้าทั้งหมด <font className ="text-2xl font-bold text-red-700">{count}</font> ชิ้น
                </p>
                <div className="flex flex-row-reverse flex-auto ">
                  <GoNext stateLocal = {props.stateLocal}/>
                  <p className="text-xl font-semibold ">
                    ราคาทั้งหมด <font className ="text-2xl font-bold text-red-700">
                    <NumberFormat value={sum} displayType={'text'} thousandSeparator={true} prefix={'$'} />       
                      </font> บาท
                  </p>
                </div>
            </div>
        </div>
      </>
    )
  }
}
  