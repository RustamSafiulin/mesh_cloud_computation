package service

import (
	"fmt"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/account_service/internal/dto"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/account_service/internal/model"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/account_service/internal/storage"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/common/errors_helper"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/common/middleware"
	"github.com/pkg/errors"
	"github.com/sarulabs/di"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"time"
)

type AccountService struct {
	accountStorage storage.BaseAccountStorage
}

func PrepareAccountServiceDef(store storage.BaseAccountStorage) di.Def {
	return di.Def{
		Name:  "AccountService",
		Build: func(ctn di.Container) (interface{}, error) {
			return &AccountService{accountStorage: store}, nil
		},
	}
}

func (s *AccountService) Signup(si *dto.SignupInfoDto) (*model.Account, error) {

	existingAcc, err := s.accountStorage.FindByEmail(strings.ToLower(si.Email))
	if existingAcc != nil {
		return nil, errors.WithMessage(errors_helper.ErrAccountAlreadyExists, fmt.Sprintf("Account ID: %s, Reason: %s", existingAcc.ID, err.Error()))
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(si.Password), 8)
	if err != nil {
		return nil, errors.WithMessage(errors_helper.ErrPasswordHashGeneration, fmt.Sprintf("Reason: %s", err.Error()))
	}

	storedAccount, err := s.accountStorage.Insert(&model.Account{
		ID: bson.NewObjectId(),
		Name: si.Username,
		Email: strings.ToLower(si.Email),
		PasswordHash: string(hashedPassword[:]),
		CreatedAt: time.Now().Unix(),
	})

	if err != nil {
		return nil, errors.WithMessage(errors_helper.ErrStorageError, fmt.Sprintf("Reason: %s", err.Error()))
	}

	return storedAccount, nil
}

func (s *AccountService) Signin(si *dto.SigninInfoDto) (*dto.SessionInfoDto, error) {

	existingAccount, err := s.accountStorage.FindByEmail(strings.ToLower(si.Email))
	if err != nil {
		return nil, errors.WithMessage(errors_helper.ErrAccountNotExists, fmt.Sprintf("Account email: %s, Reason: %s", si.Email, err.Error()))
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existingAccount.PasswordHash), []byte(si.Password)); err != nil {
		return nil, errors.WithMessage(errors_helper.ErrWrongPassword, fmt.Sprintf("Reason: %s", err.Error()))
	}

	tokenString, err := middleware.CreateToken(existingAccount.ID.Hex())

	if err != nil {
		return nil, errors.WithMessage(errors_helper.ErrCreateJwtToken, fmt.Sprintf("Reason: %s", err.Error()))
	}

	storeSession := &dto.SessionInfoDto{ AccountID: existingAccount.ID.Hex(), SessionToken: tokenString, UserName: existingAccount.Name, Email: strings.ToLower(existingAccount.Email) }
	return storeSession, nil
}

func (s *AccountService) GetAccountInfo(accountId string) (*model.Account, error) {

	existingAccount, err := s.accountStorage.FindById(accountId)
	if err != nil {
		return existingAccount, errors.WithMessage(errors_helper.ErrAccountNotExists, fmt.Sprintf("Account ID: %s, reason: %s", accountId, err.Error()))
	}

	return existingAccount, nil
}

func (s *AccountService) UpdateAccountInfo() error {
	return nil
}

