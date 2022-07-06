package account

import (
	"net/http"
	"userservice/userservice/api/v1/handler"

	"github.com/go-chi/chi"
)

// Init initializes all the accounts endpoints
func Init(r chi.Router) {

	//r.Method(http.MethodPost, "/activate-user", api.Handler(activateUser))
	r.Method(http.MethodPost, "/register", handler.Handler(registerUser))
	r.Method(http.MethodPost, "/send/otp", handler.Handler(sendOtp))
	r.Method(http.MethodPost, "/verify/otp", handler.Handler(verifyOtp))
	r.Method(http.MethodPost, "/password/reset", handler.Handler(resetPassword))
	r.Method(http.MethodPost, "/auth/login", handler.Handler(authEmailLogin))

}
