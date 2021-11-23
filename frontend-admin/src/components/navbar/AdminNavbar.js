import React, { useState } from "react";
import "./AdminNavbar.css";
import UserDropdown from "../carditems/UserDropdown.js";
// import { MdShoppingCart,MdPerson } from "react-icons/md";

export default function AdminNavbar() {
  const [navbarOpen, setNavbarOpen] = React.useState(false);

  return (
    <>
      {/* Navbar */}
      <nav className="absolute top-0 left-0 w-full z-10 bg-transparent md:flex-row md:flex-nowrap md:justify-start flex items-center p-4 bg-theme">
        <div className="w-full mx-autp items-center flex justify-between md:flex-nowrap flex-wrap md:px-10 px-5 ">
          {/* Brand */}
          <a
            className="text-white text-sm uppercase hidden lg:inline-block font-semibold"
            href="#pablo"
            onClick={(e) => e.preventDefault()}
          >
            หน้าหลัก
          </a>
          {/* User */}
          <UserDropdown />
        </div>
      </nav>
      {/* End Navbar */}
    </>
  );
}

