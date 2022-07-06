package v1

import (
	"userservice/userservice/api/v1/account"
	"userservice/userservice/api/v1/bookservice"
	"userservice/userservice/api/v1/payment"

	"github.com/go-chi/chi"
)

// Routes - all the registered routes
func Routes(router chi.Router) {
	router.Route("/v1", Init)

}

// Init initializes all the v1 routes
func Init(r chi.Router) {
	r.Group(initUnAuthorizedRoutes)
	r.Group(initAuthorizedRoutes)
}

// initUnAuthorizedRoutes initializes all the routes that are public
func initUnAuthorizedRoutes(r chi.Router) {
	// Accounts related endpoints
	r.Route("/accounts", account.Init)
}

func initAuthorizedRoutes(r chi.Router) {
	r.Route("/auth", bookservice.Init)
	r.Route("/authpay", payment.Init)
}
