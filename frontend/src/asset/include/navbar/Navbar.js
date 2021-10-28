import "./Navbar.css";
import React, { useState } from "react";
import Applogo from "../../LogoBanner.png";

export default function Navbar() {
  return (
        <nav className ="bg-white border-indigo-500 px-2 md:px-4 lg:px-5">
          <div className="container mx-auto flex flex-wrap items-center justify-between">
            <a href="#" className="flex w-1/4">
            <img src={Applogo} alt="Logo"/>
            </a>
          </div>
        </nav>
  )
}

