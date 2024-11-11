package dto

import (
	"base-gin/domain"
	"base-gin/domain/dao"
	"time"
)

type AuthorCreateReq struct {
	Fullname  string     `json:"fullname" binding:"required,max=56"`
	Gender    *string    `json:"gender" binding:"omitempty,oneof=m f"`
	BirthDate *time.Time `json:"birthdate" binding:"omitempty"`
}

func (o *AuthorCreateReq) ToEntity() dao.Author {
	var item dao.Author
	var gender domain.TypeGender
	if item.Gender == nil {
		gender = "-"
	} else if *item.Gender == domain.GenderFemale {
		gender = "wanita"
	} else {
		gender = "pria"
	}
	item.Fullname = o.Fullname
	item.Gender = &gender
	item.BirthDate = o.BirthDate

	return item
}

type AuthorResp struct {
	ID        int        `json:"id"`
	Fullname  string     `json:"fullname"`
	Gender    *string    `json:"gender,omitempty"`
	BirthDate *time.Time `json:"birth_date,omitempty"`
}

func (o *AuthorResp) FromEntity(item *dao.Author) {
	o.ID = int(item.ID)
	o.Fullname = item.Fullname
}

type AuthorUpdateReq struct {
	ID        uint   `json:"-"`
	Fullname  string `json:"fullname" binding:"required,min=2,max=56"`
	Gender    string `json:"gender" binding:"omitempty,oneof=m f"`
	BirthDate string `json:"birth_date"`
}
