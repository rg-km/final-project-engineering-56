import React from 'react'
import './Cards.css'

const Cards = ({ book }) => {
  console.log(book)
  return (
    <>
      {
        book.map((item) => {
          let thumbnail = item.volumeInfo.imageLinks && item.volumeInfo.imageLinks.smallThumbnail;
          if (thumbnail!== undefined) {
            return (
              <>
                <div className='card'>
                  <img src={thumbnail} alt='' />
                  <div className='bottom'>
                    <h3 className='title'>{item.volumeInfo.title}</h3>
                    <p className='description'>This book is very good for student colage</p>
                  </div>
                </div>
              </>
            )
          }


        })
      }
    </>
  )
}

export default Cards;



