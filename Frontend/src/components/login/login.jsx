import profile from "../../Profile-login.png"
import "./login.css"

function App() {
  return (
    <div className="main">
    <h1 className="text-masuk"> Masuk </h1>
      <div className="images"> 
          <img src={profile} alt="" />
      </div>
      <form>
        <div className="email">
        <input className="text-email" type="email" placeholder="your email" />
        </div> 
     <div className="password">
     <input className="text-password" type="password" placeholder="your password" />
      </div>
      <button>Masuk</button>
      <p className="text-white">dont have an account? <a href="signup"> sign up </a> </p>
      </form>
    </div>   
  );
}

export default App;