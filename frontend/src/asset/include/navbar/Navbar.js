import { Fragment } from "react";
import { Disclosure, Menu, Transition } from "@headlessui/react";
import Applogo from "../../LogoBanner.png";
import "./Navbar.css";
import React, { Component } from "react";
import { AiOutlineShopping } from "react-icons/ai";
import { MdOutlineReviews } from "react-icons/md";

const navigation = [
  {
    name: <MdOutlineReviews style={{ color: "black", fontSize: "50px" }} />,
    href: "#",
    current: false,
  },
  { name: <AiOutlineShopping />, href: "#", current: false },
];

function classNames(...classes) {
  return classes.filter(Boolean).join(" ");
}

// export default function Example() {
//   return (
//     <Disclosure as="nav" className="navbar-theme">
//       {({ open }) => (
//         <>
//           <div className="max-w-7xl mx-auto px-2 sm:px-6 lg:px-8">
//             <div className="relative flex items-center justify-between h-16">
//               <div className="absolute inset-y-0 left-0 flex items-center sm:hidden">

//                 {/* Mobile menu button*/}
//                 <Disclosure.Button className="inline-flex items-center justify-center p-2 rounded-md text-gray-400 hover:text-white hover:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-white">
//                   <span className="sr-only">Open main menu</span>
//                 </Disclosure.Button>
//               </div>
//               <div className="flex-1 flex items-center justify-center sm:items-stretch sm:justify-start">
//                 <div className="flex-shrink-0 flex items-center">
//                   <img
//                     className="block lg:hidden h-8 w-auto"
//                     src={Applogo}
//                     alt="Workflow"
//                   />

//                   <img
//                     className="hidden lg:block h-8 w-auto"
//                     src={Applogo}
//                     alt="Workflow"
//                   />
//                 </div>

//               <div class="shadow flex">
//                 <input
//                   class="hidden lg:block w-full rounded p-2 bg-green-100 border border-gray-300 text-gray-900"
//                   type="text"
//                   placeholder="Search..."
//                 />
//                 <button class="hidden lg:block bg-white w-auto flex justify-end items-center text-blue-500 p-2 hover:text-blue-400">
//                   <i class="material-icons">search</i>
//                 </button>
//               </div>

//                 <div className="hidden sm:block sm:ml-6 ">
//                   <div className="flex space-x-4">
//                     {navigation.map((item) => (
//                       <a
//                         key={item.name}
//                         href={item.href}
//                         className={classNames(
//                           item.current
//                             ? "bg-gray-900 text-white"
//                             : "text-gray-300 hover:bg-gray-700 hover:text-white",
//                           "px-10 py-5 rounded-md text-sm font-medium"
//                         )}
//                         aria-current={item.current ? "page" : undefined}
//                       >
//                         {item.name}
//                       </a>
//                     ))}
//                   </div>
//                 </div>
//               </div>

//               <div className="absolute inset-y-0 right-0 flex items-center pr-2 sm:static sm:inset-auto sm:ml-6 sm:pr-0">
//                 <button
//                   type="button"
//                   className="bg-gray-800 p-1 rounded-full text-gray-400 hover:text-white focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-gray-800 focus:ring-white"
//                 >
//                   <span className="sr-only">View notifications</span>
//                 </button>

//                 {/* Profile dropdown */}
//                 <Menu as="div" className="ml-3 relative">
//                   <div>
//                     <Menu.Button className=" flex text-sm rounded-full focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-gray-800 focus:ring-white">
//                       <span className="sr-only">Open user menu</span>
//                       {/*<img
//                         className="h-8 w-8 rounded-full"
//                         src="https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80"
//                         alt=""
//                       />*/}

