import "./Navbar.css";
import React, { useState } from "react";
import Applogo from "./LogoBanner.png";
import Plussq from "./plus-square.svg";
import User from "./user.svg";
import Search from "./search.svg";
import Backpack from "./shopping-bag.svg";
import More from "./more.svg";
import { LoginUser } from "../../../models/User";
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
              "lg:flex flex-grow items-center" +
              (navbarOpen ? " flex" : " hidden")
            }
            id="example-navbar-danger"
          >
            <ul className="flex flex-col list-none lg:flex-row lg:ml-auto">
            <li className="nav-item">
                <a
                  className="flex items-center px-3 py-2 text-xs font-bold leading-snug text-green-600 uppercase hover:opacity-75"
                  href="/login"
                >
                  Login
                </a>
              </li>
              <li className="nav-item">
                <a
                  className="flex items-center px-3 py-2 text-xs font-bold leading-snug text-green-600 uppercase hover:opacity-75"
                  href="#"
                >
                  <img src={Search} alt="Search" />
                </a>
              </li>
              <li className="nav-item">
                <a
                  className="flex items-center px-3 py-2 text-xs font-bold leading-snug text-green-600 uppercase hover:opacity-75"
                  href="/review"
                >
                  <img src={Plussq} alt="Plussq" />
                </a>
              </li>
              <li className="nav-item">
                <a
                  className="flex items-center px-3 py-2 text-xs font-bold leading-snug text-green-600 uppercase hover:opacity-75"
                  href={`/order/${LoginUser.username}`}
                >
                  <img src={Backpack} alt="Backpack" />
                </a>
              </li>
              <li className="nav-item">
                <a
                  className="flex items-center px-3 py-2 text-xs font-bold leading-snug text-green-600 uppercase hover:opacity-75"
                  href="/profile"
                >
                  <button type="button" onClick={logout}>
                    <img src={User} alt="User" />
                  </button>
                </a>
              </li>
            </ul>
          </div>
        </div>
      </nav>
    </>
  );
}
