import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Home from './pages/Home';
import './App.css';
import BookShow from './pages/BookShow';
// import Header from './components/Header/Header';

function App() {
  return (
    <>
      <Router>
        <Routes>
        <Route exect path="/" element={<Home />} />
        <Route exect path="/book-show" element={<BookShow />} />
        </Routes>
      </Router>
    </>
  );
}

export default App;
