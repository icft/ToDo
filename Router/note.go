package Router

import (
	"ToDo/Server"
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/ToDo/Note", Server.GetAll).
		Methods("GET")
	r.HandleFunc("/ToDo/Note", Server.CreateProblem).
		Methods("POST")
	r.HandleFunc("/ToDo/Note/{id}", Server.SetExecution).
		Methods("PUT")
	r.HandleFunc("/ToDo/CancelNote/{id}", Server.SetExecution).
		Methods("PUT")
	r.HandleFunc("/ToDo/DeleteNote/{id}", Server.DeleteProblem).
		Methods("DELETE")
	r.HandleFunc("/ToDo/DeleteAll/{id}", Server.DelAll).
		Methods("DELETE")
	return r
}
