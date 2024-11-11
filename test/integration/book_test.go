package integration_test

import (
	"base-gin/domain"
	"base-gin/domain/dao"
	"base-gin/domain/dto"
	"base-gin/server"
	"base-gin/util"
	"fmt"
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func createBook() dao.Book {
	p := dao.Publisher{
		Name: util.RandomStringAlpha(8),
		City: util.RandomStringAlpha(10),
	}
	_ = publisherRepo.Create(&p)
	
	gender := domain.GenderFemale
	a := dao.Author{
		Fullname: util.RandomStringAlpha(8),
		Gender:   &gender,
	}

	_ = authorRepo.Create(&a)
	o := dao.Book{
		Title: util.RandomStringAlpha(6),
		PublisherID: p.ID,
		AuthorID:    a.ID,
	}
	_ = bookRepo.Create(&o)

	return o
}

func TestBook_Create_Success(t *testing.T) {

	p := dao.Publisher{
		Name: util.RandomStringAlpha(8),
		City: util.RandomStringAlpha(10),
	}
	_ = publisherRepo.Create(&p)

	gender := domain.GenderFemale
	a := dao.Author{
		Fullname: util.RandomStringAlpha(8),
		Gender:   &gender,
	}
	_ = authorRepo.Create(&a)

	params := dto.BookCreateReq{
		Title:       util.RandomStringAlpha(6),
		PublisherID: p.ID,
		AuthorID:    a.ID,
	}

	w := doTest(
		"POST",
		server.RootBook,
		params,
		createAuthAccessToken(dummyAdmin.Account.Username),
	)
	assert.Equal(t, 201, w.Code)
}

func TestBook_Update_Success(t *testing.T) {
	// requirement
	o := createBook()
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

	params := dto.BookUpdateReq{
		Title:       util.RandomStringAlpha(6),
		PublisherID: p.ID,
		AuthorID:    a.ID,
	}

	w := doTest(
		"PUT",
		fmt.Sprintf("%s/%d", server.RootBook, o.ID),
		params,
		createAuthAccessToken(dummyAdmin.Account.Username),
	)

	// output
	assert.Equal(t, 200, w.Code)

	item, _ := bookRepo.GetByID(o.ID)
	assert.Equal(t, params.Title, item.Title)
	assert.Equal(t, params.PublisherID, item.PublisherID)
	assert.Equal(t, params.AuthorID, item.AuthorID)
	// assert.Equal(t, params.Gender, item.Gender)
	assert.Equal(t, false, item.DeletedAt.Valid)
}

func TestBook_Delete_Success(t *testing.T) {
	o := createBook()
	_ = bookRepo.Create(&o)

	w := doTest(
		"DELETE",
		fmt.Sprintf("%s/%d", server.RootBook, o.ID),
		nil,
		createAuthAccessToken(dummyAdmin.Account.Username),
	)
	assert.Equal(t, 200, w.Code)

	item, _ := bookRepo.GetByID(o.ID)
	assert.Nil(t, item)
}

func TestBook_GetList_Success(t *testing.T) {
	o1 := createBook()
	_ = bookRepo.Create(&o1)

	o2 := createBook()
	_ = bookRepo.Create(&o2)

	w := doTest(
		"GET",
		server.RootBook,
		nil,
		"",
	)
	assert.Equal(t, 200, w.Code)

	body := w.Body.String()
	assert.Contains(t, body, o1.Title)
	assert.Contains(t, body, o2.Title)

	w = doTest(
		"GET",
		server.RootBook+"?q="+o1.Title,
		nil,
		"",
	)
	assert.Equal(t, 200, w.Code)

	body = w.Body.String()
	assert.Contains(t, body, o1.Title)
	assert.NotContains(t, body, o2.Title)
}

func TestBook_GetDetail_Success(t *testing.T) {
	o := createBook()
	_ = bookRepo.Create(&o)

	w := doTest(
		"GET",
		fmt.Sprintf("%s/%d", server.RootBook, o.ID),
		nil,
		"",
	)
	assert.Equal(t, 200, w.Code)

	body := w.Body.String()
	assert.Contains(t, body, o.Title)
}
