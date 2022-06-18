// import profile from "../../Profile-login.png"
import "./App.css";
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Login from "./page/login";
import Registrasi from "./page/registrasi";

function App() {
  return (
    <>
      <center>
      <Router>
        <Routes>
          <Route path="/" element={<Login />} />
          <Route path="/signup" element={<Registrasi />} />
        </Routes>
      </Router>
      </center>
    </>
  );
} 

export default App;