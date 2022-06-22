import React, { useState } from "react";
import { Helmet } from "react-helmet";
import '../pages/admin.css'
import logo from '../assets/Logo-admin.png'
import { Link } from 'react-router-dom'


export default function Admin(props) {

  

    return (
        <div>
        <Helmet>
          <style>{"body { background-color:#010038 ; }"}</style>
        </Helmet>
            <img src={logo} className="float-logo"></img>
        <container>
        <p className="judul">Selamat Datang Admin</p>
        <table>
          <thead>
        <tr>
          <th>Judul</th>
          <th>Penulis</th>
          <th>Tanggal Upload</th>
          <th></th>
        </tr>
          </thead>
          {/* <tbody>{tableRows}</tbody> */}
        <tr>
          <td>Lorem ipsum</td>
          <td>Lorem ipsum</td>
          <td>Lorem ipsum</td>
          <td><Link to="/edit"><button>Cek</button></Link></td>
        </tr>
        <tr>
          <td>Lorem ipsum</td>
          <td>Lorem ipsum</td>
          <td>Lorem ipsum</td>
          <td><Link to="/edit"><button>Cek</button></Link></td>
        </tr>
        <tr>
          <td>Lorem ipsum</td>
          <td>Lorem ipsum</td>
          <td>Lorem ipsum</td>
          <td><Link to="/edit"><button>Cek</button></Link></td>
        </tr>
      </table>
      </container>
      <container>
      <Link to="/tambah"><button className="tambah">Tambah Buku</button></Link>
      </container>
      </div>
    );
}