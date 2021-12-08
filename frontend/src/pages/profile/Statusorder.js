import react from "react";
import allLocation from "./allLocation";
import Applogo from "../../asset/include/navbar/LogoBanner.png";
import Statusbar from "../order/Statusbar";

export default function statusorder() {
  return (
    <div className="w-full h-full font-body ">
      <div className="bg-gray-50 font-body rounded-lg">
        <div className="max-w-2xl px-4 py-8 mx-auto rounded-lg sm:py-8 sm:px-6 lg:max-w-7xl lg:px-8">
          {/* <div className="grid justify-between grid-cols-1 mt-2 gap-y-10 gap-x-6 sm:grid-cols-1 lg:grid-cols-2 xl:gap-x-8"> */}
          {/* <div className="grid justify-between grid-row-3 mt-2 gap-y-10 gap-x-6  xl:gap-x-8"> */}
          <div className="grid justify-center items-center grid-cols-1 mt-2 gap-y-10 gap-x-6 xl:gap-x-8 sm:grid-cols-1 lg:grid-cols-4">
            {/* <div className="flex flex-row p-4 pt-2 pb-2 mt-2 mb-2 bg-white border border-gray-200 rounded-lg shadow-md max-w"> */}
            <div className="">
              <img classname="object-contain" src={Applogo} alt="Applogo" />
            </div>
            <div class="col-span-2">
              <Statusbar />
            </div>
            <div>
              {allLocation.map((location) => (
                <>
                  <div className="flex flex-row p-4 pt-2 pb-2 mt-2 mb-2 bg-white border border-gray-200 rounded-lg shadow-md max-w">
                    <div href="#" className="px-2">
                      <p className="mb-2 text-lg font-bold tracking-tight text-gray-900">
                        {location.name}
                      </p>
                      <p class="font-normal text-gray-700">
                        {location.disFirst}
                      </p>
                      <p class="font-normal text-gray-700">
                        {location.disSecond}
                      </p>
                      <p class="font-normal text-gray-700">
                        {location.postNumber}
                      </p>
                      <p class="font-normal text-gray-700">
                        {location.phoneNumber}
                      </p>
                    </div>
                  </div>
                </>
              ))}
            </div>
          </div>
          <div className="grid justify-between grid-cols-4 mt-2 gap-y-10 gap-x-6 xl:gap-x-8 ">
            <div className="relative flex justify-center ">
              <img
                src={
                  "https://www.jorakay.co.th/ckfinder/userfiles/images/%E0%B8%9E%E0%B8%B2%E0%B8%A2%E0%B8%B8%E0%B8%9D%E0%B8%99%E0%B9%80%E0%B8%AA%E0%B8%B5%E0%B9%88%E0%B8%A2%E0%B8%87%E0%B8%97%E0%B8%B3%E0%B8%95%E0%B9%89%E0%B8%99%E0%B9%84%E0%B8%A1%E0%B9%89%E0%B8%A5%E0%B9%89%E0%B8%A1%E0%B8%97%E0%B8%B1%E0%B8%9A%E0%B8%9A%E0%B9%89%E0%B8%B2%E0%B8%99%20%E0%B8%A1%E0%B8%B5%E0%B8%A7%E0%B8%B4%E0%B8%98%E0%B8%B5%E0%B8%9B%E0%B9%89%E0%B8%AD%E0%B8%87%E0%B8%81%E0%B8%B1%E0%B8%99%E0%B8%AD%E0%B8%A2%E0%B9%88%E0%B8%B2%E0%B8%87%E0%B9%84%E0%B8%A3%20%20%E0%B9%83%E0%B8%99%201.jpg"
                }
                className="object-contain w-10 h-10 rounded-full"
              />
            </div>
            <div className="relative flex justify-center items-center">
              <p className="text-gray-500">T001</p>
            </div>
            <div className="relative flex justify-center items-center">
              <p className="font-green">ต้นกุหลาบหิน</p>
            </div>
            <div className="relative flex justify-center items-center">
              <p className="text-black">2500 บาท</p>
            </div>
            <div className="relative flex justify-center ">
              <img
                src={
                  "https://www.jorakay.co.th/ckfinder/userfiles/images/%E0%B8%9E%E0%B8%B2%E0%B8%A2%E0%B8%B8%E0%B8%9D%E0%B8%99%E0%B9%80%E0%B8%AA%E0%B8%B5%E0%B9%88%E0%B8%A2%E0%B8%87%E0%B8%97%E0%B8%B3%E0%B8%95%E0%B9%89%E0%B8%99%E0%B9%84%E0%B8%A1%E0%B9%89%E0%B8%A5%E0%B9%89%E0%B8%A1%E0%B8%97%E0%B8%B1%E0%B8%9A%E0%B8%9A%E0%B9%89%E0%B8%B2%E0%B8%99%20%E0%B8%A1%E0%B8%B5%E0%B8%A7%E0%B8%B4%E0%B8%98%E0%B8%B5%E0%B8%9B%E0%B9%89%E0%B8%AD%E0%B8%87%E0%B8%81%E0%B8%B1%E0%B8%99%E0%B8%AD%E0%B8%A2%E0%B9%88%E0%B8%B2%E0%B8%87%E0%B9%84%E0%B8%A3%20%20%E0%B9%83%E0%B8%99%201.jpg"
                }
                className="object-contain w-10 h-10 rounded-full"
              />
            </div>
            <div className="relative flex justify-center items-center">
              <p className="text-gray-500">T001</p>
            </div>
            <div className="relative flex justify-center items-center">
              <p className="font-green">ต้นกุหลาบหิน</p>
            </div>
            <div className="relative flex justify-center items-center">
              <p className="text-black">2500 บาท</p>
            </div>
          </div>

          <div className="flex justify-between">
            <div>
              <p className="font-green text-xl font-bold mt-4">ราคา 2500 บาท</p>
            </div>
            {/* <div className="flex flex-col px-4">
              <img
                src={
                  "https://www.jorakay.co.th/ckfinder/userfiles/images/%E0%B8%9E%E0%B8%B2%E0%B8%A2%E0%B8%B8%E0%B8%9D%E0%B8%99%E0%B9%80%E0%B8%AA%E0%B8%B5%E0%B9%88%E0%B8%A2%E0%B8%87%E0%B8%97%E0%B8%B3%E0%B8%95%E0%B9%89%E0%B8%99%E0%B9%84%E0%B8%A1%E0%B9%89%E0%B8%A5%E0%B9%89%E0%B8%A1%E0%B8%97%E0%B8%B1%E0%B8%9A%E0%B8%9A%E0%B9%89%E0%B8%B2%E0%B8%99%20%E0%B8%A1%E0%B8%B5%E0%B8%A7%E0%B8%B4%E0%B8%98%E0%B8%B5%E0%B8%9B%E0%B9%89%E0%B8%AD%E0%B8%87%E0%B8%81%E0%B8%B1%E0%B8%99%E0%B8%AD%E0%B8%A2%E0%B9%88%E0%B8%B2%E0%B8%87%E0%B9%84%E0%B8%A3%20%20%E0%B9%83%E0%B8%99%201.jpg"
                }
                className="object-contain w-20 h-20 rounded-full"
              />
            </div> */}
          </div>
        </div>
      </div>
    </div>
  );
}
