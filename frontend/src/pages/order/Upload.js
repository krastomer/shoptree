import React, { useState } from "react";
import Add from "../review/add.svg";
import { addSlip } from "../service/addSlip/addSlip";

export default function Upload() {
  const [image, setImage] = useState({ preview: "", raw: "" });
  const [file , setFile] = useState(null)
  const handleChange = (e) => {
    let file = e.target.files[0];
    setFile(e.target.files[0]);
    if (e.target.files.length) {
      setImage({
        preview: URL.createObjectURL(e.target.files[0]),
        raw: e.target.files[0],
      });
    }
  };
 
  const handleUpload = async (e) => {
    e.preventDefault();
<<<<<<< HEAD
    
=======
    console.log("file ", file);
    const status = await addSlip(file)
>>>>>>> be65979f1c40de1212aacdc7780dca9ecaa4ea96
  };
  return (
    <div>
      <h2 className="py-4 text-2xl tracking-tight text-gray-600">
        ส่งหลักฐานการโอนเงิน
      </h2>
      <label htmlFor="upload-button">
        {image.preview ? (
          <div>
            รูปภาพปัจจุบัน
            <img src={image.preview} alt="dummy" width="300" height="300" />
            คลิกรูปภาพอีกครั้งเพื่อเปลี่ยนรูปภาพ
          </div>
        ) : (
          <>
            <div
              className="text-black border-4 border-dashed font-body md:border-dashed"
              type="file"
              onChange={handleChange}
            >
              <div className="flex flex-col items-center p-5 font-bold leading-snug font-theme ">
                อัพโหลดสลิป
                <div className="flex p-2">
                  <img src={Add} alt="dummy" />
                </div>
              </div>
            </div>
          </>
        )}
      </label>
      <input
        type="file"
        id="upload-button"
        style={{ display: "none" }}
        onChange={handleChange}
      />
      <br />
      <button onClick={handleUpload}>Upload</button>
    </div>
  );
}
