import React, { useEffect, useState } from "react";
import "./Review.css";
import Dropdown from "./Dropdown";
import Add from "./add.svg";
import { createPopper } from "@popperjs/core";
import ItemCard from "./ItemCard";
import { Listbox, Transition } from "@headlessui/react";
import Review from "./Review";
import Box from "@mui/material/Box";
import Rating from "@mui/material/Rating";
import Typography from "@mui/material/Typography";

export default function Modal(props) {
  const [review, setReview] = useState(null);
  const [point, setPoint] = useState(null);

  const OnchangeReview = (e) => {
    console.log(e.target.value);
    setReview(e.target.value);
  };

  const onAddNewPoint = (newPoint) => {
    console.log("ข้อมูลมาจาก point = ", typeof newPoint);
    setPoint(newPoint);
  };

  const saveReview = (e) => {
    e.preventDefault();
    console.log("บันทึกข้อมูลเรียบร้อย");
    const reviewData = {
      id: 5,
      desc: review,
      point: point,
    };
    setReview(null);
    setPoint(null);
    props.onAddReview(reviewData);
    console.log(reviewData);
  };

  const [showModal, setShowModal] = useState(false);
  return (
    <>
      <button
        //className="bg-theme text-white active:bg-pink-600 font-bold uppercase text-sm px-6 py-3 rounded shadow hover:shadow-lg outline-none focus:outline-none mr-1 mb-1 ease-linear transition-all duration-150"
        className="text-black font-body border-dashed md:border-dashed border-4"
        type="button"
        onClick={() => setShowModal(true)}
      >
        <div className="p-5 flex flex-col items-center font-bold leading-snug  font-theme ">
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
                <div className = "px-4"><Dropdown /></div>

                {/*body*/}
                <div className="bgg-theme relative px-6 py-2">
                  <div className="my-2 text-lg leading-relaxed flex flex-col  items-center ">
                  <StarRating onAddPoint={onAddNewPoint} />
                    <form onSubmit={saveReview} className = "flex flex-col  items-center">
                      <textarea className = "border rounded-md"
                        id="review"
                        name="review"
                        type="text"
                        value={review}
                        onChange={OnchangeReview}
                        placeholder="รีวิว"
                      />
                      <div className = "py-3 object-center">
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
                </div>
                {/*footer*/}
                {/* <div className="bgg-theme flex items-center justify-end p-6 border-t border-solid border-blueGray-200 rounded-b">
                </div> */}
              </div>
            </div>
          </div>
          <div className="bgg-theme opacity-25 fixed inset-0 z-40"></div>
        </>
      ) : null}
    </>
  );
}

function StarRating(props) {
  const [point, setPoint] = React.useState(0);
  return (
    <Box
      sx={{
        "& > legend": { mt: 2 },
      }}
    >
      <Rating
        name="simple-controlled"
        point={point}
        onChange={(event, newPoint) => {
          setPoint(newPoint);
          // console.log(newPoint);
          props.onAddPoint(newPoint);
        }}
      />
    </Box>
  );
}
