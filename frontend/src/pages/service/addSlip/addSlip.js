import axios from "axios";

export const addSlip = async (image) => {
  let response;
  var fs = require("fs");
  var data = new FormData();
  console.log(image)
  data.append(
    "image",
    fs.createReadStream(image.raw)
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
