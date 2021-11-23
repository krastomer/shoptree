import "./Navbar.css";
import React, { useEffect, useState } from "react";
import Applogo from "./LogoBanner.png";
import Plussq from "./plus-square.svg";
import User from "./user.svg";
import Search from "./search.svg";
import Backpack from "./shopping-bag.svg";
import More from "./more.svg";
import UserDropdown from "./UserDropdown";
import { LoginUser } from "../../../models/User";
import { useForm } from "react-hook-form";

// import { MdShoppingCart,MdPerson } from "react-icons/md";
const logout = () => {
  console.log("logout");
  console.log(LoginUser);
  if (LoginUser.auth.loggedIn) {
    localStorage.removeItem("user");
    window.location.reload();
  }
};
export default function Navbar() {
  const [navbarOpen, setNavbarOpen] = useState(false);
  const {
    register,
  } = useForm({});
  return (
    <>
      <nav className="relative flex flex-wrap items-center justify-between px-2 py-3 mb-3 bg-white">
        <div className="container flex flex-wrap items-center justify-between px-4 mx-auto">
          <div className="relative flex justify-between w-full lg:w-auto lg:static lg:block lg:justify-start">
            <a
              className="inline-block py-2 mr-4 text-sm font-bold leading-relaxed text-green-600 uppercase whitespace-nowrap"
              href="/"
            >
              <img src={Applogo} alt="Applogo" />
            </a>
            <button
              className="block px-3 py-1 text-xl leading-none text-white bg-transparent border border-transparent border-solid rounded outline-none cursor-pointer lg:hidden focus:outline-none"
              type="button"
              onClick={() => setNavbarOpen(!navbarOpen)}
            >
              <img src={More} alt="Plussq" />
            </button>
          </div>
          <div
            className={
              "lg:flex justify-end flex-grow items-center" +
              (navbarOpen ? " flex" : " hidden")
            }
            id="example-navbar-danger"
          >
            <ul className="flex flex-col text-right list-none justify-items-end lg:flex-row lg:ml-auto">
              <li className="nav-item">
                <div class="flex items-center justify-center">
                  <div class="flex border-5 rounded">
                    <select {...register("Title", { required: true })}>
                      <option value="Mr">Mr</option>
                      <option value="Mrs">Mrs</option>
                      <option value="Miss">Miss</option>
                      <option value="Dr">Dr</option>
                    </select>
                    <input
                      type="text"
                      
                      class="px-4 py-2 w-80"
                      placeholder="โสดและเหงามาก"
                    />
                    <button class="flex items-center justify-center px-4 border-l">
                      <img src={Search} alt="Search" />
                    </button>
                  </div>
                </div>
              </li>
              <li className="nav-item">
                <a
                  className="flex px-3 py-2 text-xs font-bold leading-snug text-right text-green-600 uppercase hover:opacity-75"
                  href="/login"
                >
                  Login
                </a>
              </li>

              <li className="nav-item">
                <a
                  className="flex px-3 py-2 text-xs font-bold leading-snug text-right text-green-600 uppercase hover:opacity-75"
                  href="/review"
                >
                  <img src={Plussq} alt="Plussq" />
                </a>
              </li>
              <li className="nav-item">
                <a
                  className="flex px-3 py-2 text-xs font-bold leading-snug text-right text-green-600 uppercase hover:opacity-75"
                  href={`/order/${LoginUser.username}`}
                >
                  <img src={Backpack} alt="Backpack" />
                </a>
              </li>
              <li className="flex text-right nav-item">
                  <UserDropdown />
              </li>
            </ul>
          </div>
        </div>
      </nav>
    </>
  );
}
