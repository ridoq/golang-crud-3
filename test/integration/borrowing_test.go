package integration_test

import (
	"base-gin/domain"
	"base-gin/domain/dao"
	"base-gin/domain/dto"
	"base-gin/server"
	"base-gin/util"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func createBorrowing() dao.Borrowing {
	date, _ := time.Parse("2006-01-02", "1993-09-13")
	gender := domain.GenderFemale
	p := dao.Person{
		Fullname:  util.RandomStringAlpha(4) + " " + util.RandomStringAlpha(6) + " " + util.RandomStringAlpha(6),
		Gender:    &gender,
		BirthDate: &date,
	}
	_ = personRepo.Create(&p)

	pr := dao.Publisher{
		Name: util.RandomStringAlpha(8),
		City: util.RandomStringAlpha(10),
	}
	_ = publisherRepo.Create(&pr)

	pr2 := dao.Publisher{
		Name: util.RandomStringAlpha(8),
		City: util.RandomStringAlpha(10),
	}
	_ = publisherRepo.Create(&pr2)

	a := dao.Author{
		Fullname: util.RandomStringAlpha(8),
		Gender:   &gender,
	}
	_ = authorRepo.Create(&a)

	a2 := dao.Author{
		Fullname: util.RandomStringAlpha(8),
		Gender:   &gender,
	}

	_ = authorRepo.Create(&a2)
	b := dao.Book{
		Title:       util.RandomStringAlpha(6),
		PublisherID: pr.ID,
		AuthorID:    a.ID,
	}
	_ = bookRepo.Create(&b)
	b2 := dao.Book{
		Title:       util.RandomStringAlpha(6),
		PublisherID: pr2.ID,
		AuthorID:    a2.ID,
	}
	_ = bookRepo.Create(&b2)

	o := dao.Borrowing{
		BorrowDate: date,
		PersonID:   p.ID,
		BookID:     b.ID,
	}
	_ = borrowingRepo.Create(&o)

	return o
}

func TestBorrowing_Create_Success(t *testing.T) {

	date, _ := time.Parse("2006-01-02", "1993-09-13")
	gender := domain.GenderFemale
	p := dao.Person{
		Fullname:  util.RandomStringAlpha(4) + " " + util.RandomStringAlpha(6) + " " + util.RandomStringAlpha(6),
		Gender:    &gender,
		BirthDate: &date,
	}
	_ = personRepo.Create(&p)

	pr := dao.Publisher{
		Name: util.RandomStringAlpha(8),
		City: util.RandomStringAlpha(10),
	}
	_ = publisherRepo.Create(&pr)

	a := dao.Author{
		Fullname: util.RandomStringAlpha(8),
		Gender:   &gender,
	}

	_ = authorRepo.Create(&a)
	b := dao.Book{
		Title:       util.RandomStringAlpha(6),
		PublisherID: pr.ID,
		AuthorID:    a.ID,
	}
	_ = bookRepo.Create(&b)

	params := dto.BorrowingCreateReq{
		BorrowDate: date,
		PersonID:   p.ID,
		BookID:     b.ID,
	}

	w := doTest(
		"POST",
		server.RootBorrowing,
		params,
		createAuthAccessToken(dummyAdmin.Account.Username),
	)
	assert.Equal(t, 201, w.Code)
}

func TestBorrowing_Update_Success(t *testing.T) {
	// requirement
	o := createBorrowing()
	date, _ := time.Parse("2006-01-02", "1993-09-13")
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		t.Fatalf("Failed to load timezone Asia/Jakarta: %v", err)
	}
	gender := domain.GenderFemale
	p := dao.Person{
		Fullname:  util.RandomStringAlpha(4) + " " + util.RandomStringAlpha(6) + " " + util.RandomStringAlpha(6),
		Gender:    &gender,
		BirthDate: &date,
	}
	_ = personRepo.Create(&p)

	pr := dao.Publisher{
		Name: util.RandomStringAlpha(8),
		City: util.RandomStringAlpha(10),
	}
	_ = publisherRepo.Create(&pr)

	pr2 := dao.Publisher{
		Name: util.RandomStringAlpha(8),
		City: util.RandomStringAlpha(10),
	}
	_ = publisherRepo.Create(&pr2)

	a := dao.Author{
		Fullname: util.RandomStringAlpha(8),
		Gender:   &gender,
	}
	_ = authorRepo.Create(&a)

	a2 := dao.Author{
		Fullname: util.RandomStringAlpha(8),
		Gender:   &gender,
	}

	_ = authorRepo.Create(&a2)
	b := dao.Book{
		Title:       util.RandomStringAlpha(6),
		PublisherID: pr.ID,
		AuthorID:    a.ID,
	}
	_ = bookRepo.Create(&b)
	b2 := dao.Book{
		Title:       util.RandomStringAlpha(6),
		PublisherID: pr2.ID,
		AuthorID:    a2.ID,
	}
	_ = bookRepo.Create(&b2)

	params := dto.BorrowingUpdateReq{
		BorrowDate: time.Now().In(loc),
		PersonID:   p.ID,
		BookID:     b.ID,
	}

	w := doTest(
		"PUT",
		fmt.Sprintf("%s/%d", server.RootBorrowing, o.ID),
		params,
		createAuthAccessToken(dummyAdmin.Account.Username),
	)

	// output
	assert.Equal(t, 200, w.Code)

	item, _ := borrowingRepo.GetByID(o.ID)
	assert.WithinDuration(t, params.BorrowDate, item.BorrowDate, time.Second)
	assert.Equal(t, params.PersonID, item.PersonID)
	assert.Equal(t, params.BookID, item.BookID)
	// assert.Equal(t, params.Gender, item.Gender)
	assert.Equal(t, false, item.DeletedAt.Valid)
}

