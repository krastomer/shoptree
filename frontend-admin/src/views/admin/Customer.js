import React from 'react';
import { Link } from "react-router-dom";

import CateTable from '../../components/tables/CustomerTable';

import AdminNavbar from '../../components/navbar/AdminNavbar';
import AdminSidebar from '../../components/sidebar/AdminSidebar';


export default function Customer() {
  return (
    <>
    <AdminNavbar />
    <AdminSidebar />
          <div className="py-20 pl-dash">
            <div className="">
              <div className="flex justify-between text-3xl font-bold">
                  แสดงรายชื่อลูกค้า
                  <Link
                    className={"ml-2 mr-8 inline-flex justify-center px-4 py-2 text-sm font-medium text-white bg-green-500 border border-transparent rounded-full shadow-sm hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500"}
                    to="#"
                  >
                   เพิ่ม
                  </Link>
              </div>
              <div className="py-4">
                <CateTable />
              </div>
            </div>
          </div>
    </>
  );
}

