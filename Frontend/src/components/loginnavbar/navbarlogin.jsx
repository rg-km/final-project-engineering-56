import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import profile from "../../Profile-login.png";
import "./navbarlogin.css";

const App = () => {
  const navigation=useNavigate()
  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")
function onLogin(){
  if(email === "admin@email.com" && password === "123"){
    navigation("/beranda")
  }
}
  return (
    <div style={{marginTop: "100px"}}>
      <div className="container">
        <h1 className="text-masuk">Masuk {JSON.stringify(email)} </h1>
        
            <div className="images"> 
            <img src={profile} alt="" />
            </div>
            <div className="row justify-content-center">
            <div className="col-md-3">
              
            <div className="form-group">
            <label className="text-white">Email</label>
            <input onChange={(e) => setEmail(e.target.value)} value={email} type="email" placeholder="Enter your email" className="form-control"/>
            </div>

            <div className="form-group">
            <label className="text-white">Password</label>
            <input onChange={(e) => setPassword(e.target.value)} value={password} type="password" placeholder="Enter your password" className="form-control"/>
            </div>

            <button className="btn btn-primary" onClick={onLogin}>Masuk</button>
            <p className="text-white">dont have an account? <a href="/signup"> Sign Up </a> </p>
            </div>
        </div>  
      </div>
    </div>
  )
}
  
export default App;