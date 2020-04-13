package handler

import (
	"encoding/json"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/account_service/internal/dto"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/account_service/internal/service"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/common/errors_helper"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/common/helpers"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/sarulabs/di"
	"github.com/sirupsen/logrus"
	"net/http"
)

// AccountHandler handle account routes
type AccountHandler struct {
	ctn di.Container
}

func NewAccountHandler(ctn di.Container) *AccountHandler {
	return &AccountHandler{ctn: ctn}
}

func (h *AccountHandler) CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	logrus.Info("CreateAccountHandler")

	var signupInfo dto.SignupInfoDto

	err := json.NewDecoder(r.Body).Decode(&signupInfo)
	if err != nil {
		helpers.WriteJSONResponse(w, http.StatusBadRequest, dto.ErrorMsgResponse{err.Error()})
		return
	}

	accountService := h.ctn.Get("AccountService").(*service.AccountService)
	createdAccount, err := accountService.Signup(&signupInfo)

	if err != nil {

		logrus.Debugf("Error was caused. Reason: %s", err.Error())

		if errors.Cause(err) == errors_helper.ErrAccountAlreadyExists {
			helpers.WriteJSONResponse(w, http.StatusConflict, dto.ErrorMsgResponse{err.Error()})
			return
		}

		if errors.Cause(err) == errors_helper.ErrPasswordHashGeneration {
			helpers.WriteJSONResponse(w, http.StatusInternalServerError, dto.ErrorMsgResponse{err.Error()})
			return
		}

		if errors.Cause(err) == errors_helper.ErrStorageError {
			helpers.WriteJSONResponse(w, http.StatusInternalServerError, dto.ErrorMsgResponse{err.Error()})
			return
		}

		helpers.WriteJSONResponse(w, http.StatusInternalServerError, dto.ErrorMsgResponse{err.Error()})

	} else {

		createdAccountDto := dto.AccountDtoFromAccount(createdAccount)
		helpers.WriteJSONResponse(w, http.StatusOK, createdAccountDto)
	}
}

func (h *AccountHandler) GetAccountHandler(w http.ResponseWriter, r *http.Request) {

	var accountId = mux.Vars(r)["account_id"]
	service := h.ctn.Get("AccountService").(*service.AccountService)
	account, err := service.GetAccountInfo(accountId)

	if err != nil {

		if errors.Cause(err) == errors_helper.ErrAccountNotExists {
			helpers.WriteJSONResponse(w, http.StatusNotFound, dto.ErrorMsgResponse{err.Error()})
			return
		}

		helpers.WriteJSONResponse(w, http.StatusInternalServerError, dto.ErrorMsgResponse{err.Error()})

	} else {
		accountDto := dto.AccountDtoFromAccount(account)
		helpers.WriteJSONResponse(w, http.StatusOK, accountDto)
	}
}

func (h *AccountHandler) SigninHandler(w http.ResponseWriter, r *http.Request) {

	logrus.Info("SigninHandler")

	var loginInfo dto.SigninInfoDto

	err := json.NewDecoder(r.Body).Decode(&loginInfo)
	if err != nil {
		helpers.WriteJSONResponse(w, http.StatusBadRequest, dto.ErrorMsgResponse{err.Error()})
		return
	}

	service := h.ctn.Get("AccountService").(*service.AccountService)
	sessionInfo, err := service.Signin(&loginInfo)

	if err != nil {

		if errors.Cause(err) == errors_helper.ErrWrongPassword {
			helpers.WriteJSONResponse(w, http.StatusUnauthorized, dto.ErrorMsgResponse{err.Error()})
			return
		}

		helpers.WriteJSONResponse(w, http.StatusInternalServerError, dto.ErrorMsgResponse{err.Error()})

	} else {

		helpers.WriteJSONResponse(w, http.StatusOK, sessionInfo)
	}
}

func (h *AccountHandler) UpdateAccountHandler(w http.ResponseWriter, r *http.Request) {

	service := h.ctn.Get("AccountService").(*service.AccountService)
	err := service.UpdateAccountInfo()

	switch err {

	}
}

func (h *AccountHandler) ResetPasswordHandler(w http.ResponseWriter, r *http.Request) {

	//service := h.ioc.Resolve("AccountService").(*service.AccountService)

}

func (h *AccountHandler) ChangePasswordHandler(w http.ResponseWriter, r *http.Request) {

	//service := h.ioc.Resolve("AccountService").(*service.AccountService)
}