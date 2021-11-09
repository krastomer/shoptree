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
      <div className="max-w-2xl mx-auto py-16 px-4 sm:py-12 sm:px-6 lg:max-w-7xl lg:px-8">
        <h2 className="text-4xl font-theme tracking-tight">รีวิวสินค้าจากทางเรา</h2>

        <div className="mt-6 grid grid-cols-1 gap-y-10 gap-x-6 sm:grid-cols-2 lg:grid-cols-4 xl:gap-x-8">
        <AddModal onAddReview = {onAddNewReview}/>
          {reviews.map((element) => (
            <div key={element.id} className="group relative bgg-theme">
              <div className="w-full aspect-w-1 aspect-h-1 rounded-md overflow-hidden group-hover:opacity-75  lg:aspect-none">
              <span aria-hidden="true" className="absolute inset-0" />
                <ItemCard point={element.point} />      
              </div>
              <div className="py-2 px-4 mt-4 flex justify-between">
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
