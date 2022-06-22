import React from "react";
import { Helmet } from "react-helmet";
import '../pages/admin.css'
import logo from '../assets/Logo-admin.png'
import {Container, Form, Button} from "react-bootstrap";
import { Link } from 'react-router-dom'



export default function Tambah() {
    return (
        <div>
        <Helmet>
          <style>{"body { background-color:#010038 ; }"}</style>
        </Helmet>
            <img src={logo} className="float-logo"></img>

        <Container>
            <p className="judul">Edit Buku</p>
            <p className="judul2">Ubah data buku</p>

            <Form>
                <label className="pos">
                    <input type="text" className="form" name="Judul Buku" placeholder="Judul Buku" />
                    <input type="text" className="form" name="kategori" placeholder="Kategori"/>
                    <input type="text" className="form" name="penulis" placeholder="Penulis"/>
                    <input type="text" className="form" name="penerbit" placeholder="Penerbit"/>
                    <input type="text" className="textarea" name="deskripsi" placeholder="Deskripsi"/>
                    <input type="file" accept=".jpg, .png, .jpeg" className="textarea" name="import"/>
                </label>
                <Link to="/admin"><input type="submit" value="SUBMIT" className="setuju"/></Link>
                <Link to="/admin"><input type="submit" value="HAPUS" className="setuju2"/></Link>
                <Link to="/admin"><input type="submit" value="BATAL" className="setuju2"/></Link>
            </Form>
        </Container>
        </div>
    );
}