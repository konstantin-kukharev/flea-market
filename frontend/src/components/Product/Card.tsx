// import { classNames, openLink } from '@telegram-apps/sdk-react';
import React, { ReactNode, useState, useEffect } from 'react';
// import { Link as RouterLink, type LinkProps } from 'react-router-dom';

import './Card.css';

export interface Product {
    id: number;
    name: string;
    price: string;
    description: string;
  }

export interface ProductPhoto {
    [key: number]: string[];
  }

export const Card: React.FC<{ item: Product, images: string[] }> = ({ item, images }) => {
    const [isModalOpen, setIsModalOpen] = useState(false);
    const [currentIndex, setCurrentIndex] = useState(0);

    const handlePrev = () => {
        setCurrentIndex((prev) => (prev > 0 ? prev - 1 : images.length - 1));
      };
    
      const handleNext = () => {
        setCurrentIndex((prev) => (prev < images.length - 1 ? prev + 1 : 0));
      };
    
      useEffect(() => {
        const handleKeyDown = (e: KeyboardEvent) => {
          if (isModalOpen) {
            if (e.key === 'ArrowLeft') handlePrev();
            if (e.key === 'ArrowRight') handleNext();
          }
        };

        window.addEventListener('keydown', handleKeyDown);
        return () => window.removeEventListener('keydown', handleKeyDown);
      }, [isModalOpen, currentIndex]);

    return (
        <div className="product-card">
            {images ? (
            <div className="product-gallery">
            {images.map((img, index) => (
                <img
                    key={img}
                    src={img}
                    alt={`Product ${index + 1}`}
                    onClick={() => {
                    setCurrentIndex(index);
                    setIsModalOpen(true);
                    }}
                    className="product-thumb"
                />
            ))} 
            <Modal isOpen={isModalOpen} onClose={() => setIsModalOpen(false)}>
                <div className="modal-slider">
                    <button className="nav-button prev" onClick={handlePrev}>&#10094;</button>
                    
                    <div className="slide-container">
                    <img 
                        src={images[currentIndex]} 
                        alt={`Product view ${currentIndex + 1}`} 
                        className="modal-image"
                    />
                    <div className="image-counter">
                        {currentIndex + 1} / {images.length}
                    </div>
                    </div>

                    <button className="nav-button next" onClick={handleNext}>&#10095;</button>
                </div>
            </Modal>
            </div> ) : (<></>)}
            <div className="product-details">
                <h4><a href="#">{item.name}</a></h4>
                <p>{item.description}</p>
                <div className="product-bottom-details">
                    <div className="product-price">{item.price}</div>
                    <div className="product-links">
                        <a href=""><i className="fa fa-heart"></i></a>
                        <a href=""><i className="fa fa-shopping-cart"></i></a>
                    </div>
                </div>
            </div>
        </div>
    )
}

interface ModalProps {
    isOpen: boolean;
    onClose: () => void;
    children: ReactNode;
  }
  
  const Modal: React.FC<ModalProps> = ({ isOpen, onClose, children }) => {
    if (!isOpen) return null;
  
    return (
      <div className="modal-overlay" onClick={onClose}>
        <div 
          className="modal-content" 
          onClick={(e: React.MouseEvent) => e.stopPropagation()}
          role="dialog"
          aria-modal="true"
        >
          {children}
          <button 
            className="modal-close" 
            onClick={onClose}
            aria-label="Close modal"
          >
            &times;
          </button>
        </div>
      </div>
    );
  };