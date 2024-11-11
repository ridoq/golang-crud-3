package dto

import (
	"base-gin/domain/dao"
)

type BookCreateReq struct {
	Title         string        `json:"title" binding:"required,max=255"`
	Subtitle      *string       `json:"subtitle" binding:"omitempty,max=255"`
	PublisherID   uint          `json:"publisher_id" binding:"required"`
	BookPublisher dao.Publisher `json:"book_publisher" binding:"omitempty"`
	AuthorID      uint          `json:"author_id" binding:"required"`
	BookAuthor    dao.Author    `json:"book_author" binding:"omitempty"`
}

func (o *BookCreateReq) ToEntity() dao.Book {
	var item dao.Book
	item.Title = o.Title
	item.Subtitle = o.Subtitle
	item.PublisherID = o.PublisherID
	item.AuthorID = o.AuthorID
	item.BookPublisher = o.BookPublisher
	item.BookAuthor = o.BookAuthor

	return item
}

type BookResp struct {
	ID            int           `json:"id"`
	Title         string        `json:"title"`
	Subtitle      *string       `json:"subtitle"`
	PublisherID   uint          `json:"publisher_id"`
	BookPublisher dao.Publisher `json:"book_publisher"`
	AuthorID      uint          `json:"author_id"`
	BookAuthor    dao.Author    `json:"book_author"`
}

func (o *BookResp) FromEntity(item *dao.Book) {
	o.ID = int(item.ID)
	o.Title = item.Title
	o.PublisherID = item.PublisherID
	o.AuthorID = item.AuthorID
}

type BookUpdateReq struct {
	ID            uint          `json:"-"`
	Title         string        `json:"title" binding:"required,max=255"`
	Subtitle      *string       `json:"subtitle" binding:"omitempty,max=255"`
	PublisherID   uint          `json:"publisher_id" binding:"required"`
	BookPublisher dao.Publisher `json:"book_publisher" binding:"omitempty"`
	AuthorID      uint          `json:"author_id" binding:"required"`
	BookAuthor    dao.Author    `json:"book_author" binding:"omitempty"`
}
