package main

import (
	"ToDo/Router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var r = Router.InitRouter()
	fmt.Println("******************")
	fmt.Println("Create router complete")
	fmt.Printf("******************\n\n")
	fmt.Println("Starting server...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
