import React, { useState, useEffect } from "react";
import "./index.css";
import Navbar from "../../asset/include/navbar/Navbar"
import allProduct from "./allProduct";
import { Link } from "react-router-dom";
import NumberFormat from 'react-number-format';

const products = allProduct

export default function Home() {
  return (
    <div className="bg-white">
    <Navbar />
    <div className="max-w-2xl mx-auto py-16 px-4 sm:py-24 sm:px-6 lg:max-w-7xl lg:px-8 font-body">
      <div className="mt-6 grid grid-cols-1 gap-y-10 gap-x-6 sm:grid-cols-2 lg:grid-cols-4 xl:gap-x-8">
        {products.map((product) => (
          <div key={product.id} className="group relative">
            <div className="w-full min-h-80 bg-gray-200 aspect-w-1 aspect-h-1 rounded-md overflow-hidden group-hover:opacity-75 lg:h-80 lg:aspect-none">
            <Link to={`/products/${product.id}`}>
              <img
                src={product.imageSrc}
                alt={product.imageAlt}
                className="w-full h-full object-center object-cover lg:w-full lg:h-full"
              />
              </Link>
            </div>
            <div className="mt-4 flex justify-between">
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