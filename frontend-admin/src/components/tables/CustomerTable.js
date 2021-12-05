import React from 'react';
import { Link } from "react-router-dom";

import AllProduct from '../../services/database/Customer';


export default function CustomerTable() {
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
                    รหัส
                  </th>
                  <th
                    scope="col"
                    className="px-6 py-3 text-xs font-medium tracking-wider text-left text-gray-500 uppercase"
                  >
                    ชื่อ
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
                        <div className="text-sm text-gray-500 whitespace-nowrap">
                            {person.id}
                        </div>
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap">
                        <div className="text-sm text-gray-500 whitespace-nowrap">
                            {person.Cname}
                        </div>
                    </td>
                    <td align="center">
                        <Link
                            className={"inline-flex justify-center px-4 py-2 text-sm font-medium text-white bg-yellow-500 border border-transparent rounded-full shadow-sm hover:bg-yellow-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-yellow-500 md-2"}
                            to="#"
                            >
                            แก้ไข
                        </Link>
                        <Link
                            className={"ml-2 inline-flex justify-center px-4 py-2 text-sm font-medium text-white bg-red-500 border border-transparent rounded-full shadow-sm hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500"}
                            to="#"
                            >
                            ลบ
                        </Link>
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
