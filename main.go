package main

import (
	"ToDo/Router"
	"fmt"
	"net/http"
)

func main() {
	var r = Router.InitRouter()
	var _ = http.ListenAndServe(":8080", r)
	/*if err != nil {
		log.Fatal(err)
	}*/
	fmt.Println("Starting server...")
}
