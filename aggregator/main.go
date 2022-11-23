package main

import (
	"aggregator/items"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"golang.org/x/sync/semaphore"
	"log"
	"net/http"
	"time"
)

var sem1 = semaphore.NewWeighted(5)
var sem2 = semaphore.NewWeighted(5)

func ThreadWork1() {
	ctx := context.Background()
	for i := 0; i < 5; i++ {
		go func(index int) {
			for {
				if !items.DataList1.IsEmpty() {
					if err := sem1.Acquire(ctx, 1); err != nil {
						log.Fatal(err)
					}
					fmt.Println(index, "sent to Consumer.")
					SendToConsumer(items.DataList1.Dequeue())
					sem1.Release(1)
				}
				time.Sleep(time.Second * 3)
			}
		}(i)
	}
}

func ThreadWork2() {
	ctx := context.Background()
	for i := 0; i < 5; i++ {
		go func(index int) {
			for {
				if !items.DataList2.IsEmpty() {
					if err := sem2.Acquire(ctx, 1); err != nil {
						log.Fatal(err)
					}
					fmt.Println(index, "sent to Producer.")
					SendToProducer(items.DataList2.Dequeue())
					sem2.Release(1)
				}
				time.Sleep(time.Second * 3)
			}
		}(i)
	}
}

func main() {
	router := mux.NewRouter()

	go ThreadWork1()
	go ThreadWork2()

	router.HandleFunc("/fromProducer", FromProducerPost).Methods("POST")
	router.HandleFunc("/fromProducer", FromProducerGet).Methods("GET")

	router.HandleFunc("/fromConsumer", FromConsumerPost).Methods("POST")
	router.HandleFunc("/fromConsumer", FromConsumerGet).Methods("GET")

	err := http.ListenAndServe(":8001", router)
	if err != nil {
		return
	}
}

func SendToConsumer(send *items.MessageSend) {
	data, err := json.Marshal(send)
	if err != nil {
		log.Fatal(err)
	}
	response, err := http.Post("http://consumer:8002/fromAggregator", "application/json", bytes.NewBuffer(data))
	//fmt.Println("Message " + string(data) + " sent to the consumer.")
	if err != nil {
		fmt.Println("Could not make POST request to the consumer.")
		return
	}
	defer response.Body.Close()
}

func SendToProducer(send *items.MessageReceive) {
	data, err := json.Marshal(send)
	if err != nil {
		log.Fatal(err)
	}
	response, err := http.Post("http://producer:8000/fromAggregator", "application/json", bytes.NewBuffer(data))
	//fmt.Println("Message " + string(data) + " sent to the producer.")
	if err != nil {
		fmt.Println("Could not make POST request to the producer.")
	}
	defer response.Body.Close()
}

func FromProducerPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var message items.MessageSend
	_ = json.NewDecoder(r.Body).Decode(&message)
	items.DataList1.Enqueue(message)
	json.NewEncoder(w).Encode(&message)
}

func FromProducerGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items.DataList1)
}

func FromConsumerPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var message items.MessageReceive
	_ = json.NewDecoder(r.Body).Decode(&message)
	items.DataList2.Enqueue(message)
	json.NewEncoder(w).Encode(&message)
}

func FromConsumerGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items.DataList2)
}
