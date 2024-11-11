package dao

import (
	"time"

	"gorm.io/gorm"
)

type Borrowing struct {
	gorm.Model
	BorrowDate      time.Time  `gorm:"size:56;not null;"`
	ReturnDate      *time.Time `gorm:"size:56;"`
	PersonID        uint       `gorm:"not null;"`
	BorrowingPerson Person     `gorm:"foreignKey:PersonID;"`
	BookID          uint       `gorm:"not null;"`
	BorrowingBook   Book       `gorm:"foreignKey:BookID;"`
}

func (Borrowing) tableName() string {
	return "borrowing"
}
