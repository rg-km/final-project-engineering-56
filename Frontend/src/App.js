import React from 'react';
import Tambah from "./pages/tambah";
import Edit from "./pages/edit";
import Login from "./pages/login";
import Registrasi from "./pages/registrasi";
import Admin from "./pages/admin";
import Home from "./pages/Home";
// import beranda from "./pages/beranda";
import { Route, Routes } from "react-router-dom";

function App() {
  return (
    <main aria-label="App" className="App">
      <div className="routes" aria-label="routes">
        <Routes>
          <Route path="/" element={<Login />} />
          <Route path="/admin" element={<Admin />} />
          <Route path="/tambah" element={<Tambah />} />
          <Route path="/edit" element={<Edit />} />
          <Route path="/signup" element={<Registrasi />} />
          <Route path="/Home" element={<Home />} />
        </Routes>
      </div>
    </main>
  );
} 

export default App;