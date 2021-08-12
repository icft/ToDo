package Server

import (
	"ToDo/Database"
	"ToDo/Structures"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func CreateProblem(w http.ResponseWriter, r *http.Request) {
	var jsonDecoder = json.NewDecoder(r.Body)
	var rec Structures.Records
	if err := jsonDecoder.Decode(&rec); err != nil {
		log.Fatal(err)
	}
	Database.AddToDb(rec)
	if err := json.NewEncoder(w).Encode(rec); err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func SetExecution(w http.ResponseWriter, r *http.Request) {
	var vars = mux.Vars(r)
	Database.UpdateRec(vars["id"], 1)
	if err := json.NewEncoder(w).Encode(vars["id"]); err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func CancelExecution(w http.ResponseWriter, r *http.Request) {
	var vars = mux.Vars(r)
	Database.UpdateRec(vars["id"], 0)
	if err := json.NewEncoder(w).Encode(vars["id"]); err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func DeleteProblem(w http.ResponseWriter, r *http.Request) {
	var vars = mux.Vars(r)
	Database.DelInDb(vars["id"])
	if err := json.NewEncoder(w).Encode("id"); err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	list := Database.GetAll()
	if err := json.NewEncoder(w).Encode(list); err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func DelAll(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(Database.DelAll()); err != nil {
		http.Error(w, err.Error(), 500)
	}
}
