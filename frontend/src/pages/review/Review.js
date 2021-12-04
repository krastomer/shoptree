import "./Review.css";
import ItemCard from "./ItemCard";
import React, { useEffect, useState } from "react";
import Navbar from "../../asset/include/navbar/Navbar";
import Box from '@mui/material/Box';
import Modal from '@mui/material/Modal';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';
import allReview from "./AllReview";
import AddModal from "./Modal";


const initReviews = allReview

export default function Review() {
  // const [checkReview, setReview] = useState(null);
  const onAddNewReview = (newReview) =>{
    console.log("ข้อมูลมาจาก Form = ", newReview)
    setReview((prevReview) => {
      return[newReview,...prevReview]
    })
  }
  const [reviews, setReview] = useState(initReviews)

  return (
    <div className="bg-white">
    <Navbar />
    <div className="bg-white font-body">
      <div className="max-w-2xl px-4 py-16 mx-auto sm:py-12 sm:px-6 lg:max-w-7xl lg:px-8">
        <h2 className="text-4xl tracking-tight font-theme">รีวิวสินค้าจากทางเรา</h2>

        <div className="grid grid-cols-1 mt-6 gap-y-10 gap-x-6 sm:grid-cols-2 lg:grid-cols-4 xl:gap-x-8 ">
        <AddModal onAddReview = {onAddNewReview}/>
          {reviews.map((element) => (
            <div key={element.id} className="relative group bgg-theme">
              <div className="w-full overflow-hidden rounded-md aspect-w-1 aspect-h-1 group-hover:opacity-75 lg:aspect-none">
              <span aria-hidden="true" className="absolute inset-0" />
                <ItemCard point={element.point} />      
              </div>
              <div className="flex justify-between px-4 py-2 mt-4">
                <p className="text-white">{element.desc}</p>
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  </div>
  );
}
