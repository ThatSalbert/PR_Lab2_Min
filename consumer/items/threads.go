package items

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/sync/semaphore"
	"log"
	"net/http"
	"time"
)

var sem = semaphore.NewWeighted(5)

func ThreadWork() {
	ctx := context.Background()
	for i := 0; i < 5; i++ {
		go func(index int) {
			for {
				if !DataList.IsEmpty() {
					if err := sem.Acquire(ctx, 1); err != nil {
						log.Fatal(err)
					}
					fmt.Println(index, "sent to Consumer.")
					Post(GenMessageReceive(DataList.Dequeue()))
					sem.Release(1)
				}
				time.Sleep(time.Second * 3)
			}
		}(i)
	}
}

func Post(message *MessageReceive) {
	data, err := json.Marshal(message)
	if err != nil {
		log.Fatal(err)
	}
	response, err := http.Post("http://aggregator:8001/fromConsumer", "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Could not make POST request to the aggregator.")
	}
	defer response.Body.Close()
}
