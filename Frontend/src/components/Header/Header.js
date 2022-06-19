import React from 'react';
import { Link } from 'react-router-dom';
import './Header.css';

function Header() {
    return (
        <>
            <nav className='head-navbar'>
                <div className='head-nav-container'>
                    <Link to='/' className='head-nav-logo'>
                        <img src='/images/logo.png' alt='logo' />
                    </Link>
                    <ul className='head-nav-menu'>
                        <li className='head-nav-item'>
                            <Link to='/' className='btn btn-rounded'>Sign In</Link>
                        </li>
                        <li className='head-nav-item'>
                            <Link to='/' className='btn btn-rounded'>Sign Up</Link>
                        </li>
                    </ul>
                </div>
            </nav>
            <div className='header-container'>
                <div className='header-wrapper'>
                    <header>
                        <div className='hero-container'>
                            <h1>Bukuku</h1>
                            <p>Baca buku dimana saja dan kapan saja
                            </p>
                            <button className='btn-b btn-rounded'>Baca Sekarang</button>
                        </div>
                        <div className='hero-image'>
                            <img src='/images/cover.png' alt='cover' className='img-cover' />
                        </div>
                    </header> 
                </div>
            </div>
        </>
    )
}

export default Header;