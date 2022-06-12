import "./register.css"

function App() {
  return (
    <div className="main">
    <h1 className="text-daftar">Daftar</h1>
      <div className="images"> 
          <img src={profile} alt="" />
      </div>
      <form>
        <div className="username">
        <input className="text-username" type="username" placeholder="username"/>
        </div>
        
        <div className="email">
        <input className="text-email" type="email" placeholder="your email" />
        </div> 
        
        <div className="password">
        <input className="text-password" type="password" placeholder="your password" />
        </div>

        <div className="confirmpassword">
        <input className="text-confirmpassword" type="confirmpassword" placeholder="confirmpassword" />
        </div>

      <button>Sign in</button>
      </form>
    </div>   
  );
}

export default App;