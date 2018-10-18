package main

import (
	"encoding/json"
	"net/http"
	"github.com/picasion/model"

	"github.com/gorilla/mux"
)

type route struct {
	Name    string
	Method  string
	Pattern string
	Handler httpHandler
}

type httpHandler func(http.ResponseWriter, *http.Request) error

var routes = []route{
	route{
		"Index",
		"Get",
		"/",
		index,
	},
	route{
		"ready",
		"Get",
		"/probe/ready",
		ready,
	},
	route{
		"alive",
		"Get",
		"/probe/alive",
		ready,
	},
	route{
		"turn",
		"Post",
		"/turn",
		turn,
	},
}

func index(w http.ResponseWriter, r *http.Request) error {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode("Index"); err != nil {
		return err
	}

	return nil
}

func ready(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	return nil
}

func turn(w http.ResponseWriter, r *http.Request) error {

	decoder := json.NewDecoder(r.Body)
	var state model.State
	err := decoder.Decode(&state)
	if err != nil {
		return err
	}

	json.NewEncoder(w).Encode(state)

	return nil
}

func createRouter() *mux.Router {
	router := mux.NewRouter()
	for _, route := range routes {
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(logMiddleware(authMiddleware(errorMiddleware(route.Handler))))
	}

	return router
}
