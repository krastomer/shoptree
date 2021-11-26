import React from 'react';
import { Link } from "react-router-dom";

import AllProduct from '../../services/database/NewOrder';

function StatusProcess(props){
    const Nstatus = props.Ostatus;
    if(Nstatus == "ยังไม่ตรวจสอบ"){
        return(
            <span className="inline-flex px-2 text-xs font-semibold leading-5 text-white bg-red-500 rounded-full">
                ยังไม่ตรวจสอบ
            </span>
        );
    }else if(Nstatus == "ตรวจสอบแล้ว"){
        return(
            <span className="inline-flex px-2 text-xs font-semibold leading-5 text-white bg-green-500 rounded-full">
                ตรวจสอบแล้ว
            </span>
        );
    }else 
        return(
            <span className="inline-flex px-2 text-xs font-semibold leading-5 text-red-500 bg-gray-500 rounded-full">
                มีปัญหา
            </span>
        );
}
function DiliveryProcess(props){
    const Dstatus = props.Dstatus
    const Nstatus = props.Ostatus;
    if(Nstatus == "ตรวจสอบแล้ว"){
        if(Dstatus == "รอการขนส่ง"){
            return(
                <Link
                    className={"inline-flex justify-center px-4 py-2 text-sm font-medium text-white bg-yellow-500 border border-transparent rounded-full shadow-sm hover:bg-yellow-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-yellow-500 md-2"}
                    to="#"
                >
                    เริ่มการขนส่ง
                </Link>
            );
        }else if(Dstatus == "เริ่มการขนส่ง"){
            return(
                <Link
                    className={"inline-flex justify-center px-4 py-2 text-sm font-medium text-white bg-green-500 border border-transparent rounded-full shadow-sm hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500 md-2"}
                    to="#"
                >
                    เสร็จสิ้นการขนส่ง
                </Link>
            );
        }else{
            return(
                <span
                    className={"inline-flex justify-center px-4 py-2 disabled:opacity-100 text-sm font-medium text-white bg-green-700 border border-transparent rounded-full shadow-sm md-2"}
                >
                    เสร็จสินการขนส่งแล้ว
                </span>
            );
        }   
    }else 
        return(
            <span
                className={"inline-flex justify-center px-4 py-2 text-sm disabled:opacity-50 font-medium text-white bg-gray-500 border border-transparent rounded-full shadow-sm  focus:outline-none md-2"}
            >
                รอการตรวจสอบ
            </span>
        );
}

export default function DiTable() {
  const incomes = AllProduct;
  const sum = incomes.map((income)=> income.price).reduce((a,b)=> a + b);
  return (
    <div className="flex flex-col">
      <div className="-my-2 overflow-x-auto sm:-mx-6 lg:-mx-8">
        <div className="inline-block min-w-full py-2 align-middle sm:px-6 lg:px-8">
          <div className="overflow-hidden border-b border-gray-200 shadow sm:rounded-lg">
            <table className="min-w-full divide-y divide-gray-200">
              <thead className="bg-gray-50">
                <tr>
                  <th
                    scope="col"
                    className="px-6 py-3 text-xs font-medium tracking-wider text-left text-gray-500 uppercase"
                  >
                    เลขคำสั่งซื้อ
                  </th>
                  <th
                    scope="col"
                    className="px-6 py-3 text-xs font-medium tracking-wider text-left text-gray-500 uppercase"
                  >
                    ชื่อผู้ซื้อสินค้า
                  </th>
                  <th
                    align="center"
                    scope="col"
                    className="px-6 py-3 text-xs font-medium tracking-wider text-gray-500 uppercase"
                  >
                    สถานะ
                  </th>
                  <th
                    align="center"
                    scope="col"
                    className="px-6 py-3 text-xs font-medium tracking-wider text-gray-500 uppercase"
                  >
                    ตรวจสอบคำสั่งซื้อ
                  </th>
                </tr>
              </thead>
              <tbody className="bg-white divide-y divide-gray-200">
                {incomes.map((person) => (
                  <tr key={person.email}>
                    <td className="px-6 py-4 whitespace-nowrap">
                        <div className="text-sm text-gray-500 whitespace-nowrap">
                            {person.id}
                        </div>
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap">
                        <div className="text-sm text-gray-500 whitespace-nowrap">
                            {person.Cname}
                        </div>
                    </td>
                    <td align="center" className="px-6 py-4 whitespace-nowrap">
                        <StatusProcess Ostatus={person.Ostatus}/>
                    </td>
                    <td align="center">
                       <DiliveryProcess Ostatus={person.Ostatus} Dstatus={person.Dstatus}/>
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
  );
}
