import React, { useEffect, useState } from "react";
import "./Review.css";
import Dropdown from "./Dropdown";
import Add from "./add.svg";
import Box from "@mui/material/Box";
import Rating from "@mui/material/Rating";


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
        className="text-black border-4 border-dashed font-body md:border-dashed"
        type="button"
        onClick={() => setShowModal(true)}
      >
        <div className="flex flex-col items-center p-5 font-bold leading-snug font-theme ">
          เขียนรีวิวใหม่
          <div className="flex p-2">
            <img src={Add} alt="Add" />
          </div>
        </div>
      </button>
      {showModal ? (
        <>
          <div className="fixed inset-0 z-50 flex items-center justify-center overflow-x-hidden overflow-y-auto outline-none focus:outline-none">
            <div className="relative w-auto max-w-3xl mx-auto my-6">
              {/*content*/}
              <div className="relative flex flex-col w-full border-0 rounded-lg shadow-lg outline-none bgg-theme focus:outline-none">
                <button
                  className="float-right p-1 ml-auto text-3xl font-semibold leading-none text-black bg-transparent border-0 outline-none focus:outline-none"
                  onClick={() => setShowModal(false)}
                >
                  <div className="block w-6 h-6 text-2xl text-white bg-transparent outline-none focus:outline-none">
                    ×
                  </div>
                </button>
                <div className = "px-4"><Dropdown /></div>

                {/*body*/}
                <div className="relative px-6 py-2 bgg-theme">
                  <div className="flex flex-col items-center my-2 text-lg leading-relaxed ">
                  <StarRating onAddPoint={onAddNewPoint} />
                    <form onSubmit={saveReview} className = "flex flex-col items-center">
                      <textarea className = "border rounded-md"
                        id="review"
                        name="review"
                        type="text"
                        value={review}
                        onChange={OnchangeReview}
                        placeholder="รีวิว"
                      />
                      <div className = "object-center py-3">
                        <button
                          className="px-6 py-3 mb-1 mr-1 text-sm font-bold text-white uppercase transition-all duration-150 ease-linear rounded shadow outline-none submit-theme hover:shadow-lg focus:outline-none"
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
                {/* <div className="flex items-center justify-end p-6 border-t border-solid rounded-b bgg-theme border-blueGray-200">
                </div> */}
              </div>
            </div>
          </div>
          <div className="fixed inset-0 z-40 opacity-25 bgg-theme"></div>
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
