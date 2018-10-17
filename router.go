package main

import (
	"net/http"
	"encoding/json"
	"piscasion bot/model"

	"github.com/gorilla/mux"
)

type route struct {
	Name string
	Method string
	Pattern string
	Handler http.HandlerFunc
}

var routes = []route {
	route{
		"Index",
		"Get",
		"/",
		Index,
	},
	route{
		"ready",
		"Get",
		"/probe/ready",
		Ready,
	},
	route{
		"alive",
		"Get",
		"/probe/alive",
		Ready,
	},
	route{
		"turn",
		"Post",
		"/turn",
		Turn,
	},
}

func Index(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode("Index"); err != nil {
		json.NewEncoder(w).Encode(err)
	}
}

func Ready(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func Turn(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
    var state State
	err := decoder.Decode(&state)
	if err != nil {
		json.NewEncoder(w).Encode("Invalid payload" + err.Error())
	}

	json.NewEncoder(w).Encode(state)
}

func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	for _, route := range routes {
		router.Methods(route.Method).
		Path(route.Pattern).
		Name(route.Name).
		Handler(LogWrapper(AuthMiddleware(route.Handler), route.Name))
	}

	return router;
}
