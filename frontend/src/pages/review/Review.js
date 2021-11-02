import "./Review.css";
import ItemCard from "./ItemCard";
import React, { useEffect, useState } from "react";
import Navbar from "../../asset/include/navbar/Navbar";
import Box from '@mui/material/Box';
import Modal from '@mui/material/Modal';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';
import allReview from "./AllReview";

const reviews = allReview

export default function Review() {
  const [checkReview, setReview] = useState(null);
  
  return (
    <div className="bg-white">
    <Navbar />
    <div className="bg-white font-body">
      <div className="max-w-2xl mx-auto py-16 px-4 sm:py-12 sm:px-6 lg:max-w-7xl lg:px-8">
        <h2 className="text-4xl font-theme tracking-tight">รีวิวสินค้าจากทางเรา</h2>

        <div className="mt-6 grid grid-cols-1 gap-y-10 gap-x-6 sm:grid-cols-2 lg:grid-cols-4 xl:gap-x-8">
          {reviews.map((review) => (
            <div key={review.id} className="group relative bg-theme">
              <div className="w-full aspect-w-1 aspect-h-1 rounded-md overflow-hidden group-hover:opacity-75  lg:aspect-none">
              <span aria-hidden="true" className="absolute inset-0" />
                <ItemCard point={review.point} />      
              </div>
              <div className="py-2 px-4 mt-4 flex justify-between">
                <p className="text-white">{review.desc}</p>
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  </div>
  );
}
