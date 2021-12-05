import Navbar from "../../asset/include/navbar/Navbar";
import "./Itemprofile.css";
import Tree from "./all_product/Product_C.png";
import allProduct from "./allProduct";

export default function Itemprofile() {
  return (
    <div className="w-full h-full font-body">
      <Navbar />
      <div className="bg-white font-body">
        <div className="max-w-2xl px-4 py-16 mx-auto rounded-lg sm:py-12 sm:px-6 lg:max-w-7xl lg:px-8">
          <div className="grid justify-between grid-cols-1 mt-2 gap-y-10 gap-x-6 sm:grid-cols-1 lg:grid-cols-2 xl:gap-x-8">
            <div className="flex content-center object-center w-full h-full item-center ">
              <div className="p-6 ">
                <img
                  src={
                    "https://www.jorakay.co.th/ckfinder/userfiles/images/%E0%B8%9E%E0%B8%B2%E0%B8%A2%E0%B8%B8%E0%B8%9D%E0%B8%99%E0%B9%80%E0%B8%AA%E0%B8%B5%E0%B9%88%E0%B8%A2%E0%B8%87%E0%B8%97%E0%B8%B3%E0%B8%95%E0%B9%89%E0%B8%99%E0%B9%84%E0%B8%A1%E0%B9%89%E0%B8%A5%E0%B9%89%E0%B8%A1%E0%B8%97%E0%B8%B1%E0%B8%9A%E0%B8%9A%E0%B9%89%E0%B8%B2%E0%B8%99%20%E0%B8%A1%E0%B8%B5%E0%B8%A7%E0%B8%B4%E0%B8%98%E0%B8%B5%E0%B8%9B%E0%B9%89%E0%B8%AD%E0%B8%87%E0%B8%81%E0%B8%B1%E0%B8%99%E0%B8%AD%E0%B8%A2%E0%B9%88%E0%B8%B2%E0%B8%87%E0%B9%84%E0%B8%A3%20%20%E0%B9%83%E0%B8%99%201.jpg"
                  }
                  // className="object-cover object-center w-full h-full lg:w-full lg:h-full"
                  className="object-contain"
                />
              </div>
            </div>

            <div className="p-6">
              <div className="grid grid-cols-1">
                <div className="flex justify-between text-3xl font-bold">
                  <p className="font-brown">ต้นกุหลาบหิน</p>
                  <p className="font-green">B 2500</p>
                </div>
                <p className="text-gray-700">T001</p>
                <div className="flex justify-between mt-6 text-black ">
                  <p className="text-gray-700">ชื่ออังกฤษ :</p>
                  <p className="text-xl font-bold uppercase">heessub</p>
                </div>
                <p className="mt-10">
                  &nbsp;&nbsp; กุหลาบหินพันธุ์ไม้ เป็นพืชที่มีลักษณะเป็นพุ่ม
                  อวบน้ำ อายุหลายปี ซึ่งมี
                  ถิ่นกำเนิดมาจากมาดากัสกาแอฟริกาและเอเชีย
                  ซึ่งลักษณะเดิมก่อนที่จะ
                  เป็นต้นที่มีลักษณะเป็นพุ่มเล็กกะทัดรัดนี้
                  กุหลาบหินเป็นพืชที่มีต้นสูงเก้ง
                  ก้างหลังจากที่มีการคัดเลือกสายพันธุ์และผสม
                  จึงทำให้เกิดกุหลาบหิน พันธุ์ต่างๆ ที่มีลักษณะเป็นพุ่มเตี้ย
                </p>
              </div>
              <div className="pt-10">
                <button className="relative flex justify-center px-4 py-2 font-medium text-white text-green-500 border border-2 border-green-500 rounded-lg cursor-pointer group hover:bg-yellow-00 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500 hover:border-green-700 hover:text-green-700">
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
  );
}
