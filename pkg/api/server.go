package api

import (
	"RESTAPIService2/pkg/api/handler"
	"RESTAPIService2/pkg/api/middlewares"
	"RESTAPIService2/pkg/service/auth"
	"RESTAPIService2/pkg/service/item"
	"RESTAPIService2/pkg/service/list"
	"context"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
	router     *mux.Router
	subRouter  *mux.Router
}

func NewServer(middleware *middlewares.UserIdentityMiddleware) *Server {
	router := mux.NewRouter()
	api := router.PathPrefix("/api").Subrouter()
	api.Use(middleware.UserIdentity)
	return &Server{
		httpServer: &http.Server{
			Addr:           ":8000",
			MaxHeaderBytes: 1 << 20, // 1 MB
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			Handler:        router,
		},
		router:    router,
		subRouter: api,
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func (s *Server) HandleAuth(service auth.AuthorizationService) {
	s.router.HandleFunc("/auth/sign-up/", handler.SignUp(service)).Methods(http.MethodPost)
	s.router.HandleFunc("/auth/sign-in/", handler.SignIn(service)).Methods(http.MethodPost)
}

func (s *Server) HandleLists(service list.TodoListService) {
	s.subRouter.HandleFunc("/lists/", handler.CreateList(service)).Methods(http.MethodPost)
	s.subRouter.HandleFunc("/lists/", handler.GetAllLists(service)).Methods(http.MethodGet)
	s.subRouter.HandleFunc("/lists/{id}", handler.GetListById(service)).Methods(http.MethodGet)
	s.subRouter.HandleFunc("/lists/{id}", handler.DeleteList(service)).Methods(http.MethodDelete)
	s.subRouter.HandleFunc("/lists/{id}", handler.UpdateList(service)).Methods(http.MethodPut)
}

func (s *Server) HandleItems(service item.TodoItemService) {
	s.subRouter.HandleFunc("/lists/{id}/items/", handler.CreateItem(service)).Methods(http.MethodPost)
	s.subRouter.HandleFunc("/lists/{id}/items/", handler.GetAllItems(service)).Methods(http.MethodGet)
	s.subRouter.HandleFunc("/items/{id}", handler.GetItemById(service)).Methods(http.MethodGet)
	s.subRouter.HandleFunc("/items/{id}", handler.DeleteItem(service)).Methods(http.MethodDelete)
	s.subRouter.HandleFunc("/items/{id}", handler.UpdateItem(service)).Methods(http.MethodPut)
}
