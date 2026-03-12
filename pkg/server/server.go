package server

import (
	"net/http"
	"os"

	"go_final_project/pkg/api"
)

func Run() error {
	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = "7540"
	}

	api.Init()

	webDir := "./web"
	http.Handle("/", http.FileServer(http.Dir(webDir)))

	return http.ListenAndServe(":"+port, nil)
}
