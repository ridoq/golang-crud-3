package rest

import (
	"base-gin/domain/dto"
	"base-gin/exception"
	"base-gin/server"
	"base-gin/service"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	hr            *server.Handler
	service       *service.AccountService
	personService *service.PersonService
}

func NewAccountHandler(
	hr *server.Handler,
	accountService *service.AccountService,
	personService *service.PersonService,
) *AccountHandler {
	return &AccountHandler{
		hr: hr, service: accountService, personService: personService}
}

func (h *AccountHandler) Route(app *gin.Engine) {
	grp := app.Group(server.RootAccount)
	grp.POST("", h.hr.AuthAccess(), h.create)
	grp.POST(server.PathLogin, h.login)
	grp.GET("", h.hr.AuthAccess(), h.getProfile)
	grp.PUT("/:id", h.hr.AuthAccess(), h.update)
	grp.DELETE("/:id", h.hr.AuthAccess(), h.delete)
}

// login godoc
//
//	@Summary Account login
//	@Description Account login using username & password combination.
//	@Accept json
//	@Produce json
//	@Param cred body dto.AccountLoginReq true "Credential"
//	@Success 200 {object} dto.SuccessResponse[dto.AccountLoginResp]
//	@Failure 400 {object} dto.ErrorResponse
//	@Failure 422 {object} dto.ErrorResponse
//	@Failure 500 {object} dto.ErrorResponse
//	@Router /accounts/login [post]
func (h *AccountHandler) login(c *gin.Context) {
	var req dto.AccountLoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(h.hr.BindingError(err))
		return
	}

	data, err := h.service.Login(req)
	if err != nil {
		switch {
		case errors.Is(err, exception.ErrUserNotFound),
			errors.Is(err, exception.ErrUserLoginFailed):
			c.JSON(http.StatusBadRequest, h.hr.ErrorResponse(exception.ErrUserLoginFailed.Error()))
		default:
			h.hr.ErrorInternalServer(c, err)
		}
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse[dto.AccountLoginResp]{
		Success: true,
		Message: "Login berhasil",
		Data:    data,
	})
}

// getProfile godoc
//
//	@Summary Get account's profile
//	@Description Get profile of logged-in account.
//	@Produce json
//	@Security BearerAuth
//	@Success 200 {object} dto.SuccessResponse[dto.AccountProfileResp]
//	@Failure 401 {object} dto.ErrorResponse
//	@Failure 403 {object} dto.ErrorResponse
//	@Failure 404 {object} dto.ErrorResponse
//	@Failure 500 {object} dto.ErrorResponse
//	@Router /accounts [get]
func (h *AccountHandler) getProfile(c *gin.Context) {
	accountID, _ := c.Get(server.ParamTokenUserID)

	data, err := h.personService.GetAccountProfile((accountID).(uint))
	if err != nil {
		switch {
		case errors.Is(err, exception.ErrUserNotFound):
			c.JSON(http.StatusNotFound, h.hr.ErrorResponse(err.Error()))
		default:
			h.hr.ErrorInternalServer(c, err)
		}

		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse[dto.AccountProfileResp]{
		Success: true,
		Message: "Profile pengguna",
		Data:    data,
	})
}

// create godoc
//
//	@Summary Create a account
//	@Description Create a account.
//	@Accept json
//	@Produce json
//	@Security BearerAuth
//	@Param detail body dto.accountUpdateReq true "account's detail"
//	@Success 201 {object} dto.SuccessResponse[any]
//	@Failure 401 {object} dto.ErrorResponse
//	@Failure 403 {object} dto.ErrorResponse
//	@Failure 422 {object} dto.ErrorResponse
//	@Failure 500 {object} dto.ErrorResponse
//	@Router /accounts [post]
func (h *AccountHandler) create(c *gin.Context) {
	var req dto.AccountCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.hr.BindingError(err)
		return
	}

	err := h.service.Create(&req)
	if err != nil {
		h.hr.ErrorInternalServer(c, err)
		return
	}

	c.JSON(http.StatusCreated, dto.SuccessResponse[any]{
		Success: true,
		Message: "Data berhasil disimpan",
	})
}

// update godoc
//
//	@Summary Update a account's detail
//	@Description Update a account's detail.
//	@Accept json
//	@Produce json
//	@Security BearerAuth
//	@Param id path int true "account's ID"
//	@Param detail body dto.accountUpdateReq true "account's detail"
//	@Success 200 {object} dto.SuccessResponse[any]
//	@Failure 400 {object} dto.ErrorResponse
//	@Failure 401 {object} dto.ErrorResponse
//	@Failure 403 {object} dto.ErrorResponse
//	@Failure 404 {object} dto.ErrorResponse
//	@Failure 422 {object} dto.ErrorResponse
//	@Failure 500 {object} dto.ErrorResponse
//	@Router /accounts/{id} [put]
func (h *AccountHandler) update(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, h.hr.ErrorResponse("ID tidak valid"))
		return
	}

	var req dto.AccountUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(h.hr.BindingError(err))
		return
	}
	req.ID = uint(id)

	err = h.service.Update(&req)
	if err != nil {
		switch {
		case errors.Is(err, exception.ErrDataNotFound):
			c.JSON(http.StatusNotFound, h.hr.ErrorResponse(err.Error()))
		default:
			h.hr.ErrorInternalServer(c, err)
		}

		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse[any]{
		Success: true,
		Message: "Data berhasil disimpan",
	})
}

// delete godoc
//
//	@Summary Delete a account
//	@Description Delete a account.
//	@Produce json
//	@Security BearerAuth
//	@Param id path int true "account's ID"
//	@Success 200 {object} dto.SuccessResponse[any]
//	@Failure 400 {object} dto.ErrorResponse
//	@Failure 401 {object} dto.ErrorResponse
//	@Failure 403 {object} dto.ErrorResponse
//	@Failure 404 {object} dto.ErrorResponse
//	@Failure 500 {object} dto.ErrorResponse
//	@Router /accounts/{id} [delete]
func (h *AccountHandler) delete(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, h.hr.ErrorResponse("ID tidak valid"))
		return
	}

	err = h.service.Delete(uint(id))
	if err != nil {
		h.hr.ErrorInternalServer(c, err)
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse[any]{
		Success: true,
		Message: "Data berhasil dihapus",
	})
}
