import './App.css';
import profile from "./image/login.jpg";

function App() {
  return (
    <div className="main">
      <div className="sub-main">
        <div>
          <div className="imgs">
            <div className="container-image">
              <img src={profile} alt= "profile" className="profile"/>

              </div>

              </div>
              </div>
                <h1>MASUK</h1>
            <div>
              <input type="text" placeholder="your email"className="name"/>
          </div>
              <div className="second.inpput">
              <input type="text" placeholder="your password"className="name"/>
        </div>
          
      </div>
    </div>  
  );
}

export default App;