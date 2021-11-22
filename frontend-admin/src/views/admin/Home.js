import React from 'react';

import AdminNavbar from '../../components/navbar/AdminNavbar';
import AdminSidebar from '../../components/sidebar/AdminSidebar';

export default function Home() {
  return (
    <div className="font-body">
        <AdminNavbar />
        <AdminSidebar />
            ยินดีตอนรับเข้าสู่ชมรม
    </div>
  );
}

