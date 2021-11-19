import "./Profile.css";
import React, { useEffect, useState } from "react";
import Navbar from "../../asset/include/navbar/Navbar";
import allLocation from "./allLocation";

import AddAddress from "./Add"

const locations = allLocation;

export default function Profile() {
  return (
    <div className="bg-white">
      <Navbar />
      <div className="bg-white font-body">
        <div className="max-w-2xl px-4 py-16 mx-auto sm:py-12 sm:px-6 lg:max-w-7xl lg:px-8">
          <p className="text-4xl text-main-theme font-theme">โปรไฟล์ส่วนตัว</p>
          <div className="flex flex-col pt-2 mt-2">
            <div>
              <font className="text-gray-500">ชื่อ: </font>
              <font className="text-gray-900">ทรัพทวี ขี่ฮอนด้า</font>
            </div>
            <div>
              <font className="text-gray-500">อีเมล: </font>
              <font className="text-gray-900">subtawee@shoptree.com</font>
            </div>
            <div>
              <font className="text-gray-500">เบอร์โทรศัพท์: </font>
              <font className="text-gray-900">080-808-4545</font>
            </div>
            <div>
              <font className="text-gray-500">เปลี่ยนรหัสผ่าน </font>
            </div>
          </div>
        </div>
        <div className="max-w-2xl px-4 py-16 mx-auto sm:py-12 sm:px-6 lg:max-w-7xl lg:px-8">
          <p className="text-4xl text-main-theme font-theme">ที่จัดส่ง{allLocation.length}</p>
          <div className="grid grid-cols-1 mt-2 gap-y-10 gap-x-6 sm:grid-cols-2 lg:grid-cols-3 xl:gap-x-8">
            {locations.map((location) => (
              <>
                <div className="flex flex-row max-w-sm p-4 bg-white border border-gray-200 rounded-lg shadow-md">
                  <a href="#" className="px-2">
                    <p className="mb-2 text-lg font-bold tracking-tight text-gray-900">
                      {location.name}
                    </p>
                    <p class="font-normal text-gray-700">{location.disFirst}</p>
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
            {locations.map((location) => (
              <>
                <div className="w-full my-4">
                  <div className="p-4 rounded-lg bg-theme">
                    <div className="flex flex-row justify-between">
                      <div>
                        <font className="text-white ">เลขคำสั่งซื้อ </font>
                        <font className="font-medium text-white">
                          #00001234
                        </font>
                      </div>
                      <div>
                        <font className="font-medium text-white">
                          สถานะกำลังจัดส่ง
                        </font>
                      </div>
                    </div>
                  </div>
                </div>
              </>
            ))}
          </div>
        </div>
      </div>
    </div>
  );
}
