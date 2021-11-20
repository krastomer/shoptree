import Navbar from "../../asset/include/navbar/Navbar";
import "./Itemprofile.css";
import Tree from "./all_product/Product_C.png"
import allProduct from "./allProduct";

export default function Itemprofile() {
  return (
    <div className="w-full h-full bg-purple-600 font-body">
      <Navbar />
      <div className="grid grid-cols-1 mt-2 gap-y-10 gap-x-6 sm:grid-cols-2 lg:grid-cols-2 xl:gap-x-8 justify-between">
          <div className="bg-theme w-full h-full flex item-center content-center object-center ">
          <div className="p-12 ">
          <img
            src= {Tree}
            // className="w-full h-full object-center object-cover lg:w-full lg:h-full"
            className="w-full h-full"
          />
        </div>
          </div>

        <div className="bg-purple-100 p-5">
            <div className = "flex justify-between font-bold text-3xl">
                <span className = "font-brown">ต้นกุหลาบหิน</span>
                <span className = "font-green">B 1,500.00</span>
            </div>
            <div className = "mt-3">T001</div> 
            <div className = "mt-3">ชื่ออังกฤษ</div>
            <div className = "mt-3">&nbsp;&nbsp; กุหลาบหินพันธุ์ไม้ เป็นพืชที่มีลักษณะเป็นพุ่ม อวบน้ำ อายุหลายปี ซึ่งมี ถิ่นกำเนิดมาจากมาดากัสกาแอฟริกาและเอเชีย  ซึ่งลักษณะเดิมก่อนที่จะ เป็นต้นที่มีลักษณะเป็นพุ่มเล็กกะทัดรัดนี้ กุหลาบหินเป็นพืชที่มีต้นสูงเก้ง ก้างหลังจากที่มีการคัดเลือกสายพันธุ์และผสม จึงทำให้เกิดกุหลาบหิน พันธุ์ต่างๆ ที่มีลักษณะเป็นพุ่มเตี้ย</div>
            <div>
            <button
            
                  className="mt-3 btn-theme group relative flex justify-center py-2 px-4 border border-3 border-green-500  font-medium rounded-md text-white hover:bg-yellow-00 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500"
                >
                  เพิ่มลงตะกร้า
                </button>
            </div>
            <div className = "mt-3">
                <span>สินค้าที่เกี่ยวข้อง</span>
                <div></div>
            </div>
        </div>
      </div>
    </div>
  );
}
