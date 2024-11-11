package integration_test

import (
	"base-gin/domain/dao"
	"base-gin/domain/dto"
	"base-gin/server"
	"testing"
	"fmt"

	"github.com/stretchr/testify/assert"
)

func TestAccount_Login_Success(t *testing.T) {
	req := dto.AccountLoginReq{
		Username: "admin",
		Password: password,
	}

	w := doTest("POST", server.RootAccount+server.PathLogin, req, "")
	assert.Equal(t, 200, w.Code)
}

func TestAccount_GetProfile_Success(t *testing.T) {
	accessToken := createAuthAccessToken(dummyAdmin.Account.Username)

	w := doTest("GET", server.RootAccount, nil, accessToken)
	assert.Equal(t, 200, w.Code)

	resp := w.Body.String()
	assert.Contains(t, resp, dummyAdmin.Fullname)
}

func TestAccount_GetProfile_ErrorAccessToken(t *testing.T) {
	w := doTest("GET", server.RootAccount, nil, "")
	assert.Equal(t, 401, w.Code)

	w = doTest("GET", server.RootAccount, nil, "accessToken")
	assert.Equal(t, 401, w.Code)
}

func createAccount() dao.Account {
	o := dao.Account{
		Username: "ridoq",
		Password:  "12345678" ,
	}
	_ = accountRepo.Create(&o)
	
	return o
}

func TestAccount_Create_Success(t *testing.T) {
	params := dto.AccountCreateReq{
		Username: "ridoq2",
		Password:  "123456782" ,
	}

	w := doTest(
		"POST",
		server.RootAccount,
		params,
		createAuthAccessToken(dummyAdmin.Account.Username),
	)
	assert.Equal(t, 201, w.Code)
}

func TestAccount_Update_Password_Success(t *testing.T) {
	// requirement
	o := createAccount()

	// action
	params := dto.AccountUpdateReq{
		Password:  "password" ,
	}

	w := doTest(
		"PUT",
		fmt.Sprintf("%s/%d", server.RootAccount, o.ID),
		params,
		createAuthAccessToken(dummyAdmin.Account.Username),
	)

	// output
	assert.Equal(t, 200, w.Code)

	item, _ := accountRepo.GetByUsername(o.Username)
	assert.Equal(t, params.Password, item.Password)
}

func TestAccount_Delete_Success(t *testing.T) {
	o := createAccount()
	_ = accountRepo.Create(&o)

	w := doTest(
		"DELETE",
		fmt.Sprintf("%s/%d", server.RootAccount, o.ID),
		nil,
		createAuthAccessToken(dummyAdmin.Account.Username),
	)
	assert.Equal(t, 200, w.Code)

	item, err := accountRepo.GetByID(o.ID)
	assert.Error(t, err)
	assert.Nil(t, item)
}