import React from 'react';
import { HiAnnotation } from "react-icons/hi";
import truck from '../../assets/truck.svg'
export default function ItemOrderDelivery() {
  return (
    <div className="font-body">
      <div className="grid grid-cols-4 gap-4 rounded-3xl bg-gradient-to-r from-blue-100 to-blue-400">
        <div className="col-span-2 pt-2 mt-2">&nbsp;</div>
        <div className="col-span-2 pt-2 mt-2 text-xl font-bold">สินค้ารอจัดส่ง</div>
        <div className="col-span-2">x
          <img src={truck} width="100" height="83.33"/></div>
        <div className="flex flex-row pb-2 mb-2 text-4xl font-bold">0<p className="px-2 text-lg">รายการ</p></div>
      </div>
    </div>
  );
}
