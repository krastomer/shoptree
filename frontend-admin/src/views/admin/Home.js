import React from 'react';

import AdminNavbar from '../../components/navbar/AdminNavbar';
import AdminSidebar from '../../components/sidebar/AdminSidebar';
import HeaderStats from "../../components/Headers/HeaderStats.js";

export default function Home() {
  return (
    <>
          <div className="px-4 md:px-10 mx-auto w-full">
              <HeaderStats />
          </div>
    </>
  );
}

