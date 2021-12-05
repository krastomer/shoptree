import React from 'react';

import HomeIncome from '../../services/database/HomeIncome';

export default function HomeTable() {
  const incomes = HomeIncome;
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
                    ชื่อสินค้า
                  </th>
                  <th
                    scope="col"
                    className="px-6 py-3 text-xs font-medium tracking-wider text-left text-gray-500 uppercase"
                  >
                    ชื่อผู้ซื้อสินค้า
                  </th>
                  <th
                    scope="col"
                    className="px-6 py-3 text-xs font-medium tracking-wider text-left text-gray-500 uppercase"
                  >
                    สถานะ
                  </th>
                  <th
                    align="center"
                    scope="col"
                    className="px-6 py-3 text-xs font-medium tracking-wider text-gray-500 uppercase"
                  >
                    ราคาที่ขายได้
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
                      <div className="text-sm text-gray-900">{person.Cname}</div>
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap">
                      <span className="inline-flex px-2 text-xs font-semibold leading-5 text-green-800 bg-green-100 rounded-full">
                        ขายได้แล้ว
                      </span>
                    </td>
                    <td align="center" className="px-6 py-4 text-sm text-gray-500 whitespace-nowrap">${person.price}</td>
                  </tr>
                ))}
              </tbody>
              <tfoot className="py-8">
                <tr className="px-6 py-6 mt-6 font-medium tracking-wider text-left text-gray-500 uppercase md-6">
                    <td align="center">รายได้ทั้งหมด</td>
                    <td></td>
                    <td></td>
                    <td align="center" className="font-bold text-green-500">${sum}</td>
                </tr>
              </tfoot>
            </table>
          </div>
        </div>
      </div>
    </div>
  );
}
