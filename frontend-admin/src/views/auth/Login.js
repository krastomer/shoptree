import Applogo from "../../logo.svg";
import React, { useEffect, useState, useRef } from "react";
import { Link } from "react-router-dom";
import { letLogin } from "../../services/auth/letLogin";
export default function Login() {
    letLogin();
    return (
        <div className="grid h-screen md:grid-cols-2 font-prompt font-body ">
          <div
            className="flex flex-col hidden bg-green-600 md:flex bg-primary "
            style={{ boxShadow: "0 4px 4px #000" }}
          >
          </div>
          <div className="flex items-center justify-center min-h-full px-4 py-12 sm:px-6 lg:px-8">
            <div className="w-full max-w-md space-y-8">
              <div>
                <h2 className="mt-6 text-3xl font-extrabold text-left text-gray-900">
                  ล็อกอิน
                </h2>
              </div>
              <form className="mt-8 space-y-6" action="#" method="POST">
                <input type="hidden" name="remember" defaultValue="true" />
                <div className="-space-y-px rounded-md shadow-sm">
                  <div>
                    <p className="text-gray-500 ">อีเมล</p>
                    <label htmlFor="email-address" className="sr-only">
                      Email address
                    </label>
                    <input
                      type="email"
                      autoComplete="email"
                      required
                      className="relative block w-full px-3 py-2 text-gray-900 placeholder-gray-500 border border-gray-300 rounded-none appearance-none rounded-t-md focus:outline-none focus:ring-green-500 focus:border-green-500 focus:z-10 sm:text-sm"
                    />
                  </div>
                  <div>
                    <p className="text-gray-500 ">รหัสผ่าน</p>
                    <label htmlFor="password" className="sr-only">
                      Password
                    </label>
                    <input
                      id="password"
                      name="password"
                      type="password"
                      autoComplete="current-password"
                      required
                      className="relative block w-full px-3 py-2 text-gray-900 placeholder-gray-500 border border-gray-300 rounded-none appearance-none rounded-b-md focus:outline-none focus:ring-green-500 focus:border-green-500 focus:z-10 sm:text-sm"
                    />
                  </div>
                </div>
                <div className="flex items-center justify-between">
                  <div className="flex items-center"></div>
                  <div>
                    <Link
                        to="/"
                        className="relative flex justify-center w-full px-4 py-2 text-sm font-medium text-white bg-green-500 border border-transparent rounded-md group btn-theme hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500"
                    >
                      เข้าสู่ระบบ
                    </Link>
                  </div>
                </div>
              </form>
            </div>
          </div>
        </div>
      );
}