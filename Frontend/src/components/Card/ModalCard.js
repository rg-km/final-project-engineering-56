import axios, { Axios } from 'axios';
import React, { useState } from 'react';
import './ModalCard.css';




const ModalCard = ({ show, item, onClose}) => {
    if (!show) {
        return null;
    }
    let thumbnail = item.volumeInfo.imageLinks && item.volumeInfo.imageLinks.smallThumbnail;

    function AddFavorit(){
        axios.post('http://localhost:8080/api/buku/setfavorit',{
           id_user : id_user,
           id_books : id_books
        })
        .then(res => {
            console.log(res.data)
            
        })
        .catch((err) => {
            console.log(err)
        })
    }
    return (
        <>
            <div className='overlay'>
                <div className='overlay-inner'>
                    <button className='close' onClick={onClose}><i className='fas fa-times' /> </button>
                    <div className='inner-box'>
                        <img src={thumbnail} alt='' />
                        <div className='info'>
                        <li key="{item.ID}">{item.ID}</li>
                            <h1>{item.volumeInfo.title}</h1>
                            <h3>{item.volumeInfo.authors}</h3>
                            <h4>{item.volumeInfo.publisher}<br></br><span>{item.volumeInfo.publishedDate}</span></h4>
                            <a href={item.volumeInfo.previewLink}><button>Info</button></a>
                            <a  onClick={AddFavorit()} ><button>Add Favorit</button></a>
                        </div>
                    </div>
                    <h4  className='description'>{item.volumeInfo.description}</h4>
                </div>
            </div>
        </>
    )
}

export default ModalCard;