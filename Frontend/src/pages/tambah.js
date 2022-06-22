import React, { useState } from "react";
import { Helmet } from "react-helmet";
import '../pages/admin.css'
import logo from '../assets/Logo-admin.png'
import {Container, Form, Button} from "react-bootstrap";
import { Link } from 'react-router-dom'


export default function Tambah(props) {
    const [judul, setJudul] = useState('');
    const [penulis, setPenulis] = useState('');
    const [penerbit, setPenerbit] = useState('');

    const changeJudul = event => {
        setJudul(event.target.value);
    };

    const changePenulis = event => {
        setPenulis(event.target.value);
    };

    const changePenerbit = event => {
        setPenerbit(event.target.value);
    };

    const transferValue = event => {
        event.preventDefault();
        const val = {
        judul,
        penulis,
        penerbit
        };
        props.func(val);
        clearState();
    };

    const clearState = () => {
        setJudul('');
        setPenulis('');
        setPenerbit('');
    };
    return (
        <div>
        <Helmet>
          <style>{"body { background-color:#010038 ; }"}</style>
        </Helmet>
            <img src={logo} className="float-logo"></img>

        <Container>
            <p className="judul">Tambah Buku</p>
            <p className="judul2">Masukkan data buku</p>

            <Form>
                <label className="pos">
                    <input type="text" value={judul} onChange={changeJudul} className="form" name="Judul Buku" placeholder="Judul Buku" />
                    <input type="text" className="form" name="kategori" placeholder="Kategori"/>
                    <input type="text" value={penulis} onChange={changePenulis} className="form" name="penulis" placeholder="Penulis"/>
                    <input type="text" value={penerbit} onChange={changePenerbit}className="form" name="penerbit" placeholder="Penerbit"/>
                    <input type="text" className="textarea" name="deskripsi" placeholder="Deskripsi"/>
                    <input type="file" accept=".jpg, .png, .jpeg" className="textarea" name="import"/>
                </label>
                <Link to="/admin"><input type="submit" value="SUBMIT" className="setuju"/></Link>
            </Form>
        </Container>
        </div>
    );
}