import React from 'react';
import shipped from '../../assets/shipped.svg';
export default function ItemOrderDelivered() {
  return (
    <div className="font-body">
      <div className="grid grid-cols-4 gap-4 rounded-3xl bg-gradient-to-r from-blue-400 to-blue-700">
        <div className="col-span-2 pt-2">&nbsp;</div>
        <div className="col-span-2 pt-2 mt-2 text-xl font-bold">สินค้ารอจัดส่ง</div>
        <div className="col-span-2">
          <img className="mb-4 ml-4" src={shipped}/>
        </div>
        <div className="flex flex-row pb-2 mb-2 text-4xl font-bold">0<p className="px-2 text-lg">รายการ</p></div>
      </div>
    </div>
  );
}
