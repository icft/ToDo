package Router

import (
	"ToDo/Server"
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/ToDo/problem", Server.GetAll).
		Methods("GET")
	r.HandleFunc("/ToDo/problem", Server.CreateProblem).
		Methods("POST")
	r.HandleFunc("/ToDo/problem/{id}", Server.SetExecution).
		Methods("PUT")
	r.HandleFunc("/ToDo/cancelProblem/{id}", Server.SetExecution).
		Methods("PUT")
	r.HandleFunc("/ToDo/deleteProblem/{id}", Server.DeleteProblem).
		Methods("DELETE")
	r.HandleFunc("/ToDo/deleteAll/{id}", Server.DelAll).
		Methods("DELETE")
	return r
}
