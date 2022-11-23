package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"producer/items"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	router := mux.NewRouter()

	go items.GenThreads(5)

	router.HandleFunc("/fromAggregator", FromAggregatorPost).Methods("POST")
	router.HandleFunc("/fromAggregator", FromAggregatorGet).Methods("GET")

	err := http.ListenAndServe(":8000", router)
	if err != nil {
		return
	}
}

func FromAggregatorPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var message items.MessageReceive
	_ = json.NewDecoder(r.Body).Decode(&message)
	items.DataList.Enqueue(message)
	json.NewEncoder(w).Encode(&message)
}

func FromAggregatorGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items.DataList)
}
