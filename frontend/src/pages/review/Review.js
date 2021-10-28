import "./Review.css";
import ItemCard from "./ItemCard"
import React, { useState } from "react";

export default function Review() {
  return (
        <div className="font-body container mx-auto">
          <div className=" text-yellow-800">
            <h3 className="mt-6 text-left text-3xl font-extrabold">รีวิวสินค้าจากทางเรา</h3>
          </div>
          <div className="grid grid-cols-3 gap-4">
              <ItemCard 
                point="3" 
                detail ="ต้นไม้สวยงาม มีกลื่นหอม ตรงตามรูป จัดส่งได้เร็วทันใจวัยรุ่น"
              />
              <ItemCard 
                point="4" 
                detail ="ไอซับใจล่มๆ"
              />
              <ItemCard 
                point="2" 
                detail ="แก้มน้องนางนั้นเขียวกว่าใคร"
              />
              <ItemCard 
                point="2" 
                detail ="หล่อมากกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกกก"
              />
              <ItemCard 
                point="0" 
                detail ="ไอซับไอสัสใจเย็นๆ"
              />
          </div>
        </div>
  )
}

