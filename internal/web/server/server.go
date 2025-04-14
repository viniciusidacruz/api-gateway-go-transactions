package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/viniciusidacruz/api-gateway-go-transactions/internal/modules/account/services"
	"github.com/viniciusidacruz/api-gateway-go-transactions/internal/modules/invoice/service"
	"github.com/viniciusidacruz/api-gateway-go-transactions/internal/web/handlers"
	"github.com/viniciusidacruz/api-gateway-go-transactions/internal/web/middleware"
)

type Server struct {
	router *chi.Mux
	server *http.Server
	port string

	accountService *services.AccountService
	invoiceService *service.InvoiceService
}

func NewServer(port string, accountService *services.AccountService, invoiceService *service.InvoiceService) *Server {
	return &Server{
		port: port,
		accountService: accountService,
		invoiceService: invoiceService,
		router: chi.NewRouter(),
	}
}

func (s *Server) ConfigureRoutes() {
	accountHandler := handlers.NewAccountHandler(s.accountService)	
	invoiceHandler := handlers.NewInvoiceHandler(s.invoiceService)
	authMiddleware := middleware.NewAuthMiddleware(s.accountService)

	s.router.Post("/accounts", accountHandler.Create)
	s.router.Get("/accounts", accountHandler.Get)

	s.router.Group(func(r chi.Router) {
		r.Use(authMiddleware.Authenticate)
		r.Post("/invoice", invoiceHandler.Create)
		r.Get("/invoice/{id}", invoiceHandler.GetByID)
		r.Get("/invoice", invoiceHandler.GetAllByAccountID)
	})
}

func (s *Server) Start() error {
	s.server = &http.Server{
		Addr:    ":" + s.port,
		Handler: s.router,
	}

	return s.server.ListenAndServe()
}

