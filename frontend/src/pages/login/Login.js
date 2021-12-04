import "./Login.css";
import Applogo from "../../logo.svg";
import React, { useState, useRef } from "react";
import { useDispatch, useSelector } from "react-redux";
import { Redirect } from "react-router-dom";
import Form from "react-validation/build/form";
import Input from "react-validation/build/input";
import CheckButton from "react-validation/build/button";
import { login } from "../../actions/auth";
export default function Login(props) {
  const required = (value) => {
    if (!value) {
      return (
        <div className="alert alert-danger" role="alert">
          This field is required!
        </div>
      );
    }
  };
  const form = useRef();
  const checkBtn = useRef();
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [loading, setLoading] = useState(false);

  const { isLoggedIn } = useSelector((state) => state.auth);
  const { message } = useSelector((state) => state.message);

  const dispatch = useDispatch();

  const onChangeUsername = (e) => {
    const username = e.target.value;
    setUsername(username);
  };

  const onChangePassword = (e) => {
    const password = e.target.value;
    setPassword(password);
  };

  const handleLogin = (e) => {
    e.preventDefault();

    setLoading(true);
    form.current.validateAll();
    if (checkBtn.current.context._errors.length === 0) {
      dispatch(login(username, password))
        .then(() => {
          props.history.push("/");
          window.location.reload();
        })
        .catch(() => {
          setLoading(false);
        });
    } else {
      setLoading(false);
    }
  };

  if (isLoggedIn) {
    return <Redirect to="/" />;
  }
  return (
    <div className="grid  md:grid-cols-2 h-screen font-prompt font-body ">
      <div
        className="md:flex hidden flex bg-primary bg-green-600 flex-col "
        style={{ boxShadow: "0 4px 4px #000" }}
      >
        <a href="/">
          <div className="my-10 text-white">
            &nbsp;&nbsp;&nbsp;&nbsp;กลับสู่หน้าหลัก
          </div>
        </a>
        <div className="mx-auto my-auto ">
          <img src={Applogo} alt="Logo" />
        </div>
      </div>
      <div className="min-h-full flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
        <div className="max-w-md w-full space-y-8">
          <div>
            <h2 className="mt-6 text-left text-3xl font-extrabold text-gray-900">
              ล็อกอิน
            </h2>
          </div>
          <Form  onSubmit={handleLogin} ref={form}>
            <input type="hidden" name="remember" defaultValue="true" />
            <div className="rounded-md shadow-sm -space-y-px">
              <div>
                <p className=" text-gray-500">อีเมล</p>
                <label htmlFor="email-address" className="sr-only">
                  Email address
                </label>
                <Input
                  type="text"
                  className="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-b-md focus:outline-none focus:ring-green-500 focus:border-green-500 focus:z-10 sm:text-sm"
                  name="username"
                  value={username}
                  onChange={onChangeUsername}
                  validations={[required]}
                />
              </div>
              <div>
                <p className=" text-gray-500">รหัสผ่าน</p>
                <label htmlFor="password" className="sr-only">
                  Password
                </label>
                <Input
                  type="password"
                  className="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-t-md focus:outline-none focus:ring-green-500 focus:border-green-500 focus:z-10 sm:text-sm"
                  name="password"
                  value={password}
                  onChange={onChangePassword}
                  validations={[required]}
                />
              </div>
            </div>
            <div className="flex items-center justify-between">
              <div className="flex items-center"></div>
              <div>
                <button
                  onClick={handleLogin}
                  className="group relative w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white btn-theme hover:bg-yellow-00 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500"
                  disabled={loading}
                >
                  {loading && (
                    <span className="spinner-border spinner-border-sm"></span>
                  )}
                  เข้าสู่ระบบ
                </button>
              </div>
            </div>
            <div>
              <div className=" text-center">
                <a
                  href="register"
                  className="font-medium text-gray-600 hover:text-gray-400"
                >
                  ถ้ายังไม่มีบัญชี คลิ๊กเพื่อสมัครสมาชิก
                </a>
              </div>
            </div>
            {message && (
              <div className="form-group">
                <div className="alert alert-danger" role="alert">
                  {message}
                </div>
              </div>
            )}
            <CheckButton style={{ display: "none" }} ref={checkBtn} />
          </Form>
        </div>
      </div>
    </div>
  );
}
