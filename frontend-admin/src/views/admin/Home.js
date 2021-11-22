import React from 'react';

import AdminNavbar from '../../components/navbar/AdminNavbar';
import AdminSidebar from '../../components/sidebar/AdminSidebar';

import ItemOrderWait from '../../components/carditems/ItemOrderWait';
import ItemOrderConfirm from '../../components/carditems/ItemOrderConfirm';
import ItemOrderDelivery from '../../components/carditems/ItemOrderDelivery';
import ItemOrderDelivered from '../../components/carditems/ItemOrderDelivered';

export default function Home() {
  return (
    <div className="font-body">
        <AdminNavbar />
        <AdminSidebar />       
        <div className="py-16 pl-dash"> 
            <div className="grid w-full grid-cols-4 gap-4">
                <ItemOrderWait />
                <ItemOrderConfirm />
                <ItemOrderDelivery />
                <ItemOrderDelivered />
            </div>
        </div>
    </div>
  );
}

