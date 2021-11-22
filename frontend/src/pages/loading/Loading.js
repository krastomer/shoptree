import Applogo from "../../LogoBanner.png";


export default function Loading() {
  return (

      <div className = "flex flex-col items-center justify-center h-screen ">
        <div className = "w-40 h-40 border-t-4 border-b-4 border-green-900 rounded-full animate-spin"></div>
        <div className = "w-40 h-40">
          <img src={Applogo} alt="Logo" />
        </div>
      </div>
  );
}
