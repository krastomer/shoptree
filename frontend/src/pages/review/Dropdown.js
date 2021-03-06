import React from "react";
import "./Review.css";
import "./AllReview";
import { Listbox, Transition } from "@headlessui/react";
import { CheckIcon, SelectorIcon } from '@heroicons/react/outline'
import { Fragment, useState, useEffect } from 'react'
import { getReviewNumbers } from "../service/review/getOrderCustomer";

const numberOrder = getReviewNumbers();

const labelText = [
    {
      id: "ยังไม่ได้เลือก"
    },

  ]
  
  function classNames(...classes) {
    return classes.filter(Boolean).join(' ')
  }
  
export default function Example() {
  const [selected, setSelected] = useState(labelText[0])

  const [numOrder, setNumOrder] = useState(null);

  useEffect(() => {
    numberOrder.then(function (data) {
      setNumOrder(data.data);
    });
  });
  if(!numOrder) return null;
  
    return (
      <Listbox value={selected} onChange={setSelected}>
        {({ open }) => (
          <>
            <Listbox.Label className="block font-medium modal-theme ">เลขคำสั่งซื้อสินค้า</Listbox.Label>
            <div className="relative mt-1 ">
              <Listbox.Button className="relative w-full py-2 pl-3 pr-10 text-left border-b-2 rounded-md shadow-sm cursor-default modal-theme focus:outline-none focus:ring-1 focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm">
                <span className="flex items-center">
                  {/* <img src={selected.avatar} alt="" className="flex-shrink-0 w-6 h-6 rounded-full" /> */}
                  <span className="block ml-3 truncate">{selected.id}</span>
                </span>
                <span className="absolute inset-y-0 right-0 flex items-center pr-2 ml-3 pointer-events-none">
                  <SelectorIcon className="w-5 h-5 text-white" aria-hidden="true" />
                </span>
              </Listbox.Button>
  
              <Transition
                show={open}
                as={Fragment}
                leave="transition ease-in duration-100"
                leaveFrom="opacity-100"
                leaveTo="opacity-0"
              >
                <Listbox.Options className="absolute z-10 w-full py-1 mt-1 overflow-auto text-base bg-white rounded-md shadow-lg max-h-56 ring-1 ring-black ring-opacity-5 focus:outline-none sm:text-sm">
                  {numOrder.map((person) => (
                    <Listbox.Option
                      key={person.id}
                      className={({ active }) =>
                        classNames(
                          active ? 'text-white bg-indigo-600' : 'text-gray-900',
                          'cursor-default select-none relative py-2 pl-3 pr-9'
                        )
                      }
                      value={person}
                    >
                      {({ selected, active }) => (
                        <>
                          <div className="flex items-center">
                            {/* <img src={person.avatar} alt="" className="flex-shrink-0 w-6 h-6 rounded-full" /> */}
                            <span
                              className={classNames(selected ? 'font-semibold' : 'font-normal', 'ml-3 block truncate')}
                            >
                              {person.id}
                            </span>
                          </div>
  
                          {selected ? (
                            <span
                              className={classNames(
                                active ? 'text-white' : 'text-indigo-600',
                                'absolute inset-y-0 right-0 flex items-center pr-4'
                              )}
                            >
                              <CheckIcon className="w-5 h-5" aria-hidden="true" />
                            </span>
                          ) : null}
                        </>
                      )}
                    </Listbox.Option>
                  ))}
                </Listbox.Options>
              </Transition>
            </div>
          </>
        )}
      </Listbox>
    )
  }