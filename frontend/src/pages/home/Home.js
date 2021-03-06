import React, { useState, useEffect } from "react";
import "./index.css";
import Navbar from "../../asset/include/navbar/Navbar";
import { Link } from "react-router-dom";
import NumberFormat from "react-number-format";
import { getHome } from "../service/home/getHome";

const items = getHome();

function productShow(product){
  console.log("logs ", product)
  if(product.status === "Available"){
    return(
      <>
        <div key={product.id} className="relative group">
          <div className="w-full overflow-hidden bg-gray-200 rounded-md min-h-80 aspect-w-1 aspect-h-1 group-hover:opacity-75 lg:h-80 lg:aspect-none">
            <Link to={`/products/${product.id}`}>
              <img
                src={`http://spaceship.trueddns.com:23720/api/v1/products/images/${product.image_id}`}
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
            </div>
            <p className="text-sm font-medium text-gray-900">
              <NumberFormat
                value={product.price}
                displayType={"text"}
                thousandSeparator={true}
                prefix={"฿"}
              />
            </p>
          </div>
        </div>
      </>
    )
  }

}

export default function Home() {
  const [item, setItem] = useState([]);

  useEffect(() => {
    if (items) {
      items.then(function (data) {
        setItem(data.data);
        // console.log("name:", data.name);
      });
    }
  });
  if (!item) return null;

  return (
    <div className="bg-white">
      <Navbar />
      <div className="max-w-2xl px-4 py-16 mx-auto sm:py-24 sm:px-6 lg:max-w-7xl lg:px-8 font-body">
        <div className="grid grid-cols-1 mt-6 gap-y-10 gap-x-6 sm:grid-cols-2 lg:grid-cols-4 xl:gap-x-8">
          {item.map((product) => (
          <>
            {productShow(product)}
          </>
          ))}
        </div>
      </div>
    </div>
  );
}
