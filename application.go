package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "A Go Web Server")
	w.WriteHeader(200)
	m := make(map[string]string)
	m["message"] = "Hello World"
	response, err := json.Marshal(m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(response)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8000", router))
}
