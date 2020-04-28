package server

import (
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/account_service/internal/handler"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/common/helpers"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/common/middleware"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Server struct {
	r  *mux.Router
	ah *handler.AccountHandler
	th *handler.TaskHandler
}

func NewServer(ah *handler.AccountHandler, th *handler.TaskHandler) *Server {
	return &Server{
		r: mux.NewRouter(),
		ah: ah,
		th: th,
	}
}

func (s *Server) Start() {

	err := http.ListenAndServe(":8081", helpers.EnableCors(s.r))
	if err != nil {
		logrus.WithError(err).Fatal("Error during start Http server on port 8081")
	}
}

func (s *Server) SetupRoutes() {

	s.r.Use(middleware.PanicRecoveryMiddleware)

	api := s.r.PathPrefix("/api/v1/").Subrouter()

	api.HandleFunc("/accounts", s.ah.CreateAccountHandler).Methods("POST")
	api.HandleFunc("/accounts/{account_id}", middleware.JwtTokenValidation(s.ah.GetAccountHandler)).Methods("GET")
	api.HandleFunc("/accounts/signin", s.ah.SigninHandler).Methods("POST")
	api.HandleFunc("/accounts/{account_id}", middleware.JwtTokenValidation(s.ah.UpdateAccountHandler)).Methods("PUT")
	api.HandleFunc("/accounts/reset_password", middleware.JwtTokenValidation(s.ah.ResetPasswordHandler)).Methods("POST")
	api.HandleFunc("/accounts/change_password", middleware.JwtTokenValidation(s.ah.ChangePasswordHandler)).Methods("POST")

	api.HandleFunc("/tasks", middleware.JwtTokenValidation(s.th.CreateTaskHandler)).Methods("POST")
	api.HandleFunc("/tasks", middleware.JwtTokenValidation(s.th.GetAllAccountTasksHandler)).Methods("GET")
	api.HandleFunc("/tasks/{task_id}/upload", middleware.JwtTokenValidation(s.th.UploadTaskDataHandler)).Methods("POST")
	api.HandleFunc("/tasks/{task_id}/download", middleware.JwtTokenValidation(s.th.DownloadTaskDataHandler)).Methods("GET")
	api.HandleFunc("/tasks/{task_id}/start", middleware.JwtTokenValidation(s.th.StartTaskHandler)).Methods("POST")
	api.HandleFunc("/tasks/{task_id}", middleware.JwtTokenValidation(s.th.GetTaskHandler)).Methods("GET")
	api.HandleFunc("/tasks/{task_id}", middleware.JwtTokenValidation(s.th.DeleteTaskHandler)).Methods("DELETE")

	s.r.PathPrefix("/").Handler(handler.IndexHandler("./public/dist"))
}
