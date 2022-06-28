package main

import (
	// "github.com/Go_project/db"
	"net/http"

	"github.com/Go_project/server"
)

func main() {

	// db.Init()
	server.Init()
	http.ListenAndServe(":8080", nil)
}
