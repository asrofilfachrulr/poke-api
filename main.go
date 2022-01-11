package main

import (
	"fmt"
	"log"
	"net/http"
	"poke_api/app"
)

func main() {
	r := app.NewRouter()
	fmt.Println("Server starting at port 8080..")
	log.Fatal(http.ListenAndServe(":8080", r))
}
