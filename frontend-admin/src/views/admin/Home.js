import React from 'react';

import ItemOrderWait from '../../components/carditems/ItemOrderWait';
import ItemOrderConfirm from '../../components/carditems/ItemOrderConfirm';
import ItemOrderDelivery from '../../components/carditems/ItemOrderDelivery';
import ItemOrderDelivered from '../../components/carditems/ItemOrderDelivered';

import AdminNavbar from '../../components/navbar/AdminNavbar';
import AdminSidebar from '../../components/sidebar/AdminSidebar';

import HomeTable from '../../components/tables/HomeTable';

export default function Home() {
  return (
    <>
    <AdminNavbar />
    <AdminSidebar />
          <div className="py-20 pl-dash">
            <div className="grid grid-cols-4 gap-4">
              <ItemOrderWait />
              <ItemOrderConfirm />
              <ItemOrderDelivery />
              <ItemOrderDelivered />
            </div>
            <div className="py-12">
              <div className="text-3xl font-bold">
                  รายการทั้งหมดที่ขายได้
              </div>
              <div className="py-4">
                <HomeTable />
              </div>
            </div>
          </div>
    </>
  );
}

