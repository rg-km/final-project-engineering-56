import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Home from './pages/Home';
import './App.css';
import NavbarLogin from './components/LoginNavbar/NavbarLogin';
import NavbarRegister from './components/RegisterNavbar/NavbarRegister';
import Cards from './components/Card/Cards';
import '../node_modules/bootstrap/dist/css/bootstrap.min.css';
import Heroes from './pages/Heroes';


function App() {
  return (
    <>
      <Router>
        <Routes>
        <Route exect path="/" element={<Heroes />} />
        <Route path="/home" element={<Home />} />
        <Route path="/login" element={<NavbarLogin />} />
        <Route path="/register" element={<NavbarRegister />} />
        <Route path="/cards" element={<Cards />} />
        </Routes>
      </Router>
    </>
  );
}

export default App;
