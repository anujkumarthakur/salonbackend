package bookservice

import (
	"net/http"
	"userservice/userservice/api/v1/handler"
	"userservice/userservice/api/v1/middleware"

	"github.com/go-chi/chi"
)

// Init initializes all the accounts endpoints
func Init(r chi.Router) {
	r.Method(http.MethodGet, "/shops", handler.Handler(getShops))
	r.Method(http.MethodPost, "/shopfilter", handler.Handler(filterShopList))
	r.Route("/booked", func(r chi.Router) {
		r.Use(middleware.AuthRequired)
		r.Method(http.MethodPost, "/service", handler.Handler(serviceBook))
	})
}
