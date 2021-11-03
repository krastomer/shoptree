import React, { useState } from "react";
import allOrder from "./allOrder";
import "./Order.css";
import NumberFormat from 'react-number-format';

const products = allOrder
const count = products.length
const sum = products.map(product => product.price).reduce((a, b) => a + b)

export default function Processbar() {
    return (
      <>   
        <div className="mt-8 p-4">
        <hr></hr>
            <div className="flex px-2 py-4">
                <p className=" text-xl font-semibold">
                    จำนวนสินค้าทั้งหมด <font className ="text-2xl font-bold text-red-700">{count}</font> ชิ้น
                </p>
                <div className="flex-auto flex flex-row-reverse ">
                    <a href="" className ="px-6 text-2xl font-bold text-green-500">เลือกที่จัดส่ง</a> 
                  <p className=" text-xl font-semibold">
                    ราคาทั้งหมด <font className ="text-2xl font-bold text-red-700">
                    <NumberFormat value={sum} displayType={'text'} thousandSeparator={true} prefix={'$'} />       
                      </font> บาท
                  </p>
                </div>
            </div>
        </div>
      </>
    );
  }
  