//                       <svg
//                         color="green"
//                         class="h-8 w-8 text-green-500"
//                         viewBox="0 0 24 24"
//                         fill="none"
//                         stroke="currentColor"
//                         stroke-width="2"
//                         stroke-linecap="round"
//                         stroke-linejoin="round"
//                       >
//                         {" "}
//                         <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2" />{" "}
//                         <circle cx="12" cy="7" r="4" />
//                       </svg>
//                     </Menu.Button>
//                   </div>
//                   <Transition
//                     as={Fragment}
//                     enter="transition ease-out duration-100"
//                     enterFrom="transform opacity-0 scale-95"
//                     enterTo="transform opacity-100 scale-100"
//                     leave="transition ease-in duration-75"
//                     leaveFrom="transform opacity-100 scale-100"
//                     leaveTo="transform opacity-0 scale-95"
//                   >
//                     <Menu.Items className="origin-top-right absolute right-0 mt-2 w-48 rounded-md shadow-lg py-1 bg-white ring-1 ring-black ring-opacity-5 focus:outline-none">
//                       <Menu.Item>
//                         {({ active }) => (
//                           <a
//                             href="#"
//                             className={classNames(
//                               active ? "bg-gray-100" : "",
//                               "block px-4 py-2 text-sm text-gray-700"
//                             )}
//                           >
//                             บัญชีของฉัน
//                           </a>
//                         )}
//                       </Menu.Item>
//                       <Menu.Item>
//                         {({ active }) => (
//                           <a
//                             href="#"
//                             className={classNames(
//                               active ? "bg-gray-100" : "",
//                               "block px-4 py-2 text-sm text-gray-700"
//                             )}
//                           >
//                             คำสั่งซื้อของฉัน
//                           </a>
//                         )}
//                       </Menu.Item>
//                       <Menu.Item>
//                         {({ active }) => (
//                           <a
//                             href="#"
//                             className={classNames(
//                               active ? "bg-gray-100" : "",
//                               "block px-4 py-2 text-sm text-gray-700"
//                             )}
//                           >
//                             ออกจากระบบ
//                           </a>
//                         )}
//                       </Menu.Item>
//                     </Menu.Items>
//                   </Transition>
//                 </Menu>
//               </div>
//             </div>
//           </div>

//           <Disclosure.Panel className="sm:hidden">
//             <div className="px-2 pt-2 pb-3 space-y-1">
//               {navigation.map((item) => (
//                 <Disclosure.Button
//                   key={item.name}
//                   as="a"
//                   href={item.href}
//                   className={classNames(
//                     item.current
//                       ? "bg-gray-900 text-white"
//                       : "text-gray-300 hover:bg-gray-700 hover:text-white",
//                     "block px-3 py-2 rounded-md text-base font-medium"
//                   )}
//                   aria-current={item.current ? "page" : undefined}
//                 >
//                   {item.name}
//                 </Disclosure.Button>
//               ))}
//             </div>
//           </Disclosure.Panel>
//         </>
//       )}
//     </Disclosure>
//   );
// }

export default function Example() {
  return (
    <nav class="navbar-theme border-gray-200 px-2 md:px-4 lg:px-5">
      <div class="container mx-auto flex flex-wrap items-center justify-between">
        <a
          href="#"
          class="flex-1 flex items-center justify-center sm:items-stretch sm:justify-start"
        >
          <img
            className="block lg:hidden h-8 w-auto"
            src={Applogo}
            alt="Workflow"
          />
          <img
            className="hidden lg:block h-8 w-auto"
            src={Applogo}
            alt="Workflow"
          />
        </a>
        <div class="flex md:order-2">
          <div class="relative mr-3 md:mr-0 hidden md:block">
            <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
              <svg
                class="w-5 h-5 text-gray-500"
                fill="currentColor"
                viewBox="0 0 20 20"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  fill-rule="evenodd"
                  d="M8 4a4 4 0 100 8 4 4 0 000-8zM2 8a6 6 0 1110.89 3.476l4.817 4.817a1 1 0 01-1.414 1.414l-4.816-4.816A6 6 0 012 8z"
                  clip-rule="evenodd"
                ></path>
              </svg>
            </div>
            <div class="shadow flex">
              <input
                class="hidden lg:block w-full rounded p-2 bg-green-100 border border-gray-300 text-gray-900"
                type="text"
                placeholder="Search..."
              />
              <button class="hidden lg:block bg-white w-auto flex justify-end items-center text-blue-500 p-2 hover:text-blue-400">
                <i class="material-icons">search</i>
              </button>
            </div>
          </div>
          <button
            data-collapse-toggle="mobile-menu-3"
            type="button"
            class="md:hidden text-gray-400 hover:text-gray-900 focus:outline-none focus:ring-2 focus:ring-blue-300 rounded-lg inline-flex items-center justify-center"
            aria-controls="mobile-menu-3"
            aria-expanded="false"
          >
            <span class="sr-only">Open main menu</span>
            <svg
              class="w-6 h-6"
              fill="currentColor"
              viewBox="0 0 20 20"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                fill-rule="evenodd"
                d="M3 5a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM3 10a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM3 15a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1z"
                clip-rule="evenodd"
              ></path>
            </svg>
            <svg
              class="hidden w-6 h-6"
              fill="currentColor"
              viewBox="0 0 20 20"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                fill-rule="evenodd"
                d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z"
                clip-rule="evenodd"
              ></path>
            </svg>
          </button>
        </div>
      </div>
    </nav>
  );
}

