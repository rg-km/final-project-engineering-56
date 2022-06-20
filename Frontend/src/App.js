import React from 'react';
import Admin from "./pages/admin";
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
  );
}

export default App;
