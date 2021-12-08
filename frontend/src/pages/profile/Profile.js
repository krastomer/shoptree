import "./Profile.css";
import React, { useEffect, useState, useCallback } from "react";
import Navbar from "../../asset/include/navbar/Navbar";
import allLocation from "./allLocation";
import { getProfile } from "../service/proflie/getProfile";
import AddAddress from "./Add";
import Applogo from "../../asset/include/navbar/LogoBanner.png";
import Statusbar from "../order/Statusbar";

const profiles = getProfile();

export default function Profile() {
  const [profile, setProfile] = useState(null);
  const [local, setLocal] = useState([]);
  const [order, setOrder] = useState([]);
  const [statusorder, setStatusorder] = useState();

  useEffect(() => {
    profiles.then(function (data) {
      setProfile(data.data);
      setLocal(data.data.address);
      setOrder(data.data.orders);
      console.log("order ", data.data.orders)
    });
  });

  if (!profile) return null;
  if(!local) return(
    <>
    <Navbar />
    <div className="max-w-2xl px-4 py-16 mx-auto sm:py-12 sm:px-6 lg:max-w-7xl lg:px-8 font-body">
          <p className="text-4xl text-main-theme font-theme">โปรไฟล์ส่วนตัว</p>
          <div className="flex flex-col pt-2 mt-2">
            <div>
              <font className="text-gray-500">ชื่อ: </font>
              <font className="text-gray-900">{profile.name}</font>
            </div>
            <div>
              <font className="text-gray-500">อีเมล: </font>
              <font className="text-gray-900">{profile.username}</font>
            </div>
            <div>
              <font className="text-gray-500">เบอร์โทรศัพท์: </font>
              <font className="text-gray-900">{profile.phone_number}</font>
            </div>
            <div>

            </div>
          </div>
        </div>
        <div className="max-w-2xl px-4 py-16 mx-auto font-body sm:py-12 sm:px-6 lg:max-w-7xl lg:px-8">
          <p className="text-4xl text-main-theme font-theme">ที่จัดส่ง</p>
          <AddAddress></AddAddress>
        </div>
        <div className="max-w-2xl px-4 py-16 mx-auto font-body sm:py-12 sm:px-6 lg:max-w-7xl lg:px-8">
          <p className="text-4xl text-main-theme font-theme">
            คำสั่งซื้อของฉัน
          </p>
          <h1 className="pt-4 text-2xl text-gary-500">ยังไม่มีคำสั่งซื้อ</h1>
        </div>
    </>
  );
  return (
    <div className="bg-white">
      <Navbar />
      <div className="bg-white font-body">
        <div className="max-w-2xl px-4 py-16 mx-auto sm:py-12 sm:px-6 lg:max-w-7xl lg:px-8">
          <p className="text-4xl text-main-theme font-theme">โปรไฟล์ส่วนตัว</p>
          <div className="flex flex-col pt-2 mt-2">
            <div>
              <font className="text-gray-500">ชื่อ: </font>
              <font className="text-gray-900">{profile.name}</font>
            </div>
            <div>
              <font className="text-gray-500">อีเมล: </font>
              <font className="text-gray-900">{profile.username}</font>
            </div>
            <div>
              <font className="text-gray-500">เบอร์โทรศัพท์: </font>
              <font className="text-gray-900">{profile.phone_number}</font>
            </div>
            <div>
            </div>
          </div>
        </div>
        <div className="max-w-2xl px-4 py-16 mx-auto sm:py-12 sm:px-6 lg:max-w-7xl lg:px-8">
          <p className="text-4xl text-main-theme font-theme">ที่จัดส่ง</p>
          <div className="grid grid-cols-1 mt-2 gap-y-10 gap-x-6 sm:grid-cols-2 lg:grid-cols-3 xl:gap-x-8">
            {local.map((location) => (
              <>
                <div className="flex flex-row max-w-sm p-4 text-black bg-white border border-gray-200 rounded-lg shadow-md">
                  <a href="#" className="px-2">
                    <p className="mb-2 text-lg font-bold tracking-tight text-gray-900">
                      ชื่อ {location.name}
                    </p>
                    <p className="font-normal text-gray-700">
                      ประเทศ {location.country} เมือง {location.city}
                    </p>
                    <p className="font-normal text-gray-700">
                      เขต/อำเภอ {location.district} แขวง/ตำบล {location.state}
                    </p>
                    <p className="font-normal text-gray-700">
                      รหัสไปษรณีย์ {location.postal_code}
                    </p>
                    <p className="font-normal text-gray-700">
                      เบอร์ติดต่อ {location.phone_number}
                    </p>
                  </a>
                  <div className="flex flex-col px-4">
                    <p>&nbsp;</p>
                    <p>&nbsp;</p>
                    <a href="./profile">
                      <svg
                        xmlns="http://www.w3.org/2000/svg"
                        width="24"
                        height="24"
                        viewBox="0 0 24 24"
                      >
                        <path d="M3 6v18h18v-18h-18zm5 14c0 .552-.448 1-1 1s-1-.448-1-1v-10c0-.552.448-1 1-1s1 .448 1 1v10zm5 0c0 .552-.448 1-1 1s-1-.448-1-1v-10c0-.552.448-1 1-1s1 .448 1 1v10zm5 0c0 .552-.448 1-1 1s-1-.448-1-1v-10c0-.552.448-1 1-1s1 .448 1 1v10zm4-18v2h-20v-2h5.711c.9 0 1.631-1.099 1.631-2h5.315c0 .901.73 2 1.631 2h5.712z" />
                      </svg>
                    </a>
                    <p>&nbsp;</p>
                    <p>&nbsp;</p>
                  </div>
                </div>
              </>
            ))}
            <AddAddress></AddAddress>
          </div>
        </div>

        <div className="max-w-2xl px-4 py-16 mx-auto sm:py-12 sm:px-6 lg:max-w-7xl lg:px-8">
          <p className="text-4xl text-main-theme font-theme">
            คำสั่งซื้อของฉัน
          </p>
          <div className="flex flex-col mt-2">
            {order.map((item) => (
              <>
                <button
                  type="button"
                  onClick={() => setStatusorder(!statusorder)}
                >
                  <div className="w-full pt-4 pb-4">
                    <div className="p-4 rounded-lg bg-theme">
                      <div className="flex flex-row justify-between">
                        <div>
                          <font className="text-white ">เลขคำสั่งซื้อ </font>
                          <font className="font-medium text-white">
                            #{item.id}
                          </font>
                        </div>
                        <div>
                          <font className="font-medium text-white">
                            {item.status}
                          </font>
                        </div>
                      </div>
                    </div>
                  </div>
                </button>
                {statusorder ? (
                  <>
                     <div className="w-full h-full">
                      <div className="rounded-lg bg-gray-50 font-body">
                        <div className="max-w-2xl px-4 py-8 mx-auto rounded-lg sm:py-8 sm:px-6 lg:max-w-7xl lg:px-8">
                          <div className="grid items-center justify-center grid-cols-1 mt-2 gap-y-10 gap-x-6 xl:gap-x-8 sm:grid-cols-1 lg:grid-cols-4">
                            <div className="">
                              <img classname="object-contain" src={Applogo} alt="Applogo" />
                            </div>
                            <div class="col-span-2">
                              <Statusbar />
                            </div>
                            <div>
                                  <div className="flex flex-row p-4 pt-2 pb-2 mt-2 mb-2 bg-white border border-gray-200 rounded-lg shadow-md max-w">
                                    <div href="#" className="px-2">
                                    <p className="mb-2 text-lg font-bold tracking-tight text-gray-900">
                                        ชื่อ {item.address.name}
                                    </p>
                                    <p className="font-normal text-gray-700">
                                        ประเทศ {item.address.country} เมือง {item.address.city}
                                    </p>
                                    <p className="font-normal text-gray-700">
                                        เขต/อำเภอ {item.address.district} แขวง/ตำบล {item.address.state}
                                    </p>
                                    <p className="font-normal text-gray-700">
                                        รหัสไปษรณีย์ {item.address.postal_code}
                                    </p>
                                    <p className="font-normal text-gray-700">
                                        เบอร์ติดต่อ {item.address.phone_number}
                                    </p>
                                    </div>
                                  </div>
                            </div>
                          </div>
                          <div className="grid justify-between grid-cols-4 mt-2 gap-y-10 gap-x-6 xl:gap-x-8 ">
                            <div className="relative flex justify-center ">
                              <img
                                src={
                                  "https://www.jorakay.co.th/ckfinder/userfiles/images/%E0%B8%9E%E0%B8%B2%E0%B8%A2%E0%B8%B8%E0%B8%9D%E0%B8%99%E0%B9%80%E0%B8%AA%E0%B8%B5%E0%B9%88%E0%B8%A2%E0%B8%87%E0%B8%97%E0%B8%B3%E0%B8%95%E0%B9%89%E0%B8%99%E0%B9%84%E0%B8%A1%E0%B9%89%E0%B8%A5%E0%B9%89%E0%B8%A1%E0%B8%97%E0%B8%B1%E0%B8%9A%E0%B8%9A%E0%B9%89%E0%B8%B2%E0%B8%99%20%E0%B8%A1%E0%B8%B5%E0%B8%A7%E0%B8%B4%E0%B8%98%E0%B8%B5%E0%B8%9B%E0%B9%89%E0%B8%AD%E0%B8%87%E0%B8%81%E0%B8%B1%E0%B8%99%E0%B8%AD%E0%B8%A2%E0%B9%88%E0%B8%B2%E0%B8%87%E0%B9%84%E0%B8%A3%20%20%E0%B9%83%E0%B8%99%201.jpg"
                                }
                                className="object-contain w-10 h-10 rounded-full"
                              />
                            </div>
                            <div className="relative flex items-center justify-center">
                              <p className="text-gray-500">T001</p>
                            </div>
                            <div className="relative flex items-center justify-center">
                              <p className="font-green">ต้นกุหลาบหิน</p>
                            </div>
                            <div className="relative flex items-center justify-center">
                              <p className="text-black">2500 บาท</p>
                            </div>
                          </div>

                          <div className="flex justify-between">
                            <div>
                              <p className="mt-4 text-xl font-bold font-green">ราคา 2500 บาท</p>
                            </div>
                            {/* <div className="flex flex-col px-4">
                              <img
                                src={
                                  "https://www.jorakay.co.th/ckfinder/userfiles/images/%E0%B8%9E%E0%B8%B2%E0%B8%A2%E0%B8%B8%E0%B8%9D%E0%B8%99%E0%B9%80%E0%B8%AA%E0%B8%B5%E0%B9%88%E0%B8%A2%E0%B8%87%E0%B8%97%E0%B8%B3%E0%B8%95%E0%B9%89%E0%B8%99%E0%B9%84%E0%B8%A1%E0%B9%89%E0%B8%A5%E0%B9%89%E0%B8%A1%E0%B8%97%E0%B8%B1%E0%B8%9A%E0%B8%9A%E0%B9%89%E0%B8%B2%E0%B8%99%20%E0%B8%A1%E0%B8%B5%E0%B8%A7%E0%B8%B4%E0%B8%98%E0%B8%B5%E0%B8%9B%E0%B9%89%E0%B8%AD%E0%B8%87%E0%B8%81%E0%B8%B1%E0%B8%99%E0%B8%AD%E0%B8%A2%E0%B9%88%E0%B8%B2%E0%B8%87%E0%B9%84%E0%B8%A3%20%20%E0%B9%83%E0%B8%99%201.jpg"
                                }
                                className="object-contain w-20 h-20 rounded-full"
                              />
                            </div> */}
                          </div>
                        </div>
                      </div>
                    </div>
                  </>
                ) : null}
              </>
            ))}
          </div>
        </div>
      </div>
    </div>
  );
}
