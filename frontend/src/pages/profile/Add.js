import react, { useRef, useState } from "react";
import Add from "../review/add.svg";
import { useForm } from "react-hook-form";
import allLocation from "./allLocation";
export default function AddAddress() {
  return (
    <>
      <button
        className="text-black font-body border-dashed md:border-dashed border-4"
        type="button"
      >
        <a href="/profile/edit">
        <div className="p-5 flex flex-col items-center font-bold leading-snug  font-theme ">
          เพิ่มที่อยู่
          <div className="p-2 flex">
            <img src={Add} alt="Add" />
          </div>
        </div>
        </a>
      </button>
      
    </>
  );
}
