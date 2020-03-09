package service

import (
	"errors"
	"github.com/RustamSafiulin/3d_reconstruction_service/account_service/internal/model"
	"github.com/RustamSafiulin/3d_reconstruction_service/account_service/internal/storage"
	"github.com/RustamSafiulin/3d_reconstruction_service/common/middleware"
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

func (s *AccountService) Signup(si *model.SignupInfoDto) error {

	logrus.Info("Signup")

	existingAcc, err := s.accountStorage.FindByEmail(si.Email)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(si.Password), 8)
	if err != nil {
		return nil
	}

	storedAccount := model.Account{Name: si.Username, Email: si.Email, PasswordHash: string(hashedPassword[:])}
	err = s.accountStorage.Insert(&storedAccount)

	if err != nil {
		return errors.New("")
	}

	return nil
}

func (s *AccountService) Signin(si *model.SigninInfoDto) (*model.SessionInfoDto, error) {
	logrus.Info("Signin")

	existingAccount, err := s.accountStorage.FindByEmail(si.Email)

	if err := bcrypt.CompareHashAndPassword([]byte(existingAccount.PasswordHash), []byte(si.Password)); err != nil {
		return nil, ErrWrongPassword
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
		return nil, ErrCreateJwtToken
	}

	storeSession := &model.SessionInfoDto{ AccountID: existingAccount.ID.String(), SessionToken: tokenString}

	return storeSession, nil
}

func (s *AccountService) Signout() error {
	logrus.Info("Signout")
	return nil
}

func (s *AccountService) GetAccountInfo(accountId string) (*model.Account, error) {
	logrus.Info("GetAccountInfo")

	existingAccount, err := s.accountStorage.FindById(accountId)


	return existingAccount, nil
}

func (s *AccountService) UpdateAccountInfo() error {
	logrus.Info("UpdateAccountInfo")
	return nil
}

