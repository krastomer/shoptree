import "./Register.css";
import Applogo from "../../logo.svg";
import React, { useState } from "react";
import Login from "../login/Login";

export default function Register() {
  const [currentUser, setCurrentUser] = useState(null);
  const handleSubmit = (e) => {
    try {
      e.preventDefault();
      var request = require("request");
      var options = {
        method: "GET",
        url: "spaceship.trueddns.com:23720/api/v1/products/1",
        headers: {},
      };
      request(options, function (error, response) {
        if (error) throw new Error(error);
        console.log(response.body);
      });
    } catch (error) {
      alert(error);
    }
  };
  return (
    <div className="grid  md:grid-cols-2 h-screen font-body ">
      <div
        className="md:flex hidden flex bg-primary bg-green-600 flex-col "
        style={{ boxShadow: "0 4px 4px #000" }}
      >
        <div className="my-10 text-white">
          &nbsp;&nbsp; &nbsp;&nbsp; กลับสู่หน้าหลัก
        </div>
        <div className="mx-auto my-auto ">
          <img src={Applogo} alt="Logo" />
        </div>
      </div>
      <div className="min-h-full flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
        <div className="max-w-md w-full space-y-8">
          <div>
            <h2 className="mt-6 text-left text-3xl font-extrabold text-gray-900">
              สมัครสมาชิก
            </h2>
          </div>
          <form className="mt-8 space-y-6" action="#" method="POST">
            <input type="hidden" name="remember" defaultValue="true" />
            <div className="rounded-md shadow-sm -space-y-px">
              <div>
                <p className=" text-gray-500">ชื่อ-นามสกุล</p>
                <label htmlFor="password" className="sr-only">
                  Name
                </label>
                <input
                  id="name"
                  name="name"
                  type="text"
                  required
                  className="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-b-md focus:outline-none focus:ring-green-500 focus:border-green-500 focus:z-10 sm:text-sm"
                />
              </div>
              <div>
                <p className=" text-gray-500">อีเมล</p>
                <label htmlFor="email-address" className="sr-only">
                  Email address
                </label>
                <input
                  id="email-address"
                  name="email"
                  type="email"
                  autoComplete="email"
                  required
                  className="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-t-md focus:outline-none focus:ring-green-500 focus:border-green-500 focus:z-10 sm:text-sm"
                />
              </div>
              <div>
                <p className=" text-gray-500">รหัสผ่าน</p>
                <label htmlFor="password" className="sr-only">
                  Password
                </label>
                <input
                  id="password"
                  name="password"
                  type="password"
                  autoComplete="current-password"
                  required
                  className="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-b-md focus:outline-none focus:ring-green-500 focus:border-green-500 focus:z-10 sm:text-sm"
                />
              </div>
              <div>
                <p className=" text-gray-500">ยืนยันรหัสผ่าน</p>
                <label htmlFor="password" className="sr-only">
                  RePassword
                </label>
                <input
                  id="repassword"
                  name="repassword"
                  type="password"
                  autoComplete="current-password"
                  required
                  className="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-b-md focus:outline-none focus:ring-green-500 focus:border-green-500 focus:z-10 sm:text-sm"
                />
              </div>
              <div>
                <p className=" text-gray-500">เบอร์โทรศัพท์</p>
                <label htmlFor="password" className="sr-only">
                  Phone
                </label>
                <input
                  id="phone"
                  name="phone"
                  type="text"
                  required
                  className="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-b-md focus:outline-none focus:ring-green-500 focus:border-green-500 focus:z-10 sm:text-sm"
                />
              </div>
            </div>
            <div className="flex items-center justify-between">
              <div className="flex items-center">
                <a
                  href="#"
                  className="font-medium text-gray-600 hover:text-gray-400"
                >
                  ฉันยอมรับข้อตกลง เพิ่มเติม
                </a>
              </div>
              <div>
                <form onClick={handleSubmit}>
                  <button
                    type="click"
                    className="group relative w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white btn-theme hover:bg-yellow-00 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500"
                  >
                    สมัครสมาชิก
                  </button>
                </form>
              </div>
            </div>
            <div>
              <div className=" text-center">
                <a
                  href="#"
                  className="font-medium text-gray-600 hover:text-gray-400"
                >
                  ถ้ามีบัญชีแล้ว คลิ๊กเพื่อเข้าสู่ระบบ
                </a>
              </div>
            </div>
          </form>
        </div>
      </div>
    </div>
  );
}
