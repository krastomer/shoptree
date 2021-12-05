/*eslint-disable*/
import React from "react";
import { Link } from "react-router-dom";
export default function Sidebar() {
  const [collapseShow, setCollapseShow] = React.useState("hidden");
  return (
    <>
      <nav className="relative z-10 flex flex-wrap items-center justify-between px-6 py-4 bg-green-700 shadow-xl md:left-0 md:block md:fixed md:top-0 md:bottom-0 md:overflow-y-auto md:flex-row md:flex-nowrap md:overflow-hidden md:w-64">
        <div className="flex flex-wrap items-center justify-between w-full px-0 mx-auto md:flex-col md:items-stretch md:min-h-full md:flex-nowrap">
          {/* Toggler */}
          <button
            className="px-3 py-1 text-xl leading-none text-white bg-transparent border border-transparent border-solid rounded opacity-50 cursor-pointer md:hidden"
            type="button"
            onClick={() => setCollapseShow("bg-white m-2 py-3 px-6")}
          >
            <i className="fas fa-bars"></i>
          </button>
          {/* Brand */}
          <Link
            className="inline-block p-4 px-0 mr-0 text-sm font-bold text-left uppercase md:block md:pb-2 text-blueGray-600 whitespace-nowrap"
            to="/"
          >
             <img
              alt="..."
              className="w-full "
              src={require("../../assets/LogoSVG.svg").default}
            />
          </Link>
          {/* Collapse */}
          <div
            className={
              "md:flex md:flex-col md:items-stretch md:opacity-100 md:relative md:mt-4 md:shadow-none shadow absolute top-0 left-0 right-0 z-40 overflow-y-auto overflow-x-hidden h-auto items-center flex-1 rounded " +
              collapseShow
            }
          >
            {/* Form */}
            <form className="mt-6 mb-4 md:hidden">
              <div className="pt-0 mb-3">
                <input
                  type="text"
                  placeholder="Search"
                  className="w-full h-12 px-3 py-2 text-base font-normal leading-snug bg-white border-0 border-solid rounded shadow-none outline-none border-blueGray-500 placeholder-blueGray-300 text-blueGray-600 focus:outline-none"
                />
              </div>
            </form>
            {/* Heading */}
            <h6 className="block pt-1 text-xs font-bold text-white no-underline uppercase md:min-w-full">
              การจัดการสินค้า
            </h6>
             {/* Divider */}
             <hr className="my-4 md:min-w-full" />
            {/* Navigation */}

            <ul className="flex flex-col list-none md:flex-col md:min-w-full">
              <li className="items-center">
                <Link
                  className={
                    "text-xs uppercase py-3 font-bold block " +
                    (window.location.href.indexOf("/admin/AllItem") !== -1
                      ? "text-white hover:text-lightBlue-600"
                      : "text-white hover:text-blueGray-500")
                  }
                  to="/product"
                >
                  <i
                    className={
                      "fas fa-tv mr-2 text-sm " +
                      (window.location.href.indexOf("/admin/AllItem") !== -1
                        ? "opacity-75"
                        : "text-white")
                    }
                  ></i>{" "}
                  แสดงรายการสินค้าทั้งหมด
                </Link>
              </li>

              <li className="items-center">
                <Link
                  className={
                    "text-xs uppercase py-3 font-bold block " +
                    (window.location.href.indexOf("/admin/AllTypes") !== -1
                    ? "text-white hover:text-lightBlue-600"
                    : "text-white hover:text-blueGray-500")
                  }
                  to="/category"
                >
                  <i
                    className={
                      "fas fa-tools mr-2 text-sm " +
                      (window.location.href.indexOf("/admin/AllTypes") !== -1
                        ? "opacity-75"
                        : "text-blueGray-300")
                    }
                  ></i>{" "}
                  แสดงประเภททั้งหมด
                </Link>
              </li>
            </ul>
            {/* Heading */}
            <h6 className="block pt-1 mt-4 text-xs font-bold text-white no-underline uppercase md:min-w-full">
              การจัดส่งสินค้า
            </h6>
            {/* Navigation */}
             {/* Divider */}
             <hr className="my-4 md:min-w-full" />

            <ul className="flex flex-col list-none md:flex-col md:min-w-full md:mb-4">
              <li className="items-center">
                <Link
                 className={
                  "text-xs uppercase py-3 font-bold block " +
                  (window.location.href.indexOf("/admin/AllItem") !== -1
                  ? "text-white hover:text-lightBlue-600"
                  : "text-white hover:text-blueGray-500")
                }
                to="/confirm"
                >
                  <i className="mr-2 text-sm text-white fas fa-fingerprint"></i>{" "}
                  ยืนยันคำสั่งซื้อ
                </Link>
              </li>

              <li className="items-center">
                <Link
                 className={
                  "text-xs uppercase py-3 font-bold block " +
                  (window.location.href.indexOf("/admin/AllItem") !== -1
                  ? "text-white hover:text-lightBlue-600"
                  : "text-white hover:text-blueGray-500")
                }
                to="/delivery"
                >
                  <i className="mr-2 text-sm fas fa-clipboard-list text-blueGray-300"></i>{" "}
                  แสดงรายการจัดส่ง
                </Link>
              </li>
            </ul>

           
            {/* Heading */}
            <h6 className="block pt-1 mt-4 text-xs font-bold text-white no-underline uppercase md:min-w-full">
              การจัดการทั่วไป
            </h6>
             {/* Divider */}
             <hr className="my-4 md:min-w-full" />
            {/* Navigation */}

            <ul className="flex flex-col list-none md:flex-col md:min-w-full md:mb-4">
              <li className="items-center">
                <Link
                className={
                  "text-xs uppercase py-3 font-bold block " +
                  (window.location.href.indexOf("/admin/AllItem") !== -1
                  ? "text-white hover:text-lightBlue-600"
                  : "text-white hover:text-blueGray-500")
                }
                to="/customer"
                >
                  <i className="mr-2 text-sm text-white fas fa-newspaper"></i>{" "}
                  แสดงรายชื่อลูกค้า
                </Link>
              </li>
            </ul>
          </div>
        </div>
      </nav>
    </>
  );
}
