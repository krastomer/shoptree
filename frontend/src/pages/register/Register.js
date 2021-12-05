import "./Register.css";
import Applogo from "../../logo.svg";
import React, { useEffect, useState, useRef } from "react";
// import { postRegister } from "../service/auth_service";
import { register_ } from "../../actions/auth";
import { useDispatch, useSelector } from "react-redux";
import { computeStyles } from "@popperjs/core";
import { useHistory } from "react-router-dom";
import { isEmail } from "validator";
import Form from "react-validation/build/form";
import Input from "react-validation/build/input";
import CheckButton from "react-validation/build/button";

export default function Register() {
  let history = useHistory();
  const form = useRef();
  const checkBtn = useRef();
  const [successful, setSuccessful] = useState(false);
  const { message } = useSelector((state) => state.message);
  const dispatch = useDispatch();
  const [username, setUsername] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState();
  const [repeat, setRepeat] = useState();
  const [mobile, setMobile] = useState();
  const required = (value) => {
    if (!value) {
      return (
        <div className="alert alert-danger" role="alert">
          This field is required!
        </div>
      );
    }
  };

  const validEmail = (value) => {
    if (!isEmail(value)) {
      return (
        <div className="alert alert-danger" role="alert">
          This is not a valid email.
        </div>
      );
    }
  };

  const vusername = (value) => {
    if (value.length < 3 || value.length > 40) {
      return (
        <div className="alert alert-danger" role="alert">
          The username must be between 3 and 20 characters.
        </div>
      );
    }
  };

  const vpassword = (value) => {
    if (value.length < 6 || value.length > 40|| value.TOstring) {
      return (
        <div className="alert alert-danger" role="alert">
          The password must be between 6 and 40 number.
        </div>
      );
    }
  };
  // const vrepeat = (value) => {
  //   if (value != password) {
  //     return (
  //       <div className="alert alert-danger" role="alert">
  //         don't format with password
  //       </div>
  //     );
  //   }
  // };
  const vmobile = (value) => {
    if (value.length != 10) {
      return (
        <div className="alert alert-danger" role="alert">
          mobile must have 10 number
        </div>
      );
    }
  };
  const onChangeUsername = (e) => {
    const username = e.target.value;
    setUsername(username);
  };

  const onChangeEmail = (e) => {
    const email = e.target.value;
    setEmail(email);
  };

  const onChangePassword = (e) => {
    const password = e.target.value;
    setPassword(password);
  };
  // const onChangeRepeat = (e) => {
  //   const repeat = e.target.value;
  //   setRepeat(repeat);
  // };
  const onChangeMobile = (e) => {
    const mobile = e.target.value;
    setMobile(mobile);
  };
  const handleRegister = (e) => {
    e.preventDefault();
    setSuccessful(false);
    form.current.validateAll();
    if (checkBtn.current.context._errors.length === 0) {
      dispatch(register_(email, username, password, mobile))
        .then(() => {
          setSuccessful(true);
          history.push("/login");
        })
        .catch(() => {
          setSuccessful(false);
        });
    }
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
            <Form onSubmit={handleRegister} ref={form}>
              <div>
                <p className=" text-gray-500">ชื่อ-นามสกุล</p>
                <label htmlFor="password" className="sr-only">
                  Name
                </label>
                <Input
                  type="text"
                  name="name"
                  placeholder="ชื่อ - นามสกุล"
                  value={username}
                  onChange={onChangeUsername}
                  validations={[required, vusername]}
                  className="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-t-md focus:outline-none focus:ring-green-500 focus:border-green-500 focus:z-10 sm:text-sm"
                />
              </div>
              <div>
                <p className=" text-gray-500">อีเมล</p>
                <label htmlFor="email-address" className="sr-only">
                  Email address
                </label>
                <Input
                  type="text"
                  placeholder="Email"
                  name="email"
                  value={email}
                  onChange={onChangeEmail}
                  validations={[required, validEmail]}
                  className="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-t-md focus:outline-none focus:ring-green-500 focus:border-green-500 focus:z-10 sm:text-sm"
                />
              </div>
              <div>
                <p className=" text-gray-500">รหัสผ่าน</p>
                <label htmlFor="password" className="sr-only">
                  Password
                </label>
                <label htmlFor="password">Password</label>
                <Input
                  type="password"
                  name="password"
                  placeholder="Password"
                  value={password}
                  onChange={onChangePassword}
                  validations={[required, vpassword]}
                  className="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-b-md focus:outline-none focus:ring-green-500 focus:border-green-500 focus:z-10 sm:text-sm"
                />
              </div>
              {/* <div>
                <p className=" text-gray-500">ยืนยันรหัสผ่าน</p>
                <label htmlFor="password" className="sr-only">
                  RePassword
                </label>
                <label htmlFor="password">Password</label>
                <Input
                  type="password"
                  name="repeat"
                  placeholder="Repeat-password"
                  value={repeat}
                  onChange={onChangeRepeat}
                  validations={[required, vrepeat]}
                  className="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-b-md focus:outline-none focus:ring-green-500 focus:border-green-500 focus:z-10 sm:text-sm"
                />
              </div> */}
              <div>
                <p className=" text-gray-500">เบอร์โทรศัพท์</p>
                <label htmlFor="password" className="sr-only">
                  Phone
                </label>
                <Input
                  type="tel"
                  name="mobile"
                  placeholder="Mobile"
                  value={mobile}
                  onChange={onChangeMobile}
                  validations={[required, vmobile]}
                  className="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-b-md focus:outline-none focus:ring-green-500 focus:border-green-500 focus:z-10 sm:text-sm"
                />
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
              <button className="group relative w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white btn-theme hover:bg-yellow-00 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500">
                สมัครสมาชิก
                </button>
              <CheckButton style={{ display: "none" }} ref={checkBtn} />
            </Form>
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
