package Server

import (
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/ToDo/Note", GetAll).
		Methods("GET")
	r.HandleFunc("/ToDo/Note", CreateNote).
		Methods("POST")
	r.HandleFunc("/ToDo/Note/{id}", SetExecution).
		Methods("PUT")
	r.HandleFunc("/ToDo/CancelNote/{id}", SetExecution).
		Methods("PUT")
	r.HandleFunc("/ToDo/DeleteNote/{id}", DeleteNote).
		Methods("DELETE")
	r.HandleFunc("/ToDo/DeleteAll/пше", DelAll).
		Methods("DELETE")
	return r
}
