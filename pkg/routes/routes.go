package routes

import (
	"lookfar/pkg/middleware"
	"net/http"
)

func CreateRouter() http.Handler {
	mux := http.NewServeMux()
	// HOME routes
	mux.Handle("/", middleware.Redirect(http.HandlerFunc(handleHome))) // need to redirect to a middleware that checks if the user is logged in
	mux.Handle("/home", middleware.Redirect(http.HandlerFunc(handleHome)))
	// AUTH routes
	mux.Handle("POST /auth/register", nil)
	mux.Handle("POST /auth/login", nil)
	mux.Handle("POST /auth/logout", nil)
	mux.Handle("POST /auth/refresh", nil)
	return mux
}

func handleHome(w http.ResponseWriter, r *http.Request) {}
