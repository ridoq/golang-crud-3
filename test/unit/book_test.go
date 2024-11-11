package unit_test

import (
	"base-gin/domain"
	"base-gin/domain/dao"
	"base-gin/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBook_GetByID_Success(t *testing.T) {
	p := dao.Publisher{
		Name: util.RandomStringAlpha(8),
		City: util.RandomStringAlpha(10),
	}
	_ = publisherRepo.Create(&p)

	gender := domain.GenderFemale
	a := dao.Author{
		Fullname: util.RandomStringAlpha(8),
		Gender: &gender,
	}
	_ = authorRepo.Create(&a)

	b := dao.Book{
		Title:       util.RandomStringAlpha(5) + " " + util.RandomStringAlpha(6),
		PublisherID: p.ID,
		AuthorID: a.ID,
	}
	_ = bookRepo.Create(&b)

	o, err := bookRepo.GetByID(b.ID)
	assert.Nil(t, err)
	assert.NotNil(t, o)
}
