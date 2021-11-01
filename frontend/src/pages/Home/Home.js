import React, { useState, useRef, useEffect } from "react";
import { Stage, Layer, Rect, Image as KonvaImag } from "react-konva";
import "./index.css";
import Navbar from "../../asset/include/navbar/Navbar"

const products = [
  {
    id: 1,
    name: 'ต้นไม้ประดับ T001',
    href: '#',
    imageSrc: 'https://blog.ghbank.co.th/wp-content/uploads/2021/02/10-air-purification-trees-for-home-01.jpg',
    imageAlt: "Front of men's Basic Tee in black.",
    price: '$3,500.00',
    status: '40:00',
    unittime: 'นาที',
  },
  {
    id: 2,
    name: 'ต้นไม้ประดับ T002',
    href: '#',
    imageSrc: 'https://lh3.googleusercontent.com/proxy/V1xh8rMzABPkhkeHhQOa66r2VaHAOZA7i-tg91kdMwFj7vmCedwZ7Poy_opjd9n0hHaUKT7QUOof-1qKfk8KKip8dEU-SY-dgjXTPKeaJSvpyk-vRHTIEkjjGMJpRdBwgZQ',
    imageAlt: "Front of men's Basic Tee in black.",
    price: '$300.00',
    status: '',
    unittime: '',
  },
  {
    id: 3,
    name: 'ต้นไม้ประดับ T003',
    href: '#',
    imageSrc: 'https://www.jorakay.co.th/ckfinder/userfiles/images/%E0%B8%9E%E0%B8%B2%E0%B8%A2%E0%B8%B8%E0%B8%9D%E0%B8%99%E0%B9%80%E0%B8%AA%E0%B8%B5%E0%B9%88%E0%B8%A2%E0%B8%87%E0%B8%97%E0%B8%B3%E0%B8%95%E0%B9%89%E0%B8%99%E0%B9%84%E0%B8%A1%E0%B9%89%E0%B8%A5%E0%B9%89%E0%B8%A1%E0%B8%97%E0%B8%B1%E0%B8%9A%E0%B8%9A%E0%B9%89%E0%B8%B2%E0%B8%99%20%E0%B8%A1%E0%B8%B5%E0%B8%A7%E0%B8%B4%E0%B8%98%E0%B8%B5%E0%B8%9B%E0%B9%89%E0%B8%AD%E0%B8%87%E0%B8%81%E0%B8%B1%E0%B8%99%E0%B8%AD%E0%B8%A2%E0%B9%88%E0%B8%B2%E0%B8%87%E0%B9%84%E0%B8%A3%20%20%E0%B9%83%E0%B8%99%201.jpg',
    imageAlt: "Front of men's Basic Tee in black.",
    price: '$1,200.00',
    status: '',
    unittime: '',
  },
  {
    id: 4,
    name: 'ต้นไม้ประดับ T004',
    href: '#',
    imageSrc: 'https://s3.theasianparent.com/tap-assets-prod/wp-content/uploads/sites/25/2021/03/%E0%B8%95%E0%B9%89%E0%B8%99%E0%B9%84%E0%B8%A1%E0%B9%891-1.jpg',
    imageAlt: "Front of men's Basic Tee in black.",
    price: '$5,500.00',
    status: '12:30',
    unittime: 'นาที',
  },
  // More products...
]
export default function Review() {
  return (
    <div className="bg-white">
    <Navbar />
    <div className="max-w-2xl mx-auto py-16 px-4 sm:py-24 sm:px-6 lg:max-w-7xl lg:px-8 font-body">
      <div className="mt-6 grid grid-cols-1 gap-y-10 gap-x-6 sm:grid-cols-2 lg:grid-cols-4 xl:gap-x-8">
        {products.map((product) => (
          <div key={product.id} className="group relative">
            <div className="w-full min-h-80 bg-gray-200 aspect-w-1 aspect-h-1 rounded-md overflow-hidden group-hover:opacity-75 lg:h-80 lg:aspect-none">
              <img
                src={product.imageSrc}
                alt={product.imageAlt}
                className="w-full h-full object-center object-cover lg:w-full lg:h-full"
              />
            </div>
            <div className="mt-4 flex justify-between">
              <div>
                <h3 className="text-sm text-gray-700">
                  <a href={product.href}>
                    <span aria-hidden="true" className="absolute inset-0" />
                    {product.name}
                  </a>
                </h3>
                <p className="mt-1 text-sm">
                  <font className ="text-red-700"> {product.status}</font>
                  <font> </font>
                  <font className ="text-gray-600">{product.unittime}</font>
                </p>
              </div>
              <p className="text-sm font-medium text-gray-900">{product.price}</p>
            </div>
          </div>
        ))}
      </div>
    </div>
  </div>
  )
}