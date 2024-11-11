package service

import (
	"base-gin/config"
	"base-gin/domain/dto"
	"base-gin/exception"
	"base-gin/repository"
	"base-gin/util"
)

type AccountService struct {
	cfg  *config.Config
	repo *repository.AccountRepository
}

func NewAccountService(
	cfg *config.Config,
	accountRepo *repository.AccountRepository,
) *AccountService {
	return &AccountService{cfg: cfg, repo: accountRepo}
}

func (s *AccountService) Login(p dto.AccountLoginReq) (dto.AccountLoginResp, error) {
	var resp dto.AccountLoginResp

	item, err := s.repo.GetByUsername(p.Username)
	if err != nil {
		return resp, err
	}

	if paswdOk := item.VerifyPassword(p.Password); !paswdOk {
		return resp, exception.ErrUserLoginFailed
	}

	aToken, err := util.CreateAuthAccessToken(*s.cfg, item.Username)
	if err != nil {
		return resp, err
	}

	rToken, err := util.CreateAuthRefreshToken(*s.cfg, item.Username)
	if err != nil {
		return resp, err
	}

	resp.AccessToken = aToken
	resp.RefreshToken = rToken

	return resp, nil
}

func (s *AccountService) Create(params *dto.AccountCreateReq) error {
	newItem := params.ToEntity()
	return s.repo.Create(&newItem)
}

func (s *AccountService) Update(params *dto.AccountUpdateReq) error {
	if params.ID <= 0 {
		return exception.ErrDataNotFound
	}

	return s.repo.Update(params)
}

func (s *AccountService) Delete(id uint) error {
	if id <= 0 {
		return exception.ErrDataNotFound
	}

	return s.repo.Delete(id)
}