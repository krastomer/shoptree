import React from 'react';
import { Link } from "react-router-dom";

import ConfirmTable from '../../components/tables/ConfirmTable';

import AdminNavbar from '../../components/navbar/AdminNavbar';
import AdminSidebar from '../../components/sidebar/AdminSidebar';


export default function Category() {
  return (
    <>
    <AdminNavbar />
    <AdminSidebar />
          <div className="py-20 pl-dash">
            <div className="">
              <div className="flex justify-between text-3xl font-bold">
                  ตรวจสอบคำสั่งซื้อ
                  <Link
                    className={"ml-2 mr-8 inline-flex justify-center px-4 py-2 text-sm font-medium text-white bg-red-500 border border-transparent rounded-full shadow-sm hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500"}
                    to="#"
                >
                 ตรวจสอบทันใจ
                </Link>
              </div>
              <div className="py-4">
                <ConfirmTable />
              </div>
            </div>
          </div>
    </>
  );
}

