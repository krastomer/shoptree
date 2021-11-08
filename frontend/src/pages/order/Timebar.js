import React, { useState } from "react";
import "./Order.css";
// import { MdShoppingCart,MdPerson } from "react-icons/md";

export default function Timebar() {
  const [navbarOpen, setNavbarOpen] = React.useState(false);

  return (
    <>
      <nav className="relative flex flex-wrap items-center justify-between px-2 py-3 bg-theme mb-3 font-body">
        <div className="container px-4 mx-auto flex flex-wrap items-center justify-between">
          <div className="w-auto static block justify-start">
            <a
              className="text-sm font-bold leading-relaxed inline-block mr-4 py-2 whitespace-nowrap uppercase text-white"
              href="/home"
            >
                <font className ="font-normal">คุณมีเวลาคงเหลือ&nbsp;</font>
                <font className ="text-xl font-semibold">
                    10
                </font>
                <font className ="font-normal">&nbsp;นาที&nbsp;</font>
                <font className ="text-xl font-semibold">
                    30
                </font>
                <font className ="font-normal">&nbsp;วินาที</font>
            </a>
            <button
              className="text-white cursor-pointer text-xl leading-none px-3 py-1 border border-solid border-transparent rounded bg-transparent block lg:hidden outline-none focus:outline-none"
              type="button"
              onClick={() => setNavbarOpen(!navbarOpen)}
            >
            
            </button>
          </div>
          <div>
            <ul className="flex flex-row list-none lg:ml-auto">
              <li className="nav-item">
                <a
                  className="px-3 py-2 flex items-center text-xs uppercase font-bold leading-snug text-white hover:opacity-75"
                  href="#"
                >
                <font className="font-medium text-base">เรียนรู้เพิ่มเติม</font>
                </a>
              </li>
            </ul>
          </div>
        </div>
      </nav>
    </>
  );
}

