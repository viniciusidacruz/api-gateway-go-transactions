package middleware

import (
	"net/http"

	"github.com/viniciusidacruz/api-gateway-go-transactions/internal/modules/account/domain"
	"github.com/viniciusidacruz/api-gateway-go-transactions/internal/modules/account/services"
)

type AuthMiddleware struct {
	accountService *services.AccountService
}

func NewAuthMiddleware(accountService *services.AccountService) *AuthMiddleware {
	return &AuthMiddleware{accountService: accountService}
}

func (m *AuthMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-Key")

		if apiKey == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		_, err := m.accountService.FindByAPIKey(apiKey)

		if err != nil {
			if err == domain.ErrAccountNotFound {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		next.ServeHTTP(w, r)
	})
}

