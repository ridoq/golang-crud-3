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

func createPerson() dao.Person {
    gender := domain.GenderMale
    date, _ := time.Parse("2006-01-02", "1993-09-13")
    o := dao.Person{
        Fullname: util.RandomStringAlpha(6),
        Gender:   &gender,
        BirthDate: &date,
    }
    _ = personRepo.Create(&o)

    return o
}

func TestPerson_Create_Success(t *testing.T) {
	gender := "f"
	date, _ := time.Parse("2006-01-02", "1993-09-13")
	params := dto.PersonCreateReq{
		Fullname:  util.RandomStringAlpha(6),
		Gender:    &gender,
		BirthDate: &date,
	}

	w := doTest(
		"POST",
		server.RootPerson,
		params,
		createAuthAccessToken(dummyAdmin.Account.Username),
	)
	assert.Equal(t, 201, w.Code)
}

func TestPerson_Update_Success(t *testing.T) {
	// requirement
	o := createPerson() 

	// action
	gender := domain.GenderFemale
	params := dto.PersonUpdateReq{
		Fullname:  util.RandomStringAlpha(6),
		Gender:    string(gender),
		BirthDateStr: "1994-09-13",
	}

	w := doTest(
		"PUT",
		fmt.Sprintf("%s/%d", server.RootPerson, o.ID),
		params,
		createAuthAccessToken(dummyAdmin.Account.Username),
	)

	assert.Equal(t, 200, w.Code)

	item, _ := personRepo.GetByID(o.ID)

	assert.Equal(t, params.Fullname, item.Fullname)
	assert.Equal(t, params.Gender, string(*item.Gender)) 
	assert.Equal(t, params.BirthDateStr, item.BirthDate.Format("2006-01-02"))

	assert.False(t, item.DeletedAt.Valid) 
}

func TestPerson_Delete_Success(t *testing.T) {
	o := createPerson()
	_ = personRepo.Create(&o)

	w := doTest(
		"DELETE",
		fmt.Sprintf("%s/%d", server.RootPerson, o.ID),
		nil,
		createAuthAccessToken(dummyAdmin.Account.Username),
	)
	assert.Equal(t, 200, w.Code)
	item, _ := personRepo.GetByID(o.ID)
	assert.Nil(t, item)
}

func TestPerson_GetList_Success(t *testing.T) {
	o := createPerson()
	_ = personRepo.Create(&o)

	w := doTest(
		"GET",
		server.RootPerson,
		nil,
		"",
	)
	assert.Equal(t, 200, w.Code)

	// Check the response body for the person's details
	body := w.Body.String()
	// Use string comparison for gender and birthdate
	assert.Contains(t, body, o.Fullname)
	assert.Contains(t, body, string(*o.Gender)) // Compare with string ("pria" or "wanita")
	assert.Contains(t, body, o.BirthDate.Format("2006-01-02"))
}


func TestPerson_GetDetail_Success(t *testing.T) {
	o := createPerson()
	_ = personRepo.Create(&o)

	w := doTest(
		"GET",
		fmt.Sprintf("%s/%d", server.RootPerson, o.ID),
		nil,
		"",
	)
	assert.Equal(t, 200, w.Code)

	// Check the response body for the person's details
	body := w.Body.String()
	// Use string comparison for gender and birthdate
	assert.Contains(t, body, o.Fullname)
	assert.Contains(t, body, string(*o.Gender)) // Compare with string ("pria" or "wanita")
	assert.Contains(t, body, o.BirthDate.Format("2006-01-02"))
}

