package main

import (
	"github.com/Go_project/db"
	"github.com/Go_project/server"
)

func main() {
    
	db.Init()
    server.Init()
}
