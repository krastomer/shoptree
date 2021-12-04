import React, { useState, useEffect } from "react";
import "./index.css";
import Navbar from "../../asset/include/navbar/Navbar"
import allProduct from "./allProduct";
import { Link } from "react-router-dom";
import { LoginUser } from "../../models/User";
import NumberFormat from 'react-number-format';
// import { product } from "../service/service";
const products = allProduct

export default function Home() {
  return (
    <div className="bg-white">
    <Navbar />
    <div className="max-w-2xl px-4 py-16 mx-auto sm:py-24 sm:px-6 lg:max-w-7xl lg:px-8 font-body">
      <div className="grid grid-cols-1 mt-6 gap-y-10 gap-x-6 sm:grid-cols-2 lg:grid-cols-4 xl:gap-x-8">
        {products.map((product) => (
          <div key={product.id} className="relative group">
            <div className="w-full overflow-hidden bg-gray-200 rounded-md min-h-80 aspect-w-1 aspect-h-1 group-hover:opacity-75 lg:h-80 lg:aspect-none">
            <Link to={`/products/${product.id}`} >
              <img
                src={product.imageSrc}
                alt={product.imageAlt}
                className="object-cover object-center w-full h-full lg:w-full lg:h-full"
              />
              <info></info>
              </Link>
            </div>
            <div className="flex justify-between mt-4">
              <div>
                <h3 className="text-sm text-gray-700">
                 <Link to={`/products/${product.id}`}>{product.name}</Link>
                </h3>
                <p className="mt-1 text-sm">
                  <font className ="text-red-700"> {product.status}</font>
                  <font> </font>
                  <font className ="text-gray-600">{product.unittime}</font>
                </p>
              </div>
              <p className="text-sm font-medium text-gray-900">
                <NumberFormat value={product.price} displayType={'text'} thousandSeparator={true} prefix={'$'} />   
              </p>
            </div>
          </div>
        ))}
      </div>
    </div>
  </div>
  )
}