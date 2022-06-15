import "./navbarregistrasi.css"

function App() {
  return (
  <div className="main">
    <h1 className="text-daftar">Daftar</h1>
    
    <form>
      <div>
        <div className="username">
        <label html='username' className='label'>
        Username :
        </label>
        <input id='username' className="text-username" type="username" placeholder="Enter your username"/>
        </div>
        
        <div className="email">
        <label html='email' className='label'>
        Email :
        </label>
        <input id='email' className="text-email" type="email" placeholder="Enter your email" />
        </div> 

        <div className="password">
        <label html='password' className='label'>
        Password :
        </label>
        <input className="text-password" type="password" placeholder="Enter your password" />
        </div>

        <div className="confirmpassword">
        <label html='confirmpassword' className='label'>
        Confirm Password :
        </label>
        <input className="text-confirmpassword" type="confirmpassword" placeholder="Enter confirm password" />
        </div>
      </div>

      <button>Masuk</button>
      <p className="text-white">have an account? <a href="/"> Login </a> </p>
      </form>
  </div>   
  );
}

export default App;