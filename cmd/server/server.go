package main

import (
	"Comments/pkg/api"
	"Comments/pkg/db"
	"log"
	"net/http"
)

type server struct {
	db  db.Storage
	api *api.API
}

func main() {
	var srv server
	db1, err := db.New("postgres://postgres:password@192.168.1.191:5432/Comments")
	if err != nil {
		log.Fatal(err)
	}
	srv.db = db1
	srv.api = api.New(srv.db)
	err = http.ListenAndServe(":8080", srv.api.Router())
	if err != nil {
		log.Fatal(err)
	}
}
