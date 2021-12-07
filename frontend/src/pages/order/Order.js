import React, { useState, Fragment, useEffect } from "react";
import Navbar from "../../asset/include/navbar/Navbar";
import Statusbar from "./Statusbar";
import Add from "../review/add.svg";
import { Link } from "react-router-dom";
import NumberFormat from "react-number-format";
import { Route, useParams, useHistory } from "react-router";
import { LoginUser } from "../../models/User";
import { SuccessOrder } from "../success/success";
import allLocation from "../profile/allLocation";
import EditAddress from "../profileEdit/Editaddress";
import "./Order.css";
import Applogo from "../../asset/ConfirmOrder.svg"
import { getCart } from "../service/orders/getCart";
import { deleteItemByID } from "../service/deleteCart/deleteCart";

export default function Order() {
  const item = getCart();
  let history = useHistory();
  const { USER } = useParams();
  const [activeState, setState] = useState(LoginUser.basket.state);
  const [content1, setConttent1] = useState();
  const [content2, setConttent2] = useState();
  const [content3, setConttent3] = useState();
  const [content4, setConttent4] = useState();
  const [content5, setConttent5] = useState();
  const [processBar, setProcessBar] = useState();
  const [editAddress, setEditAddress] = useState();
  const [orders, setOrders] = useState([]);
  const [actorder, serActorder] = useState(null);
  useEffect(() => {
    LoginUser.basket.state = activeState;
    if (activeState === 1) {
      setConttent1(true);
      setConttent2(false);
      setProcessBar(true);
    }
    if (activeState === 2) {
      setConttent1(false);
      setConttent2(true);
      setConttent3(false);
      setProcessBar(true);
    }
    if (activeState === 3) {
      setConttent2(false);
      setConttent3(true);
      setConttent4(false);
      setProcessBar(true);
    }
    if (activeState === 4) {
      setConttent3(false);
      setConttent4(true);
      setProcessBar(true);
    }
    if (activeState === 5) {
      setConttent4(false);
      setConttent5(true);
      setProcessBar(false);
    }
    if (actorder === null)
      item.then(function (data) {
        if (data) {
          console.log("data ", data.data.products);
          setOrders(data.data.products);
          serActorder("active");
        }else{
          serActorder("null")
        }
      });
  });
  if (!orders)
    return (
      <>
        <div className="bg-white">
          <Navbar />
          <div className="max-w-2xl px-4 mx-auto sm:px-6 lg:max-w-7xl lg:px-8 font-body">
            <h2 className="py-4 text-2xl tracking-tight text-gray-600">
              ไม่มีสินค้าในตะกร้า
            </h2>
          </div>
        </div>
      </>
    );
  function getPrevStepContent(stepIndex) {
    if (stepIndex > 1) {
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
        return "";
    }
  }
  function GoPrev() {
    setState(activeState - 1);
  }
  const GoNextt = () => {
    setState(activeState + 1);
  };

  if (!item) {
    return null;
  }
  const deleteOrder = async (e) => {
    e.preventDefault();
    const deleteorders = await deleteItemByID(e.target.value);
    window.location.reload();
    console.log(deleteorders);
  };
  //  calculat count of order and find total of product in order.
  const count = orders.length;
  const sum = orders.reduce((total,price)=> total =total+price.price, 0);

  return (
    <div className="bg-white">
      <Navbar />
      <div className="max-w-2xl px-4 mx-auto sm:px-6 lg:max-w-7xl lg:px-8 font-body">
        <Statusbar stateLocal={activeState} />
        {content1 ? (
          <>
            <h2 className="py-4 text-2xl tracking-tight text-gray-600">
              สินค้าในตะกร้า
            </h2>
            <div className="grid grid-cols-1 mt-3 gap-y-10 gap-x-6 sm:grid-cols-2 lg:grid-cols-4 xl:gap-x-8">
              {orders.map((product) => (
                <div key={product.id} className="relative group">
                  <div className="w-full overflow-hidden bg-gray-200 rounded-md min-h-80 aspect-w-1 aspect-h-1 lg:h-80 lg:aspect-none">
                    <Link to={`/products/${product.id}`}>
                      <img
                        src={`http://spaceship.trueddns.com:23720/api/v1/products/images/${product.image_path}`}
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
                          <button
                            value={product.id}
                            onClick={deleteOrder}
                            className="px-4 py-2 font-bold text-white bg-red-500 rounded-full hover:bg-red-700"
                          >
                            ลบสินค้านี้ออกจากตะกร้า
                          </button>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              ))}
            </div>
          </>
        ) : null}
        {content2 ? (
          <>
            <div className="max-w-2xl px-4 py-16 mx-auto sm:py-12 sm:px-6 lg:max-w-7xl lg:px-8">
              <p className="text-4xl text-main-theme font-theme">
                เลือกที่จัดส่ง
              </p>
              <div className="flex flex-col mt-2">
                {allLocation.map((location) => (
                  <>
                    <div className="flex flex-row p-4 pt-2 pb-2 mt-2 mb-2 bg-white border border-gray-200 rounded-lg shadow-md max-w hover:bg-red-700">
                      <a href="#" className="px-2">
                        <p className="mb-2 text-lg font-bold tracking-tight text-gray-900">
                          {location.name}
                        </p>
                        <p class="font-normal text-gray-700">
                          {location.disFirst}
                        </p>
                        <p class="font-normal text-gray-700">
                          {location.disSecond}
                        </p>
                        <p class="font-normal text-gray-700">
                          {location.postNumber}
                        </p>
                        <p class="font-normal text-gray-700">
                          {location.phoneNumber}
                        </p>
                      </a>
                    </div>
                  </>
                ))}
                <button
                  className="pt-2 mt-2 text-black border-4 border-dashed font-body md:border-dashed"
                  type="button"
                  onClick={() => setEditAddress(!editAddress)}
                >
                  <div className="flex flex-col items-center p-5 font-bold leading-snug font-theme ">
                    เพิ่มที่อยู่
                    <div className="flex p-2">
                      <img src={Add} alt="Add" />
                    </div>
                  </div>
                </button>
              </div>
            </div>
            {editAddress ? (
              <>
                <button
                  className="float-right p-1 ml-auto text-3xl font-semibold leading-none text-black bg-transparent border-0 outline-none focus:outline-none"
                  onClick={() => setEditAddress(false)}
                >
                  <div>×</div>
                </button>
                <EditAddress name={"เพิ่มที่จัดส่งใหม่"}></EditAddress>
              </>
            ) : null}
          </>
        ) : null}
        {content3 ? (
          <>
            <h2 className="py-4 text-2xl tracking-tight text-gray-600">
              ตรวจสอบคำสั่งซื้อ
            </h2>
            <div className="grid grid-cols-1 mt-3 gap-y-10 gap-x-6 sm:grid-cols-2 lg:grid-cols-4 xl:gap-x-8">
              {orders.map((product) => (
                <div key={product.id} className="relative group">
                  <div className="w-full overflow-hidden bg-gray-200 rounded-md min-h-80 aspect-w-1 aspect-h-1 lg:h-80 lg:aspect-none">
                    <Link to={`/products/${product.id}`}>
                      <img
                        src={`http://spaceship.trueddns.com:23720/api/v1/products/images/${product.image_path}`}
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
                    </div>
                  </div>
                </div>
              ))}
            </div>
            <div>
              <hr className="pt-2 mt-4"></hr>
              <h2 className="py-4 text-2xl tracking-tight text-gray-600">
                ที่อยู่
              </h2>
              <div className="flex flex-row p-4 pt-2 pb-2 mt-2 mb-2 bg-white border border-gray-200 rounded-lg shadow-md max-w">
                <div  className="px-2">
                  <p className="mb-2 text-lg font-bold tracking-tight text-gray-900">
                    {"test"}
                  </p>
                  <p class="font-normal text-gray-700">
                    {"test"}
                  </p>
                  <p class="font-normal text-gray-700">
                    {"test"}
                  </p>
                  <p class="font-normal text-gray-700">
                    {"test"}
                  </p>
                  <p class="font-normal text-gray-700">
                   {"test"}
                  </p>
                </div>
              </div>
            </div>  
          </>
        ) : null}
        {content4 ? (
          <>
            <h2 className="py-4 text-2xl tracking-tight text-gray-600">
              เลขที่คำสั่งซื้อ
            </h2>
            <h1 className="py-4 text-3xl tracking-tight text-gray-600">#01</h1>
            <h2 className="py-4 text-2xl tracking-tight text-gray-600">
              รายการคำสั่งซื้อ
            </h2>
            <div className="grid grid-cols-1 mt-3 gap-y-10 gap-x-6 sm:grid-cols-2 lg:grid-cols-4 xl:gap-x-8">
              {orders.map((product) => (
                <div key={product.id} className="relative group">
                  <div className="w-full overflow-hidden bg-gray-200 rounded-md min-h-80 aspect-w-1 aspect-h-1 lg:h-80 lg:aspect-none">
                    <Link to={`/products/${product.id}`}>
                      <img
                        src={`http://spaceship.trueddns.com:23720/api/v1/products/images/${product.image_path}`}
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
                    </div>
                  </div>
                </div>
              ))}
            </div>
            <h2 className="py-4 text-2xl tracking-tight text-gray-600">
              ส่งหลักฐานการโอนเงิน
            </h2>
            <button
              className="text-black border-4 border-dashed font-body md:border-dashed"
              type="button"
              onClick={() => setEditAddress(!editAddress)}
            >
              <div className="flex flex-col items-center p-5 font-bold leading-snug font-theme ">
                อัพโหลดสลิป
                <div className="flex p-2">
                  <img src={Add} alt="Add" />
                </div>
              </div>
            </button>
          </>
        ) : null}
        {content5 ? (
          <>
            <div className="grid content-center grid-cols-1 mt-6 text-center">
                  <h1 className="text-2xl ">ยืนยันคำสั่งซื้อ <font className="text-green-500">#01</font> เรียบร้อย</h1>
                  <h1 className="pb-6 text-2xl">ขอขอบคุณที่สั่งต้นไม้จาก SHOPTREE</h1>
                  <div className="grid grid-cols-3 justify-items-center">
                    <p>&nbsp;</p>
                    <img className="object-none object-center" src={Applogo} alt="Logo" />
                    <p>&nbsp;</p>
                  </div>
                  <div className="flex justify-center">
                      <Link className="pr-4" to="/">กลับสู่หน้าหลัก</Link>
                      <Link className="pl-8" to="/profile">ดูรายละเอียดคำสั่งซื้อ</Link>
                  </div>
            </div>
          </>
        ) : null}
         {processBar ? (
          <>
            <hr className="pt-4 mt-4 mb-2"></hr>
            <div className="grid grid-cols-2 pb-4 md:grid-cols-4 lg:grid-cols-4">
              <div className="flex order-3 md:order-1 lg:order-1">
                <button disabled={activeState === 0} onClick={GoPrev}>
                  <div className="text-xl font-semibold">
                    <font className="px-6 text-2xl font-bold text-green-500">
                      {getPrevStepContent(activeState)}
                    </font>
                  </div>
                </button>
              </div>
              <div className="flex order-1 md:order-2 lg:order-2">
                <p className="text-xl font-semibold ">
                  จำนวนสินค้าทั้งหมด{" "}
                  <font className="text-2xl font-bold text-red-700">{count}</font>{" "}
                  ชิ้น
                </p>
              </div>
              <div className="flex order-2 md:order-3 lg:order-3">
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
              <div className="flex order-4">
                <button disabled={activeState === 5} onClick={GoNextt}>
                  <div className="px-6 text-2xl font-bold text-green-500">
                    {getNextStepContent(activeState)}{" "}
                  </div>
                </button>
              </div>
            </div>
            </>
        ) : null}
      </div>
    </div>
  );
}
