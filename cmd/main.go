package main

import (
	"lookfar/internal/utils"
	_ "lookfar/internal/utils"
	"lookfar/pkg/routes"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8848"
	}
	addr := ":" + port
	router := routes.CreateRouter()
	server := http.Server{
		Addr:    addr,
		Handler: router,
	}
	utils.Logger().Info("Server started!", "PORT", port)
	if err := server.ListenAndServe(); err != nil {
		utils.Logger().Error(err.Error())
	}
}
