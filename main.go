package main

import (
	"ToDo/Router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var router = Router.InitRouter()
	fmt.Println("******************")
	fmt.Println("Create router complete")
	fmt.Printf("******************\n\n")
	fmt.Println("Starting server...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
