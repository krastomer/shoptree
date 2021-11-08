import React, { useState, Fragment, useEffect } from "react";
import Navbar from "../../asset/include/navbar/Navbar"
import Timebar from "./Timebar"
import Statusbar from "./Statusbar"
import Processbar from "./Processbar"
import allOrder from "./allOrder";
import { Link } from "react-router-dom";
import NumberFormat from 'react-number-format';

const products = allOrder

export default function Order() {
  
  return (
    <div className="bg-white">
    <Navbar />
    <Timebar />
    
    <div className="max-w-2xl mx-auto px-4  sm:px-6 lg:max-w-7xl lg:px-8 font-body">
    <Statusbar />
    <h2 className="text-2xl py-4 text-gray-600 tracking-tight">สินค้าในตะกร้า</h2>
      <div className="mt-3 grid grid-cols-1 gap-y-10 gap-x-6 sm:grid-cols-2 lg:grid-cols-4 xl:gap-x-8">
        {products.map((product) => (
          <div key={product.id} className="group relative">
            <div className="w-full min-h-80 bg-gray-200 aspect-w-1 aspect-h-1 rounded-md overflow-hidden  lg:h-80 lg:aspect-none">
            <Link to={`/products/${product.id}`}>
              <img
                src={product.imageSrc}
                alt={product.imageAlt}
                className="w-full h-full object-center object-cover lg:w-full lg:h-full"
              />
              </Link>
            </div>
            <div className="mt-4 flex flex-col">
              <div className = "flex flex-row justify-between">
                <h3 className="text-sm text-gray-700 text-right">
                  {product.id}
                </h3>
                <h3 className="text-sm text-gray-700 text-right">
                  {product.name}
                </h3>
              </div>
              <div>
              <div className = "flex flex-row justify-between">
                <p className="text-sm text-right font-medium text-gray-900">ราคา</p>
                <p className="text-sm text-right font-medium text-gray-900">
                  <NumberFormat value={product.price} displayType={'text'} thousandSeparator={true} prefix={'$'} /> 
                </p>
              </div>
              <div className = "flex flex-row justify-center">
                <div className="">
                  <button className="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded-full">
                    ลบสินค้านี้ออกจากตะกร้า 
                  </button>
                </div>
              </div>
              </div>
            </div>
          </div>
        ))}
      </div>
      <Processbar />
    </div>
  </div>
  )
}