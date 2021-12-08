import react from "react";
import Add from "../review/add.svg";
import { useForm } from "react-hook-form";
import Navbar from "../../asset/include/navbar/Navbar";
import { address } from "../service/editAdress/address.service";
import { getAddress } from "../service/editAdress/getAddress";
import { useHistory } from "react-router-dom";
export default function EditAddress(props) {
  let history = useHistory();
  const onSubmit = async (data) => {
    const status = await address(data.name,data.mobile_number,data.address1,data.subdistrict,data.province,data.district,data.post)
    console.log(status)
    history.push("/profile");
    window.location.reload();
  };
  const {
    register,
    formState: { errors },
    handleSubmit,
    watch,
  } = useForm({});
  return (
    <>
      <div className="flex items-center justify-center min-h-full px-4 py-12 sm:px-6 lg:px-8">
        <div className="w-full max-w-md space-y-8">
          <div>
            <h2 className="mt-6 text-3xl font-extrabold text-left text-gray-900">
              {props.name}
            </h2>
          </div>

          <input type="hidden" name="remember" defaultValue="true" />
          <div className="-space-y-px rounded-md shadow-sm">
            <form onSubmit={(e) => e.preventDefault()}>
              <div>
                <p className="text-gray-500 ">ชื่อ-นามสกุล</p>
                <label htmlFor="password" className="sr-only">
                  Name
                </label>
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
                  className="relative block w-full px-3 py-2 text-gray-900 placeholder-gray-500 border border-gray-300 rounded-none appearance-none rounded-t-md focus:outline-none focus:ring-green-500 focus:border-green-500 focus:z-10 sm:text-sm"
                />
                {errors.Name && <p>{errors.Name.message}</p>}
              </div>
              <div>
                <p className="text-gray-500 ">เบอร์โทรศัพท์</p>
                <label htmlFor="password" className="sr-only">
                  Phone
                </label>
                <input
                  type="tel"
                  placeholder="Mobile number"
                  {...register("mobile_number", {
                    required: true,
                    pattern: {
                      value: /(?=.*\d)/,
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
                  className="relative block w-full px-3 py-2 text-gray-900 placeholder-gray-500 border border-gray-300 rounded-none appearance-none rounded-b-md focus:outline-none focus:ring-green-500 focus:border-green-500 focus:z-10 sm:text-sm"
                />
                <div>
                  {errors.Mobile_number && (
                    <p>{errors.Mobile_number.message}</p>
                  )}
                </div>
              </div>
              <div>
                <p className="text-gray-500 ">ที่อยู่</p>
                <textarea
                  className="relative block w-full px-3 py-2 text-gray-900 placeholder-gray-500 border border-gray-300 rounded-none appearance-none rounded-b-md focus:outline-none focus:ring-green-500 focus:border-green-500 focus:z-10 sm:text-sm"
                  placeholder="บ้านเลขที่ ซอย ถนน"
                  {...register("address1", {})}
                />
              </div>
              <p className="text-gray-500 ">แขวง/ตำบล</p>
              <input
                  name="province"
                  type="text"
                  placeholder="แขวง/ตำบล"
                  {...register("subdistrict", {
                    required: true,
                    minLength: {
                      value: 1,
                      message: "Write your province",
                    },
                    maxLength: 80,
                  })}
                  className="relative block w-full px-3 py-2 text-gray-900 placeholder-gray-500 border border-gray-300 rounded-none appearance-none rounded-t-md focus:outline-none focus:ring-green-500 focus:border-green-500 focus:z-10 sm:text-sm"
                />
                <p className="text-gray-500 ">อำเภอ</p>
              <input
                  name="province"
                  type="text"
                  placeholder="อำเภอ"
                  {...register("district", {
                    required: true,
                    minLength: {
                      value: 1,
                      message: "Write your province",
                    },
                    maxLength: 80,
                  })}
                  className="relative block w-full px-3 py-2 text-gray-900 placeholder-gray-500 border border-gray-300 rounded-none appearance-none rounded-t-md focus:outline-none focus:ring-green-500 focus:border-green-500 focus:z-10 sm:text-sm"
                />
                <p className="text-gray-500 ">จังหวัด</p>
              <input
                  name="province"
                  type="text"
                  placeholder="จังหวัด"
                  {...register("province", {
                    required: true,
                    minLength: {
                      value: 1,
                      message: "Write your province",
                    },
                    maxLength: 80,
                  })}
                  className="relative block w-full px-3 py-2 text-gray-900 placeholder-gray-500 border border-gray-300 rounded-none appearance-none rounded-t-md focus:outline-none focus:ring-green-500 focus:border-green-500 focus:z-10 sm:text-sm"
                />
                <p className="text-gray-500 ">ไปรษณีย์</p>
                 <input
                      type="tel"
                      placeholder="ไปรษณีย์"
                      {...register("post", {
                        required: true,
                        pattern: {
                          value: /\d+/,
                          message: "must number",
                        },
                      })}
                      className="relative block w-full px-3 py-2 text-gray-900 placeholder-gray-500 border border-gray-300 rounded-none appearance-none rounded-b-md focus:outline-none focus:ring-green-500 focus:border-green-500 focus:z-10 sm:text-sm"
                    />
                    <br />
              <p className="text-gray-500 ">ประเทศไทย</p>
              <br />
              <input
                type="submit"
                value="เพิ่มที่จัดส่งใหม่"
                onClick={handleSubmit(onSubmit)}
                className="relative flex justify-center w-full px-4 py-2 text-sm font-medium text-white border border-transparent rounded-md group btn-theme hover:bg-yellow-00 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500"
              ></input>
            </form>
          </div>
        </div>
      </div>
    </>
  );
}
