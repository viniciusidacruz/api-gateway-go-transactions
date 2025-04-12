package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/viniciusidacruz/api-gateway-go-transactions/internal/modules/account/services"
	"github.com/viniciusidacruz/api-gateway-go-transactions/internal/web/handlers"
)

type Server struct {
	router *chi.Mux
	server *http.Server
	port string

	accountService *services.AccountService
	
}

func NewServer(port string, accountService *services.AccountService) *Server {
	return &Server{
		port: port,
		accountService: accountService,
		router: chi.NewRouter(),
	}
}

func (s *Server) ConfigureRoutes() {
	accountHandler := handlers.NewAccountHandler(s.accountService)	

	s.router.Post("/accounts", accountHandler.Create)
	s.router.Get("/accounts", accountHandler.Get)
}

func (s *Server) Start() error {
	s.server = &http.Server{
		Addr:    ":" + s.port,
		Handler: s.router,
	}

	return s.server.ListenAndServe()
}

