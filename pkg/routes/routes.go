package routes

import (
	"net/http"
)

func CreateRouter() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/", nil) // need to redirect to a middleware that checks if the user is logged in
	mux.Handle("POST /auth/register", nil)
	mux.Handle("POST /auth/login", nil)
	mux.Handle("POST /auth/logout", nil)
	mux.Handle("POST /auth/refresh", nil)
	return mux
}
