import "./Login.css";
import Applogo from '../../asset/logo.png';
import React, { useState } from "react";
import Register from "../register/Register";
import { Link } from 'react-router';


export default function Login() {
  return (
    <div className="grid  md:grid-cols-2 h-screen font-prompt font-body ">
        <div className='md:flex  hidden flex bg-primary text-white text-left  bg-green-600' style={{boxShadow:'0 4px 4px #000'}}>
        <div >
            <p className="text-left mx-3 my-10 top-10 row-auto"><a href="#" className="font-medium">กลับสู่หน้าหลัก</a></p>
            <img src={Applogo} alt="Logo" className='w-4/6 my-auto mx-auto justify-center ' />
        </div>
        </div>
        <div className="min-h-full flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
        <div className="max-w-md w-full space-y-8">
        <div>
            <h2 className="mt-6 text-left text-3xl font-extrabold text-gray-900">ล็อกอิน</h2>
        </div>
        <form className="mt-8 space-y-6" action="#" method="POST">
            <input type="hidden" name="remember" defaultValue="true" />
            <div className="rounded-md shadow-sm -space-y-px">

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
                className="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-t-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 sm:text-sm"
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
                className="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-b-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 sm:text-sm"
                />
            </div>
            </div>
            <div className="flex items-center justify-between">
            <div className="flex items-center">
            </div>
            <div >
                <button
                type="submit"
                className="group relative w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white btn-theme hover:bg-yellow-00 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
                >
                เข้าสู่ระบบ
                </button>
            </div>
            </div>
            <div>
            <div className=" text-center"> 
                <a href = "#"   className="font-medium text-gray-600 hover:text-gray-400">
                ถ้ายังไม่มีบัญชี คลิ๊กเพื่อสมัครสมาชิก
                </a>
            </div>
            </div>
        </form>
        </div>
    </div>
  </div>
  )
}
