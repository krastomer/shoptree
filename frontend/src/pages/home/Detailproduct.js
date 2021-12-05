import React from "react";
import Navbar from "../../asset/include/navbar/Navbar";
import { useParams } from "react-router";
import { useEffect, useState} from "react";
import { getItemByID } from "../service/getItemByID/getItemByID";
import { addItemByID } from "../service/addCart/addCard";
export default function Detail(){
    const {id} = useParams();
    const name = getItemByID(id);
    const [product, setProduct] = useState(null);
    const onAdd = async (e) =>{
      e.preventDefault();
      const addCard = await addItemByID(id)
      alert("เพิ่มสินค้าเรียบร้อย")
    }
    useEffect( ()=>{
      if(product === null){
        name.then(function(data){
            setProduct(data.data);
        })
      }
      },)  
    if(!product) return null;
    return(
        <div className="w-full h-full font-body">
        <Navbar />
        <div className="bg-white font-body">
          <div className="max-w-2xl px-4 py-16 mx-auto rounded-lg sm:py-12 sm:px-6 lg:max-w-7xl lg:px-8">
            <div className="grid justify-between grid-cols-1 mt-2 gap-y-10 gap-x-6 sm:grid-cols-1 lg:grid-cols-2 xl:gap-x-8">
              <div className="flex content-center object-center w-full h-full item-center ">
                <div className="p-6 ">
                  <img
                    src={`http://spaceship.trueddns.com:23720/api/v1/products/images/${product.image_id}`}
                    className="object-contain"
                  />
                </div>
              </div>
  
              <div className="p-6">
                <div className="grid grid-cols-1">
                  <div className="flex justify-between text-3xl font-bold">
                    <p className="font-brown">{product.name}</p>
                    <p className="font-green">{product.price} บาท</p>
                  </div>
                  <p className="text-gray-700">T{id}</p>
                  <div className="flex justify-between mt-6 text-black ">
                    <p className="text-gray-700">ชื่ออังกฤษ :</p>
                    <p className="font-bold uppercase text-l">{product.scientific_name}</p>
                  </div>
                  <p className="mt-10">
                    &nbsp;&nbsp;{product.description}
                  </p>
                </div>
                <div className="pt-10">
                  <button onClick={onAdd} className="relative flex justify-center px-4 py-2 font-medium text-white text-green-500 border border-2 border-green-500 rounded-lg cursor-pointer group hover:bg-yellow-00 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500 hover:border-green-700 hover:text-green-700">
                    เพิ่มลงตะกร้า
                  </button>
                </div>
                <div className="pt-10">
                  <p>สินค้าที่เกี่ยวข้อง</p>
                </div>
                <div className="grid grid-cols-3 pt-3 gap-y-10 gap-x-6">
                  <div>
                    <img
                      src={
                        "https://s3.theasianparent.com/tap-assets-prod/wp-content/uploads/sites/25/2021/03/%E0%B8%95%E0%B9%89%E0%B8%99%E0%B9%84%E0%B8%A1%E0%B9%891-1.jpg"
                      }
                      className="object-contain"
                    />
                   
                  </div>
                  <div>
                    <img
                      src={
                        "https://blog.ghbank.co.th/wp-content/uploads/2021/02/10-air-purification-trees-for-home-01.jpg"
                      }
                      className="object-contain"
                    />
                    {/* <p className = "pt-2 text-sm text-gray-700">T001</p> */}
                  </div>
                  <div>
                    <img
                      src={
                        "https://blog.ghbank.co.th/wp-content/uploads/2021/02/10-air-purification-trees-for-home-01.jpg"
                      }
                      className="object-contain"
                    />
                    {/* <p className = "pt-2 text-sm text-gray-700">T001</p> */}
                  </div>
                </div>
                <div className="grid grid-cols-3 pt-3 gap-y-10 gap-x-6">
                  <div>
                  <p className = "pt-2 text-sm text-gray-700">T001</p>
                   
                  </div>
                  <div>
            
                    <p className = "pt-2 text-sm text-gray-700">T001</p>
                  </div>
                  <div>
                    <p className = "pt-2 text-sm text-gray-700">T001</p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    )
}