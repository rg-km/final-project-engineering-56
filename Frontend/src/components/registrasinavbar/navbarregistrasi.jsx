import React from 'react';
import "./navbarregistrasi.css"

const App = () => {
  return (
    <div style={{marginTop: "100px"}}>
      <div className="container">
        <h1 className="text-daftar">Daftar</h1>
          <div className="row justify-content-center">
            <div className="col-md-3">
              
              <div className="form-group">
              <label className="text-white">Username</label>
              <input type="text" placeholder="Enter your username" className="form-control"/>
              </div>
              
              <div className="form-group">
              <label className="text-white">Email</label>
              <input type="email" placeholder="Enter your email" className="form-control"/>
              </div>

              <div className="form-group">
              <label className="text-white">Password</label>
              <input type="password" placeholder="Enter your password" className="form-control"/>
              </div>

              <div className="form-group">
              <label className="text-white">Confirm Password</label>
              <input type="password" placeholder="Confirm your password" className="form-control"/>
              </div>

              <button className="btn btn-primary">Masuk</button>
              <p className="text-white">have an account? <a href="/"> Login </a> </p>
            </div>
        </div>  
      </div>
    </div>
  )
}
  
export default App;