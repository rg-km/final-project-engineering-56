import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import './Navbar.css';
import axios from 'axios';
import Cards from '../Card/Cards';

function Navbar() {
    const [click, setClick] = useState(false);
    const [search,setSearch] = useState("");
    const [bookData,setData] = useState([]);
    const searchBook=(evt)=>{
        if(evt.key==="Enter")
        {
            axios.get('https://www.googleapis.com/books/v1/volumes?q='+search+'&key=AIzaSyC5NOJuFFErZ1zl7lrkYGdCX1Ez_YOjv_4'+'&maxResults=40')
            .then(res=>setData(res.data.items))
            .catch(err=>console.log(err))
        }
    }


    const handleClick = () => setClick(!click);
    const closeMobileMenu = () => setClick(false);

    return (
        <>
            <nav className="navbar">
                <div className="navbar-container">
                    <Link to='/home' className='navbar-logo' onClick={closeMobileMenu}>
                        <img src='/images/logo-small.png' alt='logo' />
                    </Link>
                    <div className="menu-icon" onClick={handleClick}>
                        <i className={click ? "fas fa-times" : "fas fa-bars"} />
                    </div>
                    <ul className={click ? "nav-menu active" : "nav-menu"}>
                        <li className="nav-item">
                            <Link to="/home" className="nav-links" onClick={closeMobileMenu}>Home</Link>
                        </li>
                        <li className="nav-item">
                            <Link to="/favorite" className="nav-links" onClick={closeMobileMenu}>Favorite</Link>
                        </li>
                        <li className="nav-item">
                            <Link to="/popular" className="nav-links" onClick={closeMobileMenu}>Popular</Link>
                        </li>
                        <li className="nav-item">
                            <Link to="/category" className="nav-links" onClick={closeMobileMenu}>Category</Link>
                        </li>
                        <li className="nav-item">
                            <Link to="/about" className="nav-links" onClick={closeMobileMenu}>About</Link>
                        </li>
                        <div className='search-icon'>
                            <input type="search" placeholder="Search" value={search} onChange={e=>setSearch(e.target.value)} onKeyPress={searchBook}/>
                        </div>
                    </ul>
                </div>
            </nav>
<h1>
    Popular
</h1>
            <div className='container'>
                {
                    <Cards book={bookData}/>
}
            </div>
        </>
    )
}

export default Navbar