import "./Register.css";
import Applogo from "../../logo.svg";
import React, { useEffect, useState } from "react";
import { postRegister } from "../api/PostRegister";
import Input from "react-validation"
export default function Register() {
  const [Name, setName] = useState(null);
  const [Email, setEmail] = useState(null);
  const [Password, setPassword] = useState(null);
  const [ConfirmPassword, setConfirmPassword] = useState(null);
  const [Phone, setPhone] = useState(null);
  const OnchangeName = (e) => {
    setName(e.target.value);
  };
  const OnchangePassword = (e) => {
    setPassword(e.target.value);
  };
  const OnchangeEmail = (e) => {
    setEmail(e.target.value);
  };
  const OnchangeConfirmPassword = (e) => {
    setConfirmPassword(e.target.value);
  };
  const OnchangePhone = (e) => {
    setPhone(e.target.value);
  };
  const handleSubmit = (e) => {
    e.preventDefault();
    const register = [
      Name,
      Email,
      Password,
      Phone
    ]
    postRegister(register)
  };
  return (
    <div className="grid  md:grid-cols-2 h-screen font-body ">
      <div
        className="md:flex hidden flex bg-primary bg-green-600 flex-col "
        style={{ boxShadow: "0 4px 4px #000" }}
      >
        <div className="my-10 text-white">
          <a href="home">กลับสู่หน้าหลัก</a>
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
          <form onSubmit={handleSubmit}>
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
                  value={Name}
                  onChange={OnchangeName}
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
                  value={Email}
                  onChange={OnchangeEmail}
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
                  value={Password}
                  onChange={OnchangePassword}
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
                  value={ConfirmPassword}
                  onChange={OnchangeConfirmPassword}
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
                  value={Phone}
                  onChange={OnchangePhone}
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
                <button className="group relative w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white btn-theme hover:bg-yellow-00 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500">
                  สมัครสมาชิก
                </button>
              </div>
            </div>
            <div>
              <div className=" text-center">
                <a
                  href="login"
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
