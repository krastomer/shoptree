import "./Error404.css";
import Applogo from "../../LogoSVGBrown.svg";
// import Applogo from "../../LogoBanner.png";

export default function Loading() {
  return (
    <div className="flex flex-col items-center justify-center h-screen ">
      <div className=" w-70 h-70">
        <img src={Applogo} alt="Logo" />
      </div>
      <div className="mt-6 font-body font-bold text-center text-4xl">
        <div>Error 404</div>
        <div>Page Not Found</div>
      </div>
      <div><a href = "/"className="mt-3 group relative flex justify-center py-2 px-4 border border-transparent text-sm font-body font-medium rounded-md text-white btn-theme hover:bg-yellow-00 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500">
        กลับสู่หน้าหลัก
      </a>
      </div>
    </div>
  );
}
