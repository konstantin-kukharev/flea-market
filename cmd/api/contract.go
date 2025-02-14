package main

import "time"

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

type Product struct {
	ID          int       `json:"id"`
	CategoryID  int       `json:"category_id"`
	State       string    `json:"state"`
	Name        string    `json:"name"`
	Price       string    `json:"price"`
	Description string    `json:"description,omitempty"`
	Author      string    `json:"author_id"`
	PublishDate time.Time `json:"publish_date"`
}

type Photo map[int][]string

type Comment struct {
	ID          int       `json:"id"`
	ProductID   int       `json:"product_id"`
	AuthorID    string    `json:"author_id"`
	RecipientID string    `json:"recipient_id"`
	Description string    `json:"description,omitempty"`
	PublishDate time.Time `json:"publish_date"`
}

type Rating struct {
	ProductID   int       `json:"product_id"`
	AuthorID    string    `json:"author_id"`
	Rate        bool      `json:"rate"`
	PublishDate time.Time `json:"publish_date"`
}

type ProductRequest struct {
	Limit int `json:"limit"`
	// Offset - ProductID
	Offset   int    `json:"offset"`
	AuthorID string `json:"author_id,omitempty"`
	State    string `json:"state,omitempty"`
}

type ProductResponse struct {
	Products     []Product   `json:"products"`
	Photo        []Photo     `json:"images,omitempty"`
	CommentCount map[int]int `json:"comment_count,omitempty"`
	Rating       map[int]int `json:"rating,omitempty"`
	Paginator    `json:"paginator"`
}

type Paginator struct {
	Limit int `json:"limit"`
	// Offset = ProductID
	Offset int `json:"offset"`
	Total  int `json:"total"`
	Page   int `json:"page"`
}

type ProductInterface interface {
	Publish(Product) error
	UnPublish(id int) error
	GetProducts(ProductRequest) (ProductResponse, error)
	GetProduct(id int) (Product, error)
	GetProductComments(id int) ([]Comment, error)
	GetProductRating(id int) (int, error)
}

type PhotoInterface interface {
	Publish(Photo) error
	UnPublish(id int) error
	GetPhotos(...int) ([]Photo, error)
}

type RatingInterface interface {
	Publish(Rating) error
	GetRate(...int) (map[int]int, error)
}

type CommentInterface interface {
	Publish(Comment) error
	Get(id int) (Comment, error)
	GetCommentCount(...int) (map[int]int, error)
}
