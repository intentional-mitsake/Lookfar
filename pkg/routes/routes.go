package routes

import "net/http"

func CreateRouter() http.Handler {
	mux := http.NewServeMux()
	return mux
}