// export default function Example() {
//   return (
//     <nav className="bg-white shadow dark:bg-gray-800">
//   <div
//     classNames="
//       container
//       px-6
//       py-3
//       mx-auto
//       md:flex md:justify-between md:items-center
//     "
//   >
//     <div className="flex items-center justify-between">
//       <div>
//         <a
//           className="
//             text-xl
//             font-bold
//             text-gray-800
//             dark:text-white
//             md:text-2xl
//             hover:text-gray-700
//             dark:hover:text-gray-300
//           "
//           href="#"
//           >Brand</a
//         >
//       </div>

//       {/* <!-- Mobile menu button --> */}
//       <div className="flex md:hidden">
//         <button
//           type="button"
//           className="
//             text-gray-500
//             dark:text-gray-200
//             hover:text-gray-600
//             dark:hover:text-gray-400
//             focus:outline-none focus:text-gray-600
//             dark:focus:text-gray-400
//           "
//           aria-label="toggle menu"
//         >
//           <svg viewBox="0 0 24 24" class="w-6 h-6 fill-current">
//             <path
//               fill-rule="evenodd"
//               d="M4 5h16a1 1 0 0 1 0 2H4a1 1 0 1 1 0-2zm0 6h16a1 1 0 0 1 0 2H4a1 1 0 0 1 0-2zm0 6h16a1 1 0 0 1 0 2H4a1 1 0 0 1 0-2z"
//             ></path>
//           </svg>
//         </button>
//       </div>
//     </div>

//     {/* <!-- Mobile Menu open: "block", Menu closed: "hidden" --> */}
//     <div className="items-center md:flex">
//       <div className="flex flex-col md:flex-row md:mx-6">
//         <a
//           className="
//             my-1
//             text-gray-700
//             dark:text-gray-200
//             hover:text-indigo-500
//             dark:hover:text-indigo-400
//             md:mx-4 md:my-0
//           "
//           href="#"
//           >Home</a
//         >
//         <a
//           className="
//             my-1
//             text-gray-700
//             dark:text-gray-200
//             hover:text-indigo-500
//             dark:hover:text-indigo-400
//             md:mx-4 md:my-0
//           "
//           href="#"
//           >Shop</a
//         >
//         <a
//           className="
//             my-1
//             text-gray-700
//             dark:text-gray-200
//             hover:text-indigo-500
//             dark:hover:text-indigo-400
//             md:mx-4 md:my-0
//           "
//           href="#"
//           >Contact</a
//         >
//         <a
//           className="
//             my-1
//             text-gray-700
//             dark:text-gray-200
//             hover:text-indigo-500
//             dark:hover:text-indigo-400
//             md:mx-4 md:my-0
//           "
//           href="#"
//           >About</a
//         >
//       </div>

//       <div className="flex justify-center md:block">
//         <a
//           className="
//             relative
//             text-gray-700
//             dark:text-gray-200
//             hover:text-gray-600
//             dark:hover:text-gray-300
//           "
//           href="#"
//         >
//           <i className="fas fa-shopping-cart"></i>

//           <span
//             className="
//               absolute
//               top-0
//               left-0
//               p-1
//               text-xs text-white
//               bg-indigo-500
//               rounded-full
//             "
//           ></span>
//         </a>
//       </div>
//     </div>
//   </div>
// </nav>
//   );
// }
