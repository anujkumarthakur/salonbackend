package payment

import (
	"net/http"
	"userservice/userservice/api/v1/handler"
	"userservice/userservice/api/v1/middleware"

	"github.com/go-chi/chi"
)

// Init initializes all the accounts endpoints
func Init(r chi.Router) {
	r.Route("/payment", func(r chi.Router) {
		r.Use(middleware.AuthRequired)
		r.Method(http.MethodPost, "/pay", handler.Handler(payamount))
	})
}
