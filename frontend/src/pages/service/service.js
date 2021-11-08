export function getProduct(item) {
    var requestOptions = {
      method: "GET",
      redirect: "follow",
    };
  
    fetch(`${Api_Url}/api/v1/products/1`, requestOptions)
      .then((response) => response.text())
      .then((result) => console.log(result))
      .catch((error) => console.log("error", error));
  }