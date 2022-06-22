import React,{ useState} from 'react';

import "./navbarregistrasi.css"
import axios from 'axios';

const App = () => {
  const [Username, setUsername] = useState("")
  const [Email, setEmail] = useState("")
  const [Password, setPassword] = useState("")
 function onRegister(){
  axios.post('http://localhost:8080/api/user/register', {
    username: Username,
    email: Email,
    password: Password
  })
  // .then(res => {
  //   console.log(res.data)
  //   if(res.data.success){
  //     alert("Register Success")
  //   }
  // })
  // .catch(err => {
  //   console.log(err)
  // }); 
  }


  return (
  <form>
    <div style={{marginTop: "100px"}}>
      <div className="container">
        <h1 className="text-daftar">Daftar</h1>
          <div className="row justify-content-center">
            <div className="col-md-3">
           
              <div className="form-group">
              <label className="text-white">Username</label>
              <input type="text" placeholder="Enter your username" className="form-control"
                      onChange = {(e) => setUsername(e.target.value)} value={Username}
              />
              </div>
              
              <div className="form-group">
              <label className="text-white">Email</label>
              <input type="email" placeholder="Enter your email" className="form-control"
                      onChange = {(e) => setEmail(e.target.value)} value={Email}
              />
              </div>

              <div className="form-group">
              <label className="text-white">Password</label>
              <input type="password" placeholder="Enter your password" className="form-control"
                      onChange = {(e) => setPassword(e.target.value)} value={Password}
              />
              </div>

              {/* <div className="form-group">
              <label className="text-white">Confirm Password</label>
              <input type="password" placeholder="Confirm your password" className="form-control"/>
              </div> */}

              <button className="btn btn-primary" onSubmit={onRegister()} > Create Akun</button>
              <p className="text-white">have an account? <a href="/"> Login </a> </p>
            </div>
        </div>  
      </div>
    </div>
  </form>
  )
}
  
export default App;