import React, { useEffect} from "react";
import { useDispatch,useSelector } from "react-redux";
import { Router, Route, Switch } from "react-router-dom";
import Login from "./pages/login/Login";
import Register from "./pages/register/Register";
import Home from "./pages/home/Home";
import Review from "./pages/review/Review";
import Modal from "./pages/review/Modal";
import Order from "./pages/order/Order";
import Navbar from "./asset/include/navbar/Navbar";
import { Provider } from "react-redux";
import Detail from "./pages/home/Detailproduct";
import allProduct from "./pages/home/allProduct";
import Dropdown from "./pages/review/Dropdown";
import StarRating from "./pages/review/StarRating";
import { LoginUser } from "./models/User";
import { SuccessOrder } from "./pages/success/success";
import Profile from "./pages/profile/Profile";
import Loading from "./pages/loading/Loading";
import Error404 from "./pages/error404/Error404";
import LayoutEdit from "./pages/profileEdit/LayoutEdit";
import Itemprofile from "./pages/itemprofile/Itemprofile";
import { clearMessage } from "./actions/message";

import EventBus from "./common/EventBus";
import { history } from "./helpers/history";
export default function App() {
  const dispatch = useDispatch();

  useEffect(() => {
    history.listen((location) => {
      dispatch(clearMessage()); // clear message when changing location
    });
  }, [dispatch]);

  return (
    <Router history={history}>
      <div className="wrap">
        <Switch>
          <Route path="/" exact>
            <Home />
          </Route>
          <Route path="/error404" exact>
            <Error404 />
          </Route>
          <Route path="/itemprofile" exact>
            <Itemprofile />
          </Route>
          <Route path="/products/:id" exact>
            <Detail id={allProduct}></Detail>
          </Route>
          <Route path="/order" exact>
            <Order/>
          </Route>
          <Route path="/login" exact>
            <Login />
          </Route>
          <Route path="/loading" exact>
            <Loading />
          </Route>
          <Route path="/register" exact>
            <Register />
          </Route>
          <Route path="/review" exact>
            <Review />
          </Route>
          <Route path="/profile" exact>
            <Profile />
          </Route>
          <Route path="/profile/edit" exact>
            <LayoutEdit />
          </Route>
          <Route path="/navbar" exact>
            <Navbar />
          </Route>
          <Route path="/modal" exact>
            <Modal />
          </Route>
          <Route path="/dropdown" exact>
            <Dropdown />
          </Route>
          <Route path="/starrating" exact>
            <StarRating />
          </Route>
          <Route path="/order/:USER/Success" exact>
            <SuccessOrder USER={LoginUser.username}></SuccessOrder>
          </Route>
        </Switch>
      </div>
    </Router>
  );
}
