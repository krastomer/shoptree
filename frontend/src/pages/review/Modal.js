import React, { useEffect, useState } from "react";
import "./Review.css";
import Dropdown from "./Dropdown";
import StarRating from "./StarRating";
import Add from "./add.svg";
import { createPopper } from "@popperjs/core";
import ItemCard from "./ItemCard";
import { Listbox, Transition } from "@headlessui/react";
import Review from "./Review";

export default function Modal(props) {
  const [review, setReview] = useState(null);

  const OnchangeReview = (e) => {
    console.log(e.target.value);
    setReview(e.target.value);
  };
  const saveReview = (e) => {
    e.preventDefault();
    console.log("บันทึกข้อมูลเรียบร้อย");
    const reviewData = {
      id: 5,
      desc : review,
      point: 5,
    }
    setReview(null)
    props.onAddReview(reviewData)
    console.log(reviewData)
  };

  const [showModal, setShowModal] = React.useState(false);
  return (
    <>
      <button
        //className="bg-theme text-white active:bg-pink-600 font-bold uppercase text-sm px-6 py-3 rounded shadow hover:shadow-lg outline-none focus:outline-none mr-1 mb-1 ease-linear transition-all duration-150"
        className="text-black font-body border-dashed md:border-dashed border-4"
        type="button"
        onClick={() => setShowModal(true)}
      >
        <div className="p-5 flex flex-col  items-center font-bold leading-snug  font-theme ">
          เขียนรีวิวใหม่
          <div className="p-2 flex">
            <img src={Add} alt="Add" />
          </div>
        </div>
      </button>
      {showModal ? (
        <>
          <div className="justify-center items-center flex overflow-x-hidden overflow-y-auto fixed inset-0 z-50 outline-none focus:outline-none">
            <div className="relative w-auto my-6 mx-auto max-w-3xl">
              {/*content*/}
              <div className=" bgg-theme border-0 rounded-lg shadow-lg relative flex flex-col w-full  outline-none focus:outline-none">
                <button
                  className="p-1 ml-auto bg-transparent border-0 text-black float-right text-3xl leading-none font-semibold outline-none focus:outline-none"
                  onClick={() => setShowModal(false)}
                >
                  <div className="text-white bg-transparent h-6 w-6 text-2xl block outline-none focus:outline-none">
                    ×
                  </div>
                </button>
                <Dropdown />

                {/*body*/}
                <div className="bgg-theme relative p-6 flex-auto">
                  <div className="my-4 text-lg leading-relaxed ">
                    <form onSubmit={saveReview}>
                      <input
                        id="review"
                        name="review"
                        type="text"
                        value={review}
                        onChange={OnchangeReview}
                        placeholder="รีวิว"
                      />
                      <div>
                        <button
                          className="submit-theme text-white font-bold uppercase text-sm px-6 py-3 rounded shadow hover:shadow-lg outline-none focus:outline-none mr-1 mb-1 ease-linear transition-all duration-150"
                          type="submit"
                          // onClick={() => setShowModal(false)}
                        >
                          เพิ่มรีวิว
                        </button>
                      </div>
                    </form>
                  </div>
                  <StarRating />
                </div>
                {/*footer*/}
                <div className="bgg-theme flex items-center justify-end p-6 border-t border-solid border-blueGray-200 rounded-b">
                  {/* <button
                    className="cancle-theme background-transparent font-bold uppercase px-6 py-2 text-sm outline-none focus:outline-none mr-1 mb-1 ease-linear transition-all duration-150"
                    type="button"
                    onClick={() => setShowModal(false)}
                  >
                    Close
                  </button> */}
                </div>
              </div>
            </div>
          </div>
          <div className="opacity-25 fixed inset-0 z-40 bg-black"></div>
        </>
      ) : null}
    </>
  );
}
