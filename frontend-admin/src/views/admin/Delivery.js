import React from 'react';

import DeliveryTable from '../../components/tables/DeliveryTable';

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
                  การขนส่ง
              </div>
              <div className="py-4">
                <DeliveryTable />
              </div>
            </div>
          </div>
    </>
  );
}

