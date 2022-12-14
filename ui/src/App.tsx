import React from 'react';
import logo from './logo.svg';
import './App.css';
import {BrowserRouter, HashRouter, Link, Route, Routes} from "react-router-dom";
import {PasswdInfo} from "./pc/PasswdInfo";
import {PasswdInfo as MPasswdInfo} from "./mobile/PasswdInfo";
import {Login as MLogin} from "./mobile/Login"
import {Register as MRegister} from "./mobile/Register"
import RouterUtil from "./common/utils/RouterUtil";
import {Login} from "./pc/Login";
import StorageUtil from "./common/utils/StorageUtil";

class App extends React.Component<any, any> {

    constructor(props: any) {
        super(props);
        this.state = {}
    }

    componentDidMount() {
        if (document.location.hash === "#/") {
            if ((navigator.userAgent.match(/(phone|pad|pod|iPhone|iPod|ios|iPad|Android|Mobile|BlackBerry|IEMobile|MQQBrowser|JUC|Fennec|wOSBrowser|BrowserNG|WebOS|Symbian|Windows Phone)/i))) {
                //手机
                RouterUtil.push("/m/login")
            } else {
                //电脑
                if (StorageUtil.get("autoLogin") === "true") {
                    RouterUtil.push("/passwdInfo")
                } else
                    RouterUtil.push("/login")
            }
        }
    }

    render() {
        return (
            <HashRouter>
                <Routes>
                    {/*  <Route component={} path={"/login"}/>*/}
                    <Route path={"/passwdInfo"} element={<PasswdInfo/>}></Route>
                    <Route path={"/login"} element={<Login/>}></Route>
                    <Route path={"/m/login"} element={<MLogin/>}></Route>
                    <Route path={"/m/register"} element={<MRegister/>}></Route>
                    <Route path={"/m/passwdInfo"} element={<MPasswdInfo/>}></Route>
                </Routes>
            </HashRouter>
        );
    }

}

export default App;
