import React, { useState, Fragment, useEffect } from "react";
import Navbar from "../../asset/include/navbar/Navbar";
import Timebar from "./Timebar";
import Statusbar from "./Statusbar";

import allOrder from "./allOrder";
import { Link } from "react-router-dom";
import NumberFormat from "react-number-format";
import { Route, useParams,useHistory } from "react-router";
import { LoginUser } from "../../models/User";
import { SuccessOrder } from "../success/success";
const products = allOrder;

export default function Order() {
  let history = useHistory()
  const products = allOrder;
  const count = products.length;
  const sum = products.map((product) => product.price).reduce((a, b) => a + b);
  const { USER } = useParams();
  const [activeState, setState] = useState(LoginUser.basket.state);
  useEffect(() => {
    LoginUser.basket.state = activeState;
    if(LoginUser.basket.state === 5){
      history.push(`/order/${USER}/Success`);
    }
  });
  function getPrevStepContent(stepIndex) {
    if (stepIndex >1){
      return "ย้อนกลับ";
    }
  }
  function getNextStepContent(stepIndex) {
    switch (stepIndex) {
      case 1:
        return "เลือกที่จัดส่ง";
      case 2:
        return "ตรวจสอบคำสั่งซื้อ";
      case 3:
        return "ชำระเงิน";
      case 4:
        return "เสร็จสิ้น";
      default:
        return "Unknown stepIndex";
    }
  }
  function GoPrev() {
    setState(activeState - 1);
  }
  const GoNextt = () => {
    setState(activeState + 1);
  };
  return (
    <div className="bg-white">
      <Navbar />
      <Timebar />
      <div className="max-w-2xl px-4 mx-auto sm:px-6 lg:max-w-7xl lg:px-8 font-body">
        <Statusbar stateLocal={activeState} />
        <h2 className="py-4 text-2xl tracking-tight text-gray-600">
          สินค้าในตะกร้า
        </h2>
        <div className="grid grid-cols-1 mt-3 gap-y-10 gap-x-6 sm:grid-cols-2 lg:grid-cols-4 xl:gap-x-8">
          {products.map((product) => (
            <div key={product.id} className="relative group">
              <div className="w-full overflow-hidden bg-gray-200 rounded-md min-h-80 aspect-w-1 aspect-h-1 lg:h-80 lg:aspect-none">
                <Link to={`/products/${product.id}`}>
                  <img
                    src={product.imageSrc}
                    alt={product.imageAlt}
                    className="object-cover object-center w-full h-full lg:w-full lg:h-full"
                  />
                </Link>
              </div>
              <div className="flex flex-col mt-4">
                <div className="flex flex-row justify-between">
                  <h3 className="text-sm text-right text-gray-700">
                    {product.id}
                  </h3>
                  <h3 className="text-sm text-right text-gray-700">
                    {product.name}
                  </h3>
                </div>
                <div>
                  <div className="flex flex-row justify-between">
                    <p className="text-sm font-medium text-right text-gray-900">
                      ราคา
                    </p>
                    <p className="text-sm font-medium text-right text-gray-900">
                      <NumberFormat
                        value={product.price}
                        displayType={"text"}
                        thousandSeparator={true}
                        prefix={"$"}
                      />
                    </p>
                  </div>
                  <div className="flex flex-row justify-center">
                    <div className="">
                      <button className="px-4 py-2 font-bold text-white bg-red-500 rounded-full hover:bg-red-700">
                        ลบสินค้านี้ออกจากตะกร้า
                      </button>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          ))}
        </div>
        <>
          <div className="p-4 mt-8">
            <hr></hr>
            <div className="flex px-2 py-4">
              {/* <GoBack stateLocal = {stateLocal}/> */}
              <button disabled={activeState === 0} onClick={GoPrev}>
                <div className="text-xl font-semibold">
                  <font className="px-6 text-2xl font-bold text-green-500">
                    {getPrevStepContent(activeState)}
                  </font>
                </div>
              </button>
              <p className="text-xl font-semibold ">
                จำนวนสินค้าทั้งหมด{" "}
                <font className="text-2xl font-bold text-red-700">{count}</font>{" "}
                ชิ้น
              </p>
              <div className="flex flex-row-reverse flex-auto ">
                <button disabled={activeState === 5} onClick={GoNextt}>
                  <div className="px-6 text-2xl font-bold text-green-500">
                    {getNextStepContent(activeState)}{" "}
                  </div>
                </button>
                {/* <GoNext stateLocal = {stateLocal}/> */}
                <p className="text-xl font-semibold ">
                  ราคาทั้งหมด{" "}
                  <font className="text-2xl font-bold text-red-700">
                    <NumberFormat
                      value={sum}
                      displayType={"text"}
                      thousandSeparator={true}
                      prefix={"$"}
                    />
                  </font>{" "}
                  บาท
                </p>
              </div>
            </div>
          </div>
        </>
      </div>
    </div>
  );
}
