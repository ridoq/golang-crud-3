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

func createAuthor() dao.Author {
	gender := domain.GenderMale
	o := dao.Author{
		Fullname: util.RandomStringAlpha(6),
		Gender:   &gender,
	}
	_ = authorRepo.Create(&o)

	return o
}

func TestAuthor_Create_Success(t *testing.T) {
	gender := "m"
	params := dto.AuthorCreateReq{
		Fullname: util.RandomStringAlpha(6),
		Gender:   &gender,
	}

	w := doTest(
		"POST",
		server.RootAuthor,
		params,
		createAuthAccessToken(dummyAdmin.Account.Username),
	)
	assert.Equal(t, 201, w.Code)
}

func TestAuthor_Update_Success(t *testing.T) {
	// requirement
	o := createAuthor()

	// action
	gender := domain.GenderFemale
	params := dto.AuthorUpdateReq{
		Fullname: util.RandomStringAlpha(6),
		Gender:   string(gender),
	}

	w := doTest(
		"PUT",
		fmt.Sprintf("%s/%d", server.RootAuthor, o.ID),
		params,
		createAuthAccessToken(dummyAdmin.Account.Username),
	)

	// output
	assert.Equal(t, 200, w.Code)

	item, _ := authorRepo.GetByID(o.ID)
	assert.Equal(t, params.Fullname, item.Fullname)
	// assert.Equal(t, params.Gender, item.Gender)
	assert.Equal(t, false, item.DeletedAt.Valid)
}

func TestAuthor_Delete_Success(t *testing.T) {
	o := createAuthor()
	_ = authorRepo.Create(&o)

	w := doTest(
		"DELETE",
		fmt.Sprintf("%s/%d", server.RootAuthor, o.ID),
		nil,
		createAuthAccessToken(dummyAdmin.Account.Username),
	)
	assert.Equal(t, 200, w.Code)

	item, _ := authorRepo.GetByID(o.ID)
	assert.Nil(t, item)
}

func TestAuthor_GetList_Success(t *testing.T) {
	o1 := createAuthor()
	_ = authorRepo.Create(&o1)

	o2 := createAuthor()
	_ = authorRepo.Create(&o2)

	w := doTest(
		"GET",
		server.RootAuthor,
		nil,
		"",
	)
	assert.Equal(t, 200, w.Code)

	body := w.Body.String()
	assert.Contains(t, body, o1.Fullname)
	assert.Contains(t, body, o2.Fullname)

	w = doTest(
		"GET",
		server.RootAuthor+"?q="+o1.Fullname,
		nil,
		"",
	)
	assert.Equal(t, 200, w.Code)

	body = w.Body.String()
	assert.Contains(t, body, o1.Fullname)
	assert.NotContains(t, body, o2.Fullname)
}

func TestAuthor_GetDetail_Success(t *testing.T) {
	o := createAuthor()
	_ = authorRepo.Create(&o)

	w := doTest(
		"GET",
		fmt.Sprintf("%s/%d", server.RootAuthor, o.ID),
		nil,
		"",
	)
	assert.Equal(t, 200, w.Code)

	body := w.Body.String()
	assert.Contains(t, body, o.Fullname)
}
