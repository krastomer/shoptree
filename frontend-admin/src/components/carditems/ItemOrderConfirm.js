import React from 'react';
import { HiAnnotation } from "react-icons/hi";
import onlinepayment from '../../assets/online-payment.svg'
export default function ItemOrderConfirm() {
  return (
    <div className="font-body">
      <div className="grid grid-cols-4 gap-4 rounded-3xl bg-gradient-to-r from-green-200 to-green-500">
        <div className="col-span-2 pt-2 ">&nbsp;</div>
        <div className="col-span-2 pt-2 mt-2 text-xl font-bold">ยืนยันแล้ว</div>
        <div className="col-span-2 ">
          <img src={onlinepayment}width="100" height="83.33" />
        </div>
        <div className="flex flex-row pb-2 mb-2 text-4xl font-bold">0<p className="px-2 text-lg">รายการ</p></div>
      </div>
    </div>
  );
}
