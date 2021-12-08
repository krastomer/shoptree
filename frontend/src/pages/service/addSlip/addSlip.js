import axios from "axios";

export const addSlip = async (image) => {
  let response;
  const config = {
    method: "post",
    url: `http://spaceship.trueddns.com:23720/api/v1/orders/payment`,
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
    },
    data:image,
  };
  try {
    response = await axios(config);
  } catch (error) {
    console.error(error);
  }
  return response?.data ? response?.data : null; // or set initial value
};
