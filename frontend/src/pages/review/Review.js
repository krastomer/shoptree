import "./Review.css";
import ItemCard from "./ItemCard";
import React, { useEffect, useState } from "react";
import Navbar from "../../asset/include/navbar/Navbar";
import Box from '@mui/material/Box';
import Modal from '@mui/material/Modal';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';

export default function Review() {
  const [checkReview, setReview] = useState(null);
  
  return (
    <div>
      <Navbar />
      <div className="font-body container mx-auto">
        <div className=" text-yellow-800">
          <h3 className="mt-6 text-left text-3xl font-extrabold">
            รีวิวสินค้าจากทางเรา
          </h3>
        </div>
        <div className="grid grid-cols-3 gap-4">
          <button  className="group relative w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white btn-theme hover:bg-yellow-00 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500">
            เพิ่มรีวิว
          </button>
          <ItemCard
            point="3"
            detail="ต้นไม้สวยงาม มีกลื่นหอม ตรงตามรูป จัดส่งได้เร็วทันใจวัยรุ่น"
          />
          <ItemCard point="4" detail="ไอซับใจล่มๆ" />
          <ItemCard point="2" detail="แก้มน้องนางนั้นเขียวกว่าใคร" />
          <ItemCard
            point="2"
            detail="หล่อมากกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกก"
          />
          <ItemCard point="0" detail="ไอซับไอสัสใจเย็นๆ" />
        </div>
      </div>
    </div>
  );
}
