import React from 'react';
import profile from "../../Profile-login.png";
import "./navbarlogin.css";

const App = () => {
  return (
    <div style={{marginTop: "100px"}}>
      <div className="container">
        <h1 className="text-masuk">Masuk</h1>
            <div className="images"> 
            <img src={profile} alt="" />
            </div>
            <div className="row justify-content-center">
            <div className="col-md-3">
              
            <div className="form-group">
            <label className="text-white">Email</label>
            <input type="email" placeholder="Enter your email" className="form-control"/>
            </div>

            <div className="form-group">
            <label className="text-white">Password</label>
            <input type="password" placeholder="Enter your password" className="form-control"/>
            </div>

            <button className="btn btn-primary">Masuk</button>
            <p className="text-white">dont have an account? <a href="/signup"> Sign Up </a> </p>
            </div>
        </div>  
      </div>
    </div>
  )
}
  
export default App;