import React, { useState } from "react";
function StatusLocation(props){
    const {local} = props;
    if(local == 1){
        return(
            <>
            <div className="p-5">
                <div className="p-4 mx-4">
                    <div className="flex items-center">
                        <div className="relative flex items-center text-white">
                            <div className="w-12 h-12 py-3 transition duration-500 ease-in-out bg-green-500 border-2 border-green-500 rounded-full">
                            <svg 
                                xmlns="http://www.w3.org/2000/svg" 
                                width="100%" 
                                height="100%" 
                                viewBox="0 0 24 24" 
                                fill="none" 
                                stroke="currentColor" 
                                stroke-width="2" 
                                stroke-linecap="round" 
                                stroke-linejoin="round" 
                                class="feather feather-shopping-bag">
                                <path d="M6 2L3 6v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V6l-3-4z"></path>
                                <line x1="3" y1="6" x2="21" y2="6"></line>
                                <path d="M16 10a4 4 0 0 1-8 0"></path>
                            </svg>
                            </div>
                            <div className="absolute top-0 w-32 mt-16 -ml-10 text-xs font-medium text-center text-green-500 uppercase">
                                ตรวจสอบตะกร้าสินค้า
                            </div>
                        </div>
                        <div className="flex-auto transition duration-500 ease-in-out border-t-2 border-gray-300"></div>
                        <div className="relative flex items-center text-gray-500">
                            <div className="w-12 h-12 py-3 transition duration-500 ease-in-out border-2 border-gray-300 rounded-full ">
                                <svg 
                                    xmlns="http://www.w3.org/2000/svg"
                                    width="100%" 
                                    height="100%" 
                                    viewBox="0 0 24 24" 
                                    fill="none" stroke="currentColor" 
                                    stroke-width="2" 
                                    stroke-linecap="round" 
                                    stroke-linejoin="round" 
                                    class="feather feather-truck">
                                    <rect x="1" y="3" width="15" height="13"></rect>
                                    <polygon points="16 8 20 8 23 11 23 16 16 16 16 8"></polygon>
                                    <circle cx="5.5" cy="18.5" r="2.5"></circle>
                                    <circle cx="18.5" cy="18.5" r="2.5"></circle>
                                    </svg>
                            </div>
                            <div className="absolute top-0 w-32 mt-16 -ml-10 text-xs font-medium text-center text-gray-500 uppercase">
                                เลือกที่จัดส่ง
                            </div>
                        </div>
                        <div className="flex-auto transition duration-500 ease-in-out border-t-2 border-gray-300"></div>
                        <div className="relative flex items-center text-gray-500">
                            <div className="w-12 h-12 py-3 transition duration-500 ease-in-out border-2 border-gray-300 rounded-full ">
                                <svg 
                                    xmlns="http://www.w3.org/2000/svg" 
                                    width="100%" 
                                    height="100%" 
                                    viewBox="0 0 24 24" 
                                    fill="none" 
                                    stroke="currentColor" 
                                    stroke-width="2" 
                                    stroke-linecap="round" 
                                    stroke-linejoin="round" 
                                    class="feather feather-dollar-sign">
                                    <line x1="12" y1="1" x2="12" y2="23"></line>
                                    <path d="M17 5H9.5a3.5 3.5 0 0 0 0 7h5a3.5 3.5 0 0 1 0 7H6"></path>
                                </svg>
                            </div>
                            <div className="absolute top-0 w-32 mt-16 -ml-10 text-xs font-medium text-center text-gray-500 uppercase ">
                            ตรวจสอบคำสั่งซื้อ
                            </div>
                        </div>
                        <div  className="flex-auto transition duration-500 ease-in-out border-t-2 border-gray-300"></div>
                        <div className="relative flex items-center text-gray-500">
                            <div className="w-12 h-12 py-3 transition duration-500 ease-in-out border-2 border-gray-300 rounded-full ">
                            <svg 
                                xmlns="http://www.w3.org/2000/svg" 
                                width="100%" 
                                height="100%" 
                                viewBox="0 0 24 24" 
                                fill="none" 
                                stroke="currentColor" 
                                stroke-width="2" 
                                stroke-linecap="round" 
                                stroke-linejoin="round"
                                class="feather feather-credit-card">
                                <rect x="1" y="4" width="22" height="16" rx="2" ry="2"></rect>
                                <line x1="1" y1="10" x2="23" y2="10"></line>
                            </svg>
                            </div>
                            <div className="absolute top-0 w-32 mt-16 -ml-10 text-xs font-medium text-center text-gray-500 uppercase">
                                ชำระเงิน
                            </div>
                        </div>
                        <div  className="flex-auto transition duration-500 ease-in-out border-t-2 border-gray-300"></div>
                        <div className="relative flex items-center text-gray-500">
                            <div className="w-12 h-12 py-3 transition duration-500 ease-in-out border-2 border-gray-300 rounded-full ">
                                <svg 
                                    xmlns="http://www.w3.org/2000/svg" 
                                    width="100%" 
                                    height="100%" 
                                    viewBox="0 0 24 24" 
                                    fill="none" 
                                    stroke="currentColor" 
                                    stroke-width="2" 
                                    stroke-linecap="round" 
                                    stroke-linejoin="round" 
                                    class="feather feather-check">
                                    <polyline points="20 6 9 17 4 12"></polyline>
                                </svg>
                            </div>
                            <div className="absolute top-0 w-32 mt-16 -ml-10 text-xs font-medium text-center text-gray-500 uppercase">
                                เสร็จสิ้น
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            </>
        )
    }
    else if(local == 2){
        return(
            <>
            <div className="p-5">
                <div className="p-4 mx-4">
                    <div className="flex items-center">
                        <div className="relative flex items-center text-green-500">
                            <div className="w-12 h-12 py-3 transition duration-500 ease-in-out border-2 border-green-500 rounded-full">
                            <svg 
                                xmlns="http://www.w3.org/2000/svg" 
                                width="100%" 
                                height="100%" 
                                viewBox="0 0 24 24" 
                                fill="none" 
                                stroke="currentColor" 
                                stroke-width="2" 
                                stroke-linecap="round" 
                                stroke-linejoin="round" 
                                class="feather feather-shopping-bag">
                                <path d="M6 2L3 6v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V6l-3-4z"></path>
                                <line x1="3" y1="6" x2="21" y2="6"></line>
                                <path d="M16 10a4 4 0 0 1-8 0"></path>
                            </svg>
                            </div>
                            <div className="absolute top-0 w-32 mt-16 -ml-10 text-xs font-medium text-center text-green-500 uppercase">
                                ตรวจสอบตะกร้าสินค้า
                            </div>
                        </div>
                        <div className="flex-auto transition duration-500 ease-in-out border-t-2 border-green-500"></div>
                        <div className="relative flex items-center text-white">
                            <div className="w-12 h-12 py-3 transition duration-500 ease-in-out bg-green-500 border-2 border-green-500 rounded-full">
                                <svg 
                                    xmlns="http://www.w3.org/2000/svg"
                                    width="100%" 
                                    height="100%" 
                                    viewBox="0 0 24 24" 
                                    fill="none" stroke="currentColor" 
                                    stroke-width="2" 
                                    stroke-linecap="round" 
                                    stroke-linejoin="round" 
                                    class="feather feather-truck">
                                    <rect x="1" y="3" width="15" height="13"></rect>
                                    <polygon points="16 8 20 8 23 11 23 16 16 16 16 8"></polygon>
                                    <circle cx="5.5" cy="18.5" r="2.5"></circle>
                                    <circle cx="18.5" cy="18.5" r="2.5"></circle>
                                    </svg>
                            </div>
                            <div className="absolute top-0 w-32 mt-16 -ml-10 text-xs font-medium text-center text-green-500 uppercase">
                                เลือกที่จัดส่ง
                            </div>
                        </div>
                        <div className="flex-auto transition duration-500 ease-in-out border-t-2 border-gray-300"></div>
                        <div className="relative flex items-center text-gray-500">
                            <div className="w-12 h-12 py-3 transition duration-500 ease-in-out border-2 border-gray-300 rounded-full ">
                                <svg 
                                    xmlns="http://www.w3.org/2000/svg" 
                                    width="100%" 
                                    height="100%" 
                                    viewBox="0 0 24 24" 
                                    fill="none" 
                                    stroke="currentColor" 
                                    stroke-width="2" 
                                    stroke-linecap="round" 
                                    stroke-linejoin="round" 
                                    class="feather feather-dollar-sign">
                                    <line x1="12" y1="1" x2="12" y2="23"></line>
                                    <path d="M17 5H9.5a3.5 3.5 0 0 0 0 7h5a3.5 3.5 0 0 1 0 7H6"></path>
                                </svg>
                            </div>
                            <div className="absolute top-0 w-32 mt-16 -ml-10 text-xs font-medium text-center text-gray-500 uppercase ">
                            ตรวจสอบคำสั่งซื้อ
                            </div>
                        </div>
                        <div  className="flex-auto transition duration-500 ease-in-out border-t-2 border-gray-300"></div>
                        <div className="relative flex items-center text-gray-500">
                            <div className="w-12 h-12 py-3 transition duration-500 ease-in-out border-2 border-gray-300 rounded-full ">
                            <svg 
                                xmlns="http://www.w3.org/2000/svg" 
                                width="100%" 
                                height="100%" 
                                viewBox="0 0 24 24" 
                                fill="none" 
                                stroke="currentColor" 
                                stroke-width="2" 
                                stroke-linecap="round" 
                                stroke-linejoin="round"
                                class="feather feather-credit-card">
                                <rect x="1" y="4" width="22" height="16" rx="2" ry="2"></rect>
                                <line x1="1" y1="10" x2="23" y2="10"></line>
                            </svg>
                            </div>
                            <div className="absolute top-0 w-32 mt-16 -ml-10 text-xs font-medium text-center text-gray-500 uppercase">
                                ชำระเงิน
                            </div>
                        </div>
                        <div  className="flex-auto transition duration-500 ease-in-out border-t-2 border-gray-300"></div>
                        <div className="relative flex items-center text-gray-500">
                            <div className="w-12 h-12 py-3 transition duration-500 ease-in-out border-2 border-gray-300 rounded-full ">
                                <svg 
                                    xmlns="http://www.w3.org/2000/svg" 
                                    width="100%" 
                                    height="100%" 
                                    viewBox="0 0 24 24" 
                                    fill="none" 
                                    stroke="currentColor" 
                                    stroke-width="2" 
                                    stroke-linecap="round" 
                                    stroke-linejoin="round" 
                                    class="feather feather-check">
                                    <polyline points="20 6 9 17 4 12"></polyline>
                                </svg>
                            </div>
                            <div className="absolute top-0 w-32 mt-16 -ml-10 text-xs font-medium text-center text-gray-500 uppercase">
                                เสร็จสิ้น
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            </>
        )
    }
    else if(local == 3){
        return(
            <>
            <div className="p-5">
                <div className="p-4 mx-4">
                    <div className="flex items-center">
                        <div className="relative flex items-center text-green-500">
                            <div className="w-12 h-12 py-3 transition duration-500 ease-in-out border-2 border-green-500 rounded-full">
                            <svg 
                                xmlns="http://www.w3.org/2000/svg" 
                                width="100%" 
                                height="100%" 
                                viewBox="0 0 24 24" 
                                fill="none" 
                                stroke="currentColor" 
                                stroke-width="2" 
                                stroke-linecap="round" 
                                stroke-linejoin="round" 
                                class="feather feather-shopping-bag">
                                <path d="M6 2L3 6v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V6l-3-4z"></path>
                                <line x1="3" y1="6" x2="21" y2="6"></line>
                                <path d="M16 10a4 4 0 0 1-8 0"></path>
                            </svg>
                            </div>
                            <div className="absolute top-0 w-32 mt-16 -ml-10 text-xs font-medium text-center text-green-500 uppercase">
                                ตรวจสอบตะกร้าสินค้า
                            </div>
                        </div>
                        <div className="flex-auto transition duration-500 ease-in-out border-t-2 border-green-500"></div>
                        <div className="relative flex items-center text-green-500">
                            <div className="w-12 h-12 py-3 transition duration-500 ease-in-out border-2 border-green-500 rounded-full">
                                <svg 
                                    xmlns="http://www.w3.org/2000/svg"
                                    width="100%" 
                                    height="100%" 
                                    viewBox="0 0 24 24" 
                                    fill="none" stroke="currentColor" 
                                    stroke-width="2" 
                                    stroke-linecap="round" 
                                    stroke-linejoin="round" 
                                    class="feather feather-truck">
                                    <rect x="1" y="3" width="15" height="13"></rect>
                                    <polygon points="16 8 20 8 23 11 23 16 16 16 16 8"></polygon>
                                    <circle cx="5.5" cy="18.5" r="2.5"></circle>
                                    <circle cx="18.5" cy="18.5" r="2.5"></circle>
                                    </svg>
                            </div>
                            <div className="absolute top-0 w-32 mt-16 -ml-10 text-xs font-medium text-center text-green-500 uppercase">
                                เลือกที่จัดส่ง
                            </div>
                        </div>
                        <div className="flex-auto transition duration-500 ease-in-out border-t-2 border-green-500"></div>
                        <div className="relative flex items-center text-white">
                            <div className="w-12 h-12 py-3 transition duration-500 ease-in-out bg-green-500 border-2 border-green-500 rounded-full">
                                <svg 
                                    xmlns="http://www.w3.org/2000/svg" 
                                    width="100%" 
                                    height="100%" 
                                    viewBox="0 0 24 24" 
                                    fill="none" 
                                    stroke="currentColor" 
                                    stroke-width="2" 
                                    stroke-linecap="round" 
                                    stroke-linejoin="round" 
                                    class="feather feather-dollar-sign">
                                    <line x1="12" y1="1" x2="12" y2="23"></line>
                                    <path d="M17 5H9.5a3.5 3.5 0 0 0 0 7h5a3.5 3.5 0 0 1 0 7H6"></path>
                                </svg>
                            </div>
                            <div className="absolute top-0 w-32 mt-16 -ml-10 text-xs font-medium text-center text-green-500 uppercase ">
                            ตรวจสอบคำสั่งซื้อ
                            </div>
                        </div>
                        <div  className="flex-auto transition duration-500 ease-in-out border-t-2 border-gray-300"></div>
                        <div className="relative flex items-center text-gray-500">
                            <div className="w-12 h-12 py-3 transition duration-500 ease-in-out border-2 border-gray-300 rounded-full ">
                            <svg 
                                xmlns="http://www.w3.org/2000/svg" 
                                width="100%" 
                                height="100%" 
                                viewBox="0 0 24 24" 
                                fill="none" 
                                stroke="currentColor" 
                                stroke-width="2" 
                                stroke-linecap="round" 
                                stroke-linejoin="round"
                                class="feather feather-credit-card">
                                <rect x="1" y="4" width="22" height="16" rx="2" ry="2"></rect>
                                <line x1="1" y1="10" x2="23" y2="10"></line>
                            </svg>
                            </div>
                            <div className="absolute top-0 w-32 mt-16 -ml-10 text-xs font-medium text-center text-gray-500 uppercase">
                                ชำระเงิน
                            </div>
                        </div>
                        <div  className="flex-auto transition duration-500 ease-in-out border-t-2 border-gray-300"></div>
                        <div className="relative flex items-center text-gray-500">
                            <div className="w-12 h-12 py-3 transition duration-500 ease-in-out border-2 border-gray-300 rounded-full ">
                                <svg 
                                    xmlns="http://www.w3.org/2000/svg" 
                                    width="100%" 
                                    height="100%" 
                                    viewBox="0 0 24 24" 
                                    fill="none" 
                                    stroke="currentColor" 
                                    stroke-width="2" 
                                    stroke-linecap="round" 
                                    stroke-linejoin="round" 
                                    class="feather feather-check">
                                    <polyline points="20 6 9 17 4 12"></polyline>
                                </svg>
                            </div>
                            <div className="absolute top-0 w-32 mt-16 -ml-10 text-xs font-medium text-center text-gray-500 uppercase">
                                เสร็จสิ้น
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            </>
        )
    }
    else if(local == 4){
        return(
            <>
            <div className="p-5">
                <div className="p-4 mx-4">
                    <div className="flex items-center">
                        <div className="relative flex items-center text-green-500">
                            <div className="w-12 h-12 py-3 transition duration-500 ease-in-out border-2 border-green-500 rounded-full">
                            <svg 
                                xmlns="http://www.w3.org/2000/svg" 
                                width="100%" 
                                height="100%" 
                                viewBox="0 0 24 24" 
                                fill="none" 
                                stroke="currentColor" 
                                stroke-width="2" 
                                stroke-linecap="round" 
                                stroke-linejoin="round" 
                                class="feather feather-shopping-bag">
                                <path d="M6 2L3 6v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V6l-3-4z"></path>
                                <line x1="3" y1="6" x2="21" y2="6"></line>
                                <path d="M16 10a4 4 0 0 1-8 0"></path>
                            </svg>
                            </div>
                            <div className="absolute top-0 w-32 mt-16 -ml-10 text-xs font-medium text-center text-green-500 uppercase">
                                ตรวจสอบตะกร้าสินค้า
                            </div>
                        </div>
                        <div className="flex-auto transition duration-500 ease-in-out border-t-2 border-green-500"></div>
                        <div className="relative flex items-center text-green-500">
                            <div className="w-12 h-12 py-3 transition duration-500 ease-in-out border-2 border-green-500 rounded-full">
                                <svg 
                                    xmlns="http://www.w3.org/2000/svg"
                                    width="100%" 
                                    height="100%" 
                                    viewBox="0 0 24 24" 
                                    fill="none" stroke="currentColor" 
                                    stroke-width="2" 
                                    stroke-linecap="round" 
                                    stroke-linejoin="round" 
                                    class="feather feather-truck">
                                    <rect x="1" y="3" width="15" height="13"></rect>
                                    <polygon points="16 8 20 8 23 11 23 16 16 16 16 8"></polygon>
                                    <circle cx="5.5" cy="18.5" r="2.5"></circle>
                                    <circle cx="18.5" cy="18.5" r="2.5"></circle>
                                    </svg>
                            </div>
                            <div className="absolute top-0 w-32 mt-16 -ml-10 text-xs font-medium text-center text-green-500 uppercase">
                                เลือกที่จัดส่ง
                            </div>
                        </div>
                        <div className="flex-auto transition duration-500 ease-in-out border-t-2 border-green-500"></div>
                        <div className="relative flex items-center text-green-500">
                            <div className="w-12 h-12 py-3 transition duration-500 ease-in-out border-2 border-green-500 rounded-full">
                                <svg 
                                    xmlns="http://www.w3.org/2000/svg" 
                                    width="100%" 
                                    height="100%" 
                                    viewBox="0 0 24 24" 
                                    fill="none" 
                                    stroke="currentColor" 
                                    stroke-width="2" 
                                    stroke-linecap="round" 
                                    stroke-linejoin="round" 
                                    class="feather feather-dollar-sign">
                                    <line x1="12" y1="1" x2="12" y2="23"></line>
                                    <path d="M17 5H9.5a3.5 3.5 0 0 0 0 7h5a3.5 3.5 0 0 1 0 7H6"></path>
                                </svg>
                            </div>
                            <div className="absolute top-0 w-32 mt-16 -ml-10 text-xs font-medium text-center text-gray-500 uppercase ">
                            ตรวจสอบคำสั่งซื้อ
                            </div>
                        </div>
                        <div className="flex-auto transition duration-500 ease-in-out border-t-2 border-green-500"></div>
                        <div className="relative flex items-center text-white">
                            <div className="w-12 h-12 py-3 transition duration-500 ease-in-out bg-green-500 border-2 border-green-500 rounded-full">
                            <svg 
                                xmlns="http://www.w3.org/2000/svg" 
                                width="100%" 
                                height="100%" 
                                viewBox="0 0 24 24" 
                                fill="none" 
                                stroke="currentColor" 
                                stroke-width="2" 
                                stroke-linecap="round" 
                                stroke-linejoin="round"
                                class="feather feather-credit-card">
                                <rect x="1" y="4" width="22" height="16" rx="2" ry="2"></rect>
                                <line x1="1" y1="10" x2="23" y2="10"></line>
                            </svg>
                            </div>
                            <div className="absolute top-0 w-32 mt-16 -ml-10 text-xs font-medium text-center text-green-500 uppercase">
                                ชำระเงิน
                            </div>
                        </div>
                        <div  className="flex-auto transition duration-500 ease-in-out border-t-2 border-gray-300"></div>
                        <div className="relative flex items-center text-gray-500">
                            <div className="w-12 h-12 py-3 transition duration-500 ease-in-out border-2 border-gray-300 rounded-full ">
                                <svg 
                                    xmlns="http://www.w3.org/2000/svg" 
                                    width="100%" 
                                    height="100%" 
                                    viewBox="0 0 24 24" 
                                    fill="none" 
                                    stroke="currentColor" 
                                    stroke-width="2" 
                                    stroke-linecap="round" 
                                    stroke-linejoin="round" 
                                    class="feather feather-check">
                                    <polyline points="20 6 9 17 4 12"></polyline>
                                </svg>
                            </div>
                            <div className="absolute top-0 w-32 mt-16 -ml-10 text-xs font-medium text-center text-gray-500 uppercase">
                                เสร็จสิ้น
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            </>
        )
    }
    else if(local == 5){
        return(
            <>
            <div className="p-5">
                <div className="p-4 mx-4">
                    <div className="flex items-center">
                        <div className="relative flex items-center text-green-500">
                            <div className="w-12 h-12 py-3 transition duration-500 ease-in-out border-2 border-green-500 rounded-full">
                            <svg 
                                xmlns="http://www.w3.org/2000/svg" 
                                width="100%" 
                                height="100%" 
                                viewBox="0 0 24 24" 
                                fill="none" 
                                stroke="currentColor" 
                                stroke-width="2" 
                                stroke-linecap="round" 
                                stroke-linejoin="round" 
                                class="feather feather-shopping-bag">
                                <path d="M6 2L3 6v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V6l-3-4z"></path>
                                <line x1="3" y1="6" x2="21" y2="6"></line>
                                <path d="M16 10a4 4 0 0 1-8 0"></path>
                            </svg>
                            </div>
                            <div className="absolute top-0 w-32 mt-16 -ml-10 text-xs font-medium text-center text-green-500 uppercase">
                                ตรวจสอบตะกร้าสินค้า
                            </div>
                        </div>
                        <div className="flex-auto transition duration-500 ease-in-out border-t-2 border-green-500"></div>
                        <div className="relative flex items-center text-green-500">
                            <div className="w-12 h-12 py-3 transition duration-500 ease-in-out border-2 border-green-500 rounded-full">
                                <svg 
                                    xmlns="http://www.w3.org/2000/svg"
                                    width="100%" 
                                    height="100%" 
                                    viewBox="0 0 24 24" 
                                    fill="none" stroke="currentColor" 
                                    stroke-width="2" 
                                    stroke-linecap="round" 
                                    stroke-linejoin="round" 
                                    class="feather feather-truck">
                                    <rect x="1" y="3" width="15" height="13"></rect>
                                    <polygon points="16 8 20 8 23 11 23 16 16 16 16 8"></polygon>
                                    <circle cx="5.5" cy="18.5" r="2.5"></circle>
                                    <circle cx="18.5" cy="18.5" r="2.5"></circle>
                                    </svg>
                            </div>
                            <div className="absolute top-0 w-32 mt-16 -ml-10 text-xs font-medium text-center text-green-500 uppercase">
                                เลือกที่จัดส่ง
                            </div>
                        </div>
                        <div className="flex-auto transition duration-500 ease-in-out border-t-2 border-green-500"></div>
                        <div className="relative flex items-center text-green-500">
                            <div className="w-12 h-12 py-3 transition duration-500 ease-in-out border-2 border-green-500 rounded-full">
                                <svg 
                                    xmlns="http://www.w3.org/2000/svg" 
                                    width="100%" 
                                    height="100%" 
                                    viewBox="0 0 24 24" 
                                    fill="none" 
                                    stroke="currentColor" 
                                    stroke-width="2" 
                                    stroke-linecap="round" 
                                    stroke-linejoin="round" 
                                    class="feather feather-dollar-sign">
                                    <line x1="12" y1="1" x2="12" y2="23"></line>
                                    <path d="M17 5H9.5a3.5 3.5 0 0 0 0 7h5a3.5 3.5 0 0 1 0 7H6"></path>
                                </svg>
                            </div>
                            <div className="absolute top-0 w-32 mt-16 -ml-10 text-xs font-medium text-center text-gray-500 uppercase ">
                            ตรวจสอบคำสั่งซื้อ
                            </div>
                        </div>
                        <div className="flex-auto transition duration-500 ease-in-out border-t-2 border-green-500"></div>
                        <div className="relative flex items-center text-green-500">
                            <div className="w-12 h-12 py-3 transition duration-500 ease-in-out border-2 border-green-500 rounded-full">
                            <svg 
                                xmlns="http://www.w3.org/2000/svg" 
                                width="100%" 
                                height="100%" 
                                viewBox="0 0 24 24" 
                                fill="none" 
                                stroke="currentColor" 
                                stroke-width="2" 
                                stroke-linecap="round" 
                                stroke-linejoin="round"
                                class="feather feather-credit-card">
                                <rect x="1" y="4" width="22" height="16" rx="2" ry="2"></rect>
                                <line x1="1" y1="10" x2="23" y2="10"></line>
                            </svg>
                            </div>
                            <div className="absolute top-0 w-32 mt-16 -ml-10 text-xs font-medium text-center text-green-500 uppercase">
                                ชำระเงิน
                            </div>
                        </div>
                        <div className="flex-auto transition duration-500 ease-in-out border-t-2 border-green-500"></div>
                        <div className="relative flex items-center text-white">
                            <div className="w-12 h-12 py-3 transition duration-500 ease-in-out bg-green-500 border-2 border-green-500 rounded-full">
                                <svg 
                                    xmlns="http://www.w3.org/2000/svg" 
                                    width="100%" 
                                    height="100%" 
                                    viewBox="0 0 24 24" 
                                    fill="none" 
                                    stroke="currentColor" 
                                    stroke-width="2" 
                                    stroke-linecap="round" 
                                    stroke-linejoin="round" 
                                    class="feather feather-check">
                                    <polyline points="20 6 9 17 4 12"></polyline>
                                </svg>
                            </div>
                            <div className="absolute top-0 w-32 mt-16 -ml-10 text-xs font-medium text-center text-green-500 uppercase">
                                เสร็จสิ้น
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            </>
        )
    }
    else{
        return(
            <>
            <div className="p-5">
                <div className="p-4 mx-4">
                    <div className="flex items-center">
                        <div className="relative flex items-center text-green-500">
                            <div className="w-12 h-12 py-3 transition duration-500 ease-in-out border-2 border-green-500 rounded-full">
                            <svg 
                                xmlns="http://www.w3.org/2000/svg" 
                                width="100%" 
                                height="100%" 
                                viewBox="0 0 24 24" 
                                fill="none" 
                                stroke="currentColor" 
                                stroke-width="2" 
                                stroke-linecap="round" 
                                stroke-linejoin="round" 
                                class="feather feather-shopping-bag">
                                <path d="M6 2L3 6v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V6l-3-4z"></path>
                                <line x1="3" y1="6" x2="21" y2="6"></line>
                                <path d="M16 10a4 4 0 0 1-8 0"></path>
                            </svg>
                            </div>
                            <div className="absolute top-0 w-32 mt-16 -ml-10 text-xs font-medium text-center text-green-500 uppercase">
                                ตรวจสอบตะกร้าสินค้า
                            </div>
                        </div>
                        <div className="flex-auto transition duration-500 ease-in-out border-t-2 border-green-500"></div>
                        <div className="relative flex items-center text-white">
                            <div className="w-12 h-12 py-3 transition duration-500 ease-in-out bg-green-500 border-2 border-green-500 rounded-full">
                                <svg 
                                    xmlns="http://www.w3.org/2000/svg"
                                    width="100%" 
                                    height="100%" 
                                    viewBox="0 0 24 24" 
                                    fill="none" stroke="currentColor" 
                                    stroke-width="2" 
                                    stroke-linecap="round" 
                                    stroke-linejoin="round" 
                                    class="feather feather-truck">
                                    <rect x="1" y="3" width="15" height="13"></rect>
                                    <polygon points="16 8 20 8 23 11 23 16 16 16 16 8"></polygon>
                                    <circle cx="5.5" cy="18.5" r="2.5"></circle>
                                    <circle cx="18.5" cy="18.5" r="2.5"></circle>
                                    </svg>
                            </div>
                            <div className="absolute top-0 w-32 mt-16 -ml-10 text-xs font-medium text-center text-green-500 uppercase">
                                เลือกที่จัดส่ง
                            </div>
                        </div>
                        <div className="flex-auto transition duration-500 ease-in-out border-t-2 border-gray-300"></div>
                        <div className="relative flex items-center text-gray-500">
                            <div className="w-12 h-12 py-3 transition duration-500 ease-in-out border-2 border-gray-300 rounded-full ">
                                <svg 
                                    xmlns="http://www.w3.org/2000/svg" 
                                    width="100%" 
                                    height="100%" 
                                    viewBox="0 0 24 24" 
                                    fill="none" 
                                    stroke="currentColor" 
                                    stroke-width="2" 
                                    stroke-linecap="round" 
                                    stroke-linejoin="round" 
                                    class="feather feather-dollar-sign">
                                    <line x1="12" y1="1" x2="12" y2="23"></line>
                                    <path d="M17 5H9.5a3.5 3.5 0 0 0 0 7h5a3.5 3.5 0 0 1 0 7H6"></path>
                                </svg>
                            </div>
                            <div className="absolute top-0 w-32 mt-16 -ml-10 text-xs font-medium text-center text-gray-500 uppercase ">
                            ตรวจสอบคำสั่งซื้อ
                            </div>
                        </div>
                        <div  className="flex-auto transition duration-500 ease-in-out border-t-2 border-gray-300"></div>
                        <div className="relative flex items-center text-gray-500">
                            <div className="w-12 h-12 py-3 transition duration-500 ease-in-out border-2 border-gray-300 rounded-full ">
                            <svg 
                                xmlns="http://www.w3.org/2000/svg" 
                                width="100%" 
                                height="100%" 
                                viewBox="0 0 24 24" 
                                fill="none" 
                                stroke="currentColor" 
                                stroke-width="2" 
                                stroke-linecap="round" 
                                stroke-linejoin="round"
                                class="feather feather-credit-card">
                                <rect x="1" y="4" width="22" height="16" rx="2" ry="2"></rect>
                                <line x1="1" y1="10" x2="23" y2="10"></line>
                            </svg>
                            </div>
                            <div className="absolute top-0 w-32 mt-16 -ml-10 text-xs font-medium text-center text-gray-500 uppercase">
                                ชำระเงิน
                            </div>
                        </div>
                        <div  className="flex-auto transition duration-500 ease-in-out border-t-2 border-gray-300"></div>
                        <div className="relative flex items-center text-gray-500">
                            <div className="w-12 h-12 py-3 transition duration-500 ease-in-out border-2 border-gray-300 rounded-full ">
                                <svg 
                                    xmlns="http://www.w3.org/2000/svg" 
                                    width="100%" 
                                    height="100%" 
                                    viewBox="0 0 24 24" 
                                    fill="none" 
                                    stroke="currentColor" 
                                    stroke-width="2" 
                                    stroke-linecap="round" 
                                    stroke-linejoin="round" 
                                    class="feather feather-check">
                                    <polyline points="20 6 9 17 4 12"></polyline>
                                </svg>
                            </div>
                            <div className="absolute top-0 w-32 mt-16 -ml-10 text-xs font-medium text-center text-gray-500 uppercase">
                                เสร็จสิ้น
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            </>
        )
    }
}
export default function Statusbar(props) {
  return (
    <StatusLocation local = {props.stateLocal}/>
  );
}

