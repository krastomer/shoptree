/* This example requires Tailwind CSS v2.0+ */
import { Fragment } from 'react'
import { HiHome,HiLogout,HiTruck } from "react-icons/hi";
import { Menu, Transition } from '@headlessui/react'
import User from "./user.svg";
import { useDispatch, useSelector } from "react-redux";
import react , {useCallback} from 'react';

function classNames(...classes) {
  return classes.filter(Boolean).join(' ')
}

export default function UserDropdown() {
  const dispatch = useDispatch();
  const user = JSON.parse(localStorage.getItem("user"));

  return (
    <Menu as="div" className="relative inline-block text-right font-body">
      <div>
        <Menu.Button className="flex items-center px-3 py-2 text-xs font-bold leading-snug text-green-600 uppercase hover:opacity-75">
        <img src={User} alt="User" />
        </Menu.Button>
      </div>
      <Transition
        as={Fragment}
        enter="transition ease-out duration-100"
        enterFrom="transform opacity-0 scale-95"
        enterTo="transform opacity-100 scale-100"
        leave="transition ease-in duration-75"
        leaveFrom="transform opacity-100 scale-100"
        leaveTo="transform opacity-0 scale-95"
      >
        <Menu.Items className="absolute right-0 mt-2 origin-top-right rounded-md shadow-lg w-36 bg-theme ring-1 ring-black ring-opacity-5 focus:outline-none">
          <div className="py-1 text-center">
            <Menu.Item>
              {({ active }) => (
                <a
                  href="/login"
                  className={classNames(
                    active ? 'text-green-500' : 'text-white',
                    'px-4 py-2 text-sm flex'
                  )}
                >
                <HiHome />
                &nbsp;
                เข้าสู่ระบบ
                </a>
              )}
            </Menu.Item>
          </div>
        </Menu.Items>
      </Transition>
    </Menu>
  )
}