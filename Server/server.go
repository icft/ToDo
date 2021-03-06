package Server

import (
	"ToDo/Database"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func CreateNote(w http.ResponseWriter, r *http.Request) {
	var jsonDecoder = json.NewDecoder(r.Body)
	var rec Database.Note
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

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	var vars = mux.Vars(r)
	Database.DelInDb(vars["id"])
	if err := json.NewEncoder(w).Encode(vars["id"]); err != nil {
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
	Database.DelAll()
}
