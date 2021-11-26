import React from 'react';

import AllProduct from '../../services/database/AllProduct';

function StatusProcess(props){
    const Nstatus = props.Ostatus;
    if(Nstatus == "ยังไม่ขาย"){
        return(
            <span className="inline-flex px-2 text-xs font-semibold leading-5 text-white bg-gray-500 rounded-full">
                ยังไม่ขาย
            </span>
        );
    }else if(Nstatus == "กำลังวางขาย"){
        return(
            <span className="inline-flex px-2 text-xs font-semibold leading-5 text-white bg-green-300 rounded-full">
                กำลังวางจำหน่าย
            </span>
        );
    }else if(Nstatus == "ขายได้แล้ว"){
        return(
            <span className="inline-flex px-2 text-xs font-semibold leading-5 text-white bg-green-700 rounded-full">
                ขายได้แล้ว
            </span>
        );
    }else if(Nstatus == "มีปัญหา"){
        return(
            <span className="inline-flex px-2 text-xs font-semibold leading-5 text-white bg-red-500 rounded-full">
                มีปัญหา
            </span>
        );
    }else if(Nstatus == "จอง"){
        return(
            <span className="inline-flex px-2 text-xs font-semibold leading-5 text-white bg-yellow-500 rounded-full">
                จอง
            </span>
        );
    }else 
        return(
            <span className="inline-flex px-2 text-xs font-semibold leading-5 text-red-500 bg-gray-500 rounded-full">
                ระบบขัดข้อง
            </span>
        );
}
export default function ProductTable() {
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
                    รายการสินค้า
                  </th>
                  <th
                    scope="col"
                    className="px-6 py-3 text-xs font-medium tracking-wider text-left text-gray-500 uppercase"
                  >
                    ประเภท
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
                    ราคา
                  </th>
                  <th
                    scope="col"
                    className="px-6 py-3 text-xs font-medium tracking-wider text-gray-500 uppercase"
                  >
                    รายละเอียด
                  </th>
                  <th
                    align="center"
                    scope="col"
                    className="px-6 py-3 text-xs font-medium tracking-wider text-gray-500 uppercase"
                  >
                    จัดการ
                  </th>
                </tr>
              </thead>
              <tbody className="bg-white divide-y divide-gray-200">
                {incomes.map((person) => (
                  <tr key={person.email}>
                    <td className="px-6 py-4 whitespace-nowrap">
                    <div className="flex items-center">
                        <div className="flex-shrink-0 w-10 h-10">
                          <img className="w-10 h-10 rounded-full" src={person.image} alt="" />
                        </div>
                        <div className="ml-4">
                          <div className="text-sm font-medium text-gray-900">{person.Oname}</div>
                          <div className="text-sm text-gray-500">{person.id}</div>
                        </div>
                      </div>
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap">
                        <div className="text-sm text-gray-500 whitespace-nowrap">
                            {person.cat}
                        </div>
                    </td>
                    <td align="center" className="px-6 py-4 whitespace-nowrap">
                        <StatusProcess Ostatus={person.Ostatus}/>
                    </td>
                    <td align="center" className="px-6 py-4 text-sm text-gray-500 whitespace-nowrap">${person.price}</td>
                    <td align="center"  className="px-4 py-3">
                        <button
                            type="submit"
                            className="inline-flex justify-center px-4 py-2 text-sm font-medium text-white bg-blue-600 border border-transparent rounded-full shadow-sm hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
                        >
                        ดูรายละเอียด
                        </button>
                    </td>
                    <td align="center"  className="px-4 py-3">
                        <button
                            type="submit"
                            className="inline-flex justify-center px-4 py-2 mr-2 text-sm font-medium text-white bg-yellow-500 border border-transparent rounded-full shadow-sm hover:bg-yellow-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-yellow-500"
                        >
                        แก้ไข
                        </button>
                        <button
                            type="submit"
                            className="inline-flex justify-center px-4 py-2 text-sm font-medium text-white bg-red-500 border border-transparent rounded-full shadow-sm hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500"
                        >
                        ลบ
                        </button>
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
