import axios from "axios";
import * as fs from 'fs';

export const addSlip = async (image) => {
  let response;
  const data = new FormData();
  console.log(image)
  let localimage = fs.createReadStream("0LYHscXWQ/w644.jpeg");
  data.append(
    "image",
    image
  );
  const config = {
    method: "post",
    url: `http://spaceship.trueddns.com:23720/api/v1/orders/payment`,
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
    },
    data: data,
  };
  try {
    response = await axios(config);
  } catch (error) {
    console.error(error);
  }
  return response?.data ? response?.data : null; // or set initial value
};
