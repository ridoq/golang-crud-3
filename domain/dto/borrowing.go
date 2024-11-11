package dto

import (
	"base-gin/domain/dao"
	"time"
)

type BorrowingCreateReq struct {
	BorrowDate      time.Time  `json:"borrow_date" binding:"required"`
	ReturnDate      *time.Time `json:"return_date" binding:"omitempty"`
	PersonID        uint       `json:"person_id" binding:"required"`
	BorrowingPerson dao.Person `json:"borrowing_person" binding:"omitempty"`
	BookID          uint       `json:"book_id" binding:"required"`
	BorrowingBook   dao.Book   `json:"borrowing_book" binding:"omitempty"`
}

func (o *BorrowingCreateReq) ToEntity() dao.Borrowing {
	var item dao.Borrowing
	item.BorrowDate = o.BorrowDate
	item.ReturnDate = o.ReturnDate
	item.PersonID = o.PersonID
	item.BookID = o.BookID
	item.BorrowingPerson = o.BorrowingPerson
	item.BorrowingBook = o.BorrowingBook

	return item
}

type BorrowingResp struct {
	ID              int        `json:"id"`
	BorrowDate      time.Time  `json:"borrow_date" `
	ReturnDate      *time.Time `json:"return_date" `
	PersonID        uint       `json:"person_id" `
	BorrowingPerson dao.Person `json:"borrowing_person" `
	BookID          uint       `json:"book_id" `
	BorrowingBook   dao.Book   `json:"borrowing_book" `
}

func (o *BorrowingResp) FromEntity(item *dao.Borrowing) {
	o.ID = int(item.ID)
	o.BorrowDate = item.BorrowDate
	o.PersonID = item.PersonID
	o.BookID = item.BookID
}

type BorrowingUpdateReq struct {
	ID              uint       `json:"-"`
	BorrowDate      time.Time  `json:"borrow_date" binding:"required"`
	ReturnDate      *time.Time `json:"return_date" binding:"omitempty"`
	PersonID        uint       `json:"person_id" binding:"required"`
	BorrowingPerson dao.Person `json:"borrowing_person" binding:"omitempty"`
	BookID          uint       `json:"book_id" binding:"required"`
	BorrowingBook   dao.Book   `json:"borrowing_book" binding:"omitempty"`
}
