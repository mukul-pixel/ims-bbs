import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import axios from "axios";

export const Login = () => {
    const [email,setEmail] = useState("");
    const [password,setPassword] = useState("");
    const navigate = useNavigate();

    const handleSubmit = async(e)=> {
        e.preventDefault();

        try {
            const response = await axios.post('http://localhost:8080/api/v1/login',{
                email,
                password
            });
            if (response.status === 202) {  // 202 is the status for successful login
                console.log('Login successful:', response.data);
                localStorage.setItem("token",response.data.token)
                localStorage.setItem("userId",response.data.userID)
                navigate('/createGRN');  // Navigate to the dashboard on successful login
            } else {
                console.log('Login failed:', response.data);
            }
        } catch (error) {
            console.error('Error during login:', error.response ? error.response.data : error.message);
        }
    }
  return (
    <div className="vh-100 wh-100 row">
      <div className="loginDisplayImage col-9 d-none d-md-block">
        <ul>
            <li>handle error at front-end using toastify</li>
        </ul>
      </div>
      <div className="formData col-md-3 col-12 p-0 text-center border-start">
        <div className="py-5 fw-bold">INVENTORY MANANGEMENT SYSTEM</div>
        <form className="manualLogin pe-3 ps-4 py-5" onSubmit={handleSubmit}>
          <div className="form-floating mb-2">
            <input
              type="email"
              className="form-control"
              id="floatingInput"
              placeholder="name@example.com"
              onChange={(e)=>setEmail(e.target.value)}
              required
            />
            <label htmlFor="floatingInput" style={{fontSize:"smaller"}}>Email address</label>
          </div>
          <div className="form-floating">
            <input
              type="password"
              className="form-control"
              id="floatingPassword"
              placeholder="Password"
              onChange={(e)=>setPassword(e.target.value)}
              required
            />
            <label htmlFor="floatingPassword" style={{fontSize:"smaller"}}>Password</label>
          </div>
          <button type="submit" className="loginButton w-100 btn my-3">
            <span>Continue</span>
          </button>
          <div className="dividingManualAndAutoLogin row my-2">
            <span className="col-5"><hr/></span>
            <span className="col-2 text-center mt-1" style={{fontSize:"smaller", color:"gray"}}> OR </span>
            <span className="col-5"><hr/></span>
          </div>
          <button className="googleLoginButton row w-100 btn my-3">
            <span className="col-2"><img src="https://www.pngmart.com/files/16/official-Google-Logo-PNG-Image.png" height="25" alt="googleLogo"/></span>
            <span className="col-10">Continue with Google</span>
          </button>
        </form>
      </div>
    </div>
  );
};
