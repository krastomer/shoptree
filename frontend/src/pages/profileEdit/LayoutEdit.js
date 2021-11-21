import react, { useRef, useState } from "react";
import Add from "../review/add.svg";
import { useForm } from "react-hook-form";
import Navbar from "../../asset/include/navbar/Navbar";
import EditAddress from "./Editaddress";
export default function LayoutEdit() {
  return (
    <>
      <Navbar />
      <EditAddress />
    </>
  );
}
