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
const logout = ()=>{
  console.log("logout")
  console.log(LoginUser)
  if(LoginUser.auth.loggedIn){
    LoginUser.auth.loggedIn = false
    alert("ออกจากระบบละนะอ้วง")
  }
}
export default function Navbar() {
  const [navbarOpen, setNavbarOpen] = useState(false);
  
  return (
    <>
      <nav className="relative flex flex-wrap items-center justify-between px-2 py-3 bg-white mb-3">
        <div className="container px-4 mx-auto flex flex-wrap items-center justify-between">
          <div className="w-full relative flex justify-between lg:w-auto lg:static lg:block lg:justify-start">
            <a
              className="text-sm font-bold leading-relaxed inline-block mr-4 py-2 whitespace-nowrap uppercase text-green-600"
              href="/"
            >
              <img src={Applogo} alt="Applogo" />
            </a>
            <button
              className="text-white cursor-pointer text-xl leading-none px-3 py-1 border border-solid border-transparent rounded bg-transparent block lg:hidden outline-none focus:outline-none"
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
             
            <ul className="flex flex-col lg:flex-row list-none lg:ml-auto">
            <li className="nav-item">
                <a
                  className="px-3 py-2 flex items-center text-xs uppercase font-bold leading-snug  text-green-600 hover:opacity-75"
                  href="#"
                >
                   <img src={Search} alt="Search" />
                </a>
              </li>
              <li className="nav-item">
                <a
                  className="px-3 py-2 flex items-center text-xs uppercase font-bold leading-snug  text-green-600 hover:opacity-75"
                  href="/review"
                >
                  <img src={Plussq} alt="Plussq" />
                </a>
              </li>
              <li className="nav-item">
                <a
                  className="px-3 py-2 flex items-center text-xs uppercase font-bold leading-snug text-green-600 hover:opacity-75"
                  href="/order"
                >
                  <img src={Backpack} alt="Backpack" />
                </a>
              </li>
              <li className="nav-item">
                <a
                  className="px-3 py-2 flex items-center text-xs uppercase font-bold leading-snug text-green-600 hover:opacity-75"
                  href="#"
                >
                  <button
                    type="button"
                    onClick={logout}
                  >
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

