<<<<<<< HEAD
import React from 'react';
import Tambah from "./pages/tambah";
import Edit from "./pages/edit";
import { Route, Routes } from "react-router-dom";

function App() {
  return (
    <main aria-label="App" className="App">
      <div className="routes" aria-label="routes">
        <Routes>
          <Route path="/admin" element={<Admin />} />
          <Route path="/tambah" element={<Tambah />} />
          <Route path="/edit" element={<Edit />} />
        </Routes>
      </div>
    </main>
=======
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
>>>>>>> origin/FixAll2
  );
} 

export default App;