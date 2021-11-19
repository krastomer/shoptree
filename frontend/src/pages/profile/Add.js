import react, { useRef, useState } from "react";
import Add from "../review/add.svg";
import { useForm } from "react-hook-form";
import allLocation from "./allLocation";
export default function AddAddress() {
  const [showModal, setShowModal] = useState(false);
  const onSubmit = (data) => {
    alert(JSON.stringify(data))
    allLocation.push({
      id:allLocation.length+1,
      name:data.name,
      disFirst: data.address1,
      disSecond: data.address2,
      postNumber: data.post,
      phoneNumber: data.mobile_number,
    })
  };
  const {
    register,
    formState: { errors },
    handleSubmit,
    watch,
  } = useForm({});
  const subder = useRef({});
  subder.current = watch("Name", "");
  return (
    <>
      <button
        className="text-black font-body border-dashed md:border-dashed border-4"
        type="button"
        onClick={() => {
          setShowModal(true);
        }}
      >
        <div className="p-5 flex flex-col items-center font-bold leading-snug  font-theme ">
          เพิ่มที่อยู่
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
                <div className="rounded-md shadow-sm -space-y-px">
                  <p>กรอกที่อยู่</p>
                  <div>
                    <form onSubmit={(e) => e.preventDefault()}>
                      <div>
                        <input
                          name="Name"
                          type="text"
                          placeholder="ชื่อ - นามสกุล"
                          {...register("name", {
                            required: true,
                            minLength: {
                              value: 1,
                              message: "write your name",
                            },
                            maxLength: 80,
                          })}
                          className="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-t-md focus:outline-none focus:ring-green-500 focus:border-green-500 focus:z-10 sm:text-sm"
                        />
                        <div>
                          <textarea
                            className="border rounded-md"
                            placeholder="บ้านเลขที่ ซอย ตำบล"
                            {...register("address1", {})}
                          />
                        </div>
                        <div>
                          <textarea
                            className="border rounded-md"
                            placeholder="อำเภอ จังหวัด"
                            {...register("address2", {})}
                          />
                        </div>
                        <input
                          type="tel"
                          placeholder="ไปรษณีย์"
                          {...register("post", {
                            required: true,
                            pattern: {
                              value: /\d+/,
                              message: "must number",
                            },
                            minLength: {
                              value: 5,
                              message: "must have 5 number",
                            },
                            maxLength: {
                              value: 5,
                              message: "must have 5 number",
                            },
                          })}
                          className="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-b-md focus:outline-none focus:ring-green-500 focus:border-green-500 focus:z-10 sm:text-sm"
                        />
                        <input
                          type="tel"
                          placeholder="เบอร์โทรศัพท์"
                          {...register("mobile_number", {
                            required: true,
                            pattern: {
                              value: /\d+/,
                              message: "must phone number",
                            },
                            minLength: {
                              value: 8,
                              message: "Phone number must have 8 -12",
                            },
                            maxLength: {
                              value: 12,
                              message: "Phone number must have 8 -12",
                            },
                          })}
                          className="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-b-md focus:outline-none focus:ring-green-500 focus:border-green-500 focus:z-10 sm:text-sm"
                        />
                      </div>
                      <input
                        className="group relative w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white btn-theme hover:bg-yellow-00 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500"
                        type="submit"
                        value={`คลิกเพื่อเพิ่มขนาดหำของคุณ${subder.current}`}
                        onClick={handleSubmit(onSubmit)}
                      />
                    </form>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div className="bgg-theme opacity-25 fixed inset-0 z-40"></div>
        </>
      ) : null}
    </>
  );
}
