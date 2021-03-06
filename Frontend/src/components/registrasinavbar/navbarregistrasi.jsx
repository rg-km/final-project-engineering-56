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
  .then(res => {
    console.log(res.data)
  })
  .catch(err => {
    console.log(err)
  }); 
  }


  return (
  <div className="container-fluid">
        <h1 className="text-daftar text-center">Daftar</h1>
          <div className="row justify-content-center">
            <div className="col-md-3">

              <div className="form-group">
              <label className="text-white py-2">Username</label>
              <input type="text" placeholder="Enter your username" className="form-control"
               onChange = {(e) => setUsername(e.target.value)} value={Username}
              />
																					  
				
              </div>

              <div className="form-group">
              <label className="text-white py-2">Email</label>
              <input type="email" placeholder="Enter your email" className="form-control"
               onChange = {(e) => setEmail(e.target.value)} value={Email}
              />
																				
              </div>

              <div className="form-group">
              <label className="text-white py-2">Password</label>
              <input type="password" placeholder="Enter your password" className="form-control"
               onChange = {(e) => setPassword(e.target.value)} value={Password}
              />
              </div>

              <button className="btn btn-primary px-3 py-2 fs-5 mt-2 mx-auto" onClick={onRegister()}>Create Akun</button>
              <p className="text-white text-center ">Have an account? <a href="/login"> Login </a> </p>
            </div>
        </div>  
      </div>
		  
 
  )
}
  
export default App;