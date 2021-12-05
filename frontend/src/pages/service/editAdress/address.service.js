import axios from "axios";
const API_URL = "http://spaceship.trueddns.com:23720";

export const address = async (
  name,
  phoneNumber,
  address_line,
  state,
  city,
  district,
  postal_code
) => {
  let response;
  const data = {
    name: name,
    phone_number: phoneNumber,
    address_line: address_line,
    country: "Thailand",
    state: state,
    city: city,
    district: district,
    postal_code: postal_code,
  };
  const config = {
    method: "post",
    url: API_URL + "/api/v1/customers/addresses",
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
