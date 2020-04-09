package service

import (
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/account_service/internal/dto"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/account_service/internal/model"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/account_service/internal/storage"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/common/errors_helper"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/common/middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/sarulabs/di"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
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

	existingAcc, err := s.accountStorage.FindByEmail(si.Email)
	if existingAcc != nil {
		return nil, errors_helper.NewApplicationError(errors_helper.ErrAccountAlreadyExists, existingAcc.ID)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(si.Password), 8)
	if err != nil {
		return nil, errors_helper.NewApplicationError(errors_helper.ErrPasswordHashGeneration)
	}

	storedAccount := model.Account{
		ID: bson.NewObjectId(),
		Name: si.Username,
		Email: si.Email,
		PasswordHash: string(hashedPassword[:]),
		CreatedAt: time.Now().Unix(),
	}
	err = s.accountStorage.Insert(&storedAccount)

	if err != nil {
		logrus.Debug(err.Error())
		return nil, errors_helper.NewApplicationError(errors_helper.ErrStorageError, err.Error())
	}

	return &storedAccount, nil
}

func (s *AccountService) Signin(si *dto.SigninInfoDto) (*dto.SessionInfoDto, error) {

	existingAccount, err := s.accountStorage.FindByEmail(si.Email)
	if err != nil {
		return nil, errors_helper.NewApplicationError(errors_helper.ErrAccountNotExists)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existingAccount.PasswordHash), []byte(si.Password)); err != nil {
		return nil, errors_helper.NewApplicationError(errors_helper.ErrWrongPassword)
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &middleware.JwtClaims{
		Username: existingAccount.ID.String(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(middleware.JwtKey)

	if err != nil {
		return nil, errors_helper.NewApplicationError(errors_helper.ErrCreateJwtToken)
	}

	storeSession := &dto.SessionInfoDto{ AccountID: existingAccount.ID.Hex(), SessionToken: tokenString}
	return storeSession, nil
}

func (s *AccountService) GetAccountInfo(accountId string) (*model.Account, error) {

	existingAccount, err := s.accountStorage.FindById(accountId)
	if err != nil {
		return nil, errors_helper.NewApplicationError(errors_helper.ErrAccountNotExists)
	}

	return existingAccount, nil
}

func (s *AccountService) UpdateAccountInfo() error {
	return nil
}

