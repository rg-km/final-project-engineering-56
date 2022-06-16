import React from 'react'; 
import profile from "../../Profile-login.png"
import "./navbarlogin.css"

function App() {
  return (
    <div className="main">
      <h1 className="text-masuk"> Masuk </h1>
        <div className="images"> 
          <img src={profile} alt="" />
        </div>
      
      <form >
        <div>
          <div className="email">
            <label html='email' className='label'>
            Email :  
            </label>
            <input id='email' className="text-email" type="text" placeholder="Enter your email" />
          </div> 
        
          <div className="password">
            <label html='password' className='label'>
            Password :
            </label>
            <input id='password'className="text-password" type="password" placeholder="Enter your password" />
          </div>
        </div>
      
        <button>Masuk</button>
        <p className="text-white">dont have an account? <a href="signup"> Sign Up </a> </p>
      </form>
    </div>   
  );
}

export default App;