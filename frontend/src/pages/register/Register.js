import "./Register.css";
import Applogo from "../../logo.svg";
import React, { useEffect, useState, useRef } from "react";
// import { postRegister } from "../service/auth_service";
import { useHistory } from "react-router";
import { useForm } from "react-hook-form";

export default function Register() {
  let history = useHistory();
  const {
    register,
    formState: { errors },
    handleSubmit,
    watch,
  } = useForm({});
  const password = useRef({});
  password.current = watch("password", "");
  const onSubmit = (data) => {
    // try {
    //   postRegister(data);
    //   history.push("/login");
    // } catch {
    //   alert("error");
    // }
  };
  return (
    <div className="grid  md:grid-cols-2 h-screen font-body ">
      <div
        className="md:flex hidden flex bg-primary bg-green-600 flex-col "
        style={{ boxShadow: "0 4px 4px #000" }}
      >
        <div className="my-10 text-white">
          <a href="/">&nbsp;&nbsp;&nbsp;&nbsp;กลับสู่หน้าหลัก</a>
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

          <input type="hidden" name="remember" defaultValue="true" />
          <div className="rounded-md shadow-sm -space-y-px">
            <form onSubmit={(e) => e.preventDefault()}>
              <div>
                <p className=" text-gray-500">ชื่อ-นามสกุล</p>
                <label htmlFor="password" className="sr-only">
                  Name
                </label>
                <input
                  name="Name"
                  type="text"
                  placeholder="ชื่อ - นามสกุล"
                  {...register("name", {
                    required: true,
                    minLength: {
                      value: 1,
                      message: "write your name",
                    },
                    maxLength: 80,
                  })}
                  className="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-t-md focus:outline-none focus:ring-green-500 focus:border-green-500 focus:z-10 sm:text-sm"
                />
                {errors.Name && <p>{errors.Name.message}</p>}
              </div>
              <div>
                <p className=" text-gray-500">อีเมล</p>
                <label htmlFor="email-address" className="sr-only">
                  Email address
                </label>
                <input
                  name="Email"
                  type="text"
                  placeholder="Email"
                  {...register("email", {
                    required: true,
                    pattern: {
                      value:  /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/,
                      message: "Email wrong",
                    },
                  })}
                  className="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-t-md focus:outline-none focus:ring-green-500 focus:border-green-500 focus:z-10 sm:text-sm"
                />
                {errors.Email && <p>{errors.Email.message}</p>}
              </div>
              <div>
                <p className=" text-gray-500">รหัสผ่าน</p>
                <label htmlFor="password" className="sr-only">
                  Password
                </label>
                <input
                  name="password"
                  type="password"
                  placeholder="Password"
                  {...register("password", {
                    required: true,
                    minLength: {
                      value: 8,
                      message: "Password must have at least 8 characters",
                    },
                    pattern:{
                      value :/(?=.*\d)(?=.*[A-Z])/,
                      message: "minimum 1 uppercase and minimum 1 number"
                    }
                  })}
                  className="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-b-md focus:outline-none focus:ring-green-500 focus:border-green-500 focus:z-10 sm:text-sm"
                />
                {errors.password && <p>{errors.password.message}</p>}
              </div>
              <div>
                <p className=" text-gray-500">ยืนยันรหัสผ่าน</p>
                <label htmlFor="password" className="sr-only">
                  RePassword
                </label>
                <input
                  name="password_repeat"
                  type="password"
                  placeholder="Confirm password"
                  {...register("password_repeat", {
                    validate: (value) =>
                      value === password.current ||
                      "The passwords do not match",
                  })}
                  className="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-b-md focus:outline-none focus:ring-green-500 focus:border-green-500 focus:z-10 sm:text-sm"
                />
                {errors.password_repeat && (
                  <p>{errors.password_repeat.message}</p>
                )}
              </div>
              <div>
                <p className=" text-gray-500">เบอร์โทรศัพท์</p>
                <label htmlFor="password" className="sr-only">
                  Phone
                </label>
                <input
                  type="tel"
                  placeholder="Mobile number"
                  {...register("mobile_number", {
                    required: true,
                    pattern: {
                      value:/(?=.*\d)/ ,
                      message : "must phone number",
                    },
                    minLength: {
                      value: 8,
                      message: "Phone number must have 8 -12",
                    },
                    maxLength: {
                      value: 12,
                      message: "Phone number must have 8 -12",
                    },
                  })}
                  className="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-b-md focus:outline-none focus:ring-green-500 focus:border-green-500 focus:z-10 sm:text-sm"
                />
                {errors.Mobile_number && <p>{errors.Mobile_number.message}</p>}
              </div>
              <div>
                <div className="flex items-center justify-between">
                  <div className="flex items-center">
                    <a
                      href="#"
                      className="font-medium text-gray-600 hover:text-gray-400"
                    >
                      ฉันยอมรับข้อตกลง เพิ่มเติม
                    </a>
                  </div>
                </div>
                <br></br>
              </div>
              <input
                type="submit"
                value="สมัครสมาชิก"
                onClick={handleSubmit(onSubmit)}
                className="group relative w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white btn-theme hover:bg-yellow-00 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500"
              ></input>
            </form>
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
        </div>
      </div>
    </div>
  );
}