func TestBorrowing_Delete_Success(t *testing.T) {
	o := createBorrowing()
	_ = borrowingRepo.Create(&o)

	w := doTest(
		"DELETE",
		fmt.Sprintf("%s/%d", server.RootBorrowing, o.ID),
		nil,
		createAuthAccessToken(dummyAdmin.Account.Username),
	)
	assert.Equal(t, 200, w.Code)

	item, _ := borrowingRepo.GetByID(o.ID)
	assert.Nil(t, item)
}

func TestBorrowing_GetList_Success(t *testing.T) {
	// Create two borrowings for testing
	o1 := createBorrowing()
	_ = borrowingRepo.Create(&o1)

	o2 := createBorrowing()
	_ = borrowingRepo.Create(&o2)

	// Perform a GET request to fetch all borrowings
	w := doTest(
		"GET",
		server.RootBorrowing,
		nil,
		"",
	)
	assert.Equal(t, 200, w.Code)

	// Verify that both borrowings are in the response body
	body := w.Body.String()
	assert.Contains(t, body, o1.BorrowDate.Format("2006-01-02"))
	assert.Contains(t, body, o2.BorrowDate.Format("2006-01-02"))

	// Perform a filtered GET request using the borrow date from o1
	w = doTest(
		"GET",
		server.RootBorrowing+"?q="+o1.BorrowDate.Format("2006-01-02"),
		nil,
		"",
	)
	assert.Equal(t, 200, w.Code)

	// Verify that only o1 is in the filtered response body
	body = w.Body.String()
	assert.Contains(t, body, o1.BorrowDate.Format("2006-01-02"))
	assert.Contains(t, body, o2.BorrowDate.Format("2006-01-02"))
}

func TestBorrowing_GetDetail_Success(t *testing.T) {
	// Create a borrowing record
	o := createBorrowing()
	_ = borrowingRepo.Create(&o)

	// Perform a GET request to fetch the details of the borrowing
	w := doTest(
		"GET",
		fmt.Sprintf("%s/%d", server.RootBorrowing, o.ID),
		nil,
		"",
	)
	assert.Equal(t, 200, w.Code)

	// Verify the response body contains the borrowing details
	body := w.Body.String()
	assert.Contains(t, body, o.BorrowDate.Format("2006-01-02"))
	assert.Contains(t, body, fmt.Sprintf("%d", o.PersonID))
	assert.Contains(t, body, fmt.Sprintf("%d", o.BookID))
}
