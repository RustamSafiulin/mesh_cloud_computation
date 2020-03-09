package handler

import (
	"encoding/json"
	"github.com/RustamSafiulin/3d_reconstruction_service/account_service/internal/model"
	"github.com/RustamSafiulin/3d_reconstruction_service/account_service/internal/service"
	"github.com/RustamSafiulin/3d_reconstruction_service/common/helpers"
	"github.com/gorilla/mux"
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

	var signupInfo model.SignupInfoDto

	err := json.NewDecoder(r.Body).Decode(&signupInfo)
	if err != nil {
		helpers.WriteJSONResponse(w, http.StatusBadRequest, model.ErrorMsgResponse{err.Error()})
		return
	}

	accountService := h.ctn.Get("AccountService").(*service.AccountService)
	err = accountService.Signup(&signupInfo)
	switch err {
	
	}
}

func (h *AccountHandler) GetAccountHandler(w http.ResponseWriter, r *http.Request) {

	var accountId = mux.Vars(r)["account_id"]
	service := h.ctn.Get("AccountService").(*service.AccountService)
	_, err := service.GetAccountInfo(accountId)

	switch err {

	}
}

func (h *AccountHandler) SigninHandler(w http.ResponseWriter, r *http.Request) {

	logrus.Info("SigninHandler")

	var loginInfo model.SigninInfoDto

	err := json.NewDecoder(r.Body).Decode(&loginInfo)
	if err != nil {
		helpers.WriteJSONResponse(w, http.StatusBadRequest, model.ErrorMsgResponse{err.Error()})
		return
	}

	service := h.ctn.Get("AccountService").(*service.AccountService)
	err = service.Signin(&loginInfo)

	switch err {

	}
}

func (h *AccountHandler) UpdateAccountHandler(w http.ResponseWriter, r *http.Request) {

	service := h.ctn.Get("AccountService").(*service.AccountService)
	err := service.UpdateAccountInfo()

	switch err {

	}
}

func (h *AccountHandler) SignoutHandler(w http.ResponseWriter, r *http.Request) {

	service := h.ctn.Get("AccountService").(*service.AccountService)
	err := service.Signout()

	switch err {

	}
}

func (h *AccountHandler) ResetPasswordHandler(w http.ResponseWriter, r *http.Request) {

	//service := h.ioc.Resolve("AccountService").(*service.AccountService)

}

func (h *AccountHandler) ChangePasswordHandler(w http.ResponseWriter, r *http.Request) {

	//service := h.ioc.Resolve("AccountService").(*service.AccountService)
}