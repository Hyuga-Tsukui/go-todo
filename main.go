package main

import (
	"fmt"
	"log"
	"net/http"
	"todo/internal/handler"
	"todo/internal/middleware"
)

func main() {
	r := middleware.NewRouter()
	r.RegistrationHandler(
		middleware.NewHandler("/", handler.Index),
	)

	fmt.Println("server start up...")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
