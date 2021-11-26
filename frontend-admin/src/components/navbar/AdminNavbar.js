import React, { useState } from "react";
import "./AdminNavbar.css";
import { Link } from "react-router-dom";

export default function AdminNavbar() {
  const [navbarOpen, setNavbarOpen] = React.useState(false);

  return (
    <>
      <nav className="absolute top-0 left-0 z-10 flex items-center w-full p-4 bg-transparent bg-green-500 md:flex-row md:flex-nowrap md:justify-start font-body">
        <div className="flex flex-wrap items-center justify-between w-full px-5 mx-autp md:flex-nowrap md:px-10 ">
          <Link
            className="hidden text-sm font-semibold text-white uppercase lg:inline-block"
            to="/"
            onClick={(e) => e.preventDefault()}
          >
            หน้าหลัก
          </Link>
        </div>
      </nav>
    </>
  );
}

