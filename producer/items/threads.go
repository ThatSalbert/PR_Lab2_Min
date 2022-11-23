package items

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func GenThreads(num int) {
	for i := 0; i < num; i++ {
		go ThreadWork()
	}
}

func ThreadWork() {
	for {
		data, err := json.Marshal(GenMessageToSend())
		if err != nil {
			log.Fatal(err)
		}
		response, err := http.Post("http://aggregator:8001/fromProducer", "application/json", bytes.NewBuffer(data))
		fmt.Println("Message " + string(data) + " sent to the aggregator.")
		if err != nil {
			fmt.Println("Could not make POST request to the aggregator.")
		}
		time.Sleep(time.Second * 2)
		defer response.Body.Close()
	}
}
