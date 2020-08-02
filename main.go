package main

import (
	"go-api/server/http"
	"log"
)

func main()  {
	srv := http.Server()
	log.Println("Server listening on ", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
