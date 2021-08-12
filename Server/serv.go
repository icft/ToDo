package Server

import (
	"ToDo/Structures"
	"encoding/json"
	"log"
	"net/http"
)

const HOST = "mongodb://localhost:27017"
const dbName = "ToDo"
const collectionName = "rec"

func CreateProblem(w http.ResponseWriter, r *http.Request) {
	var jsonDecoder = json.NewDecoder(r.Body)
	var rec Structures.Records
	err := jsonDecoder.Decode(&rec)
	if err != nil {
		log.Fatal(err)
	}

}
