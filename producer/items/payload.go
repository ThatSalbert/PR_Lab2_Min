package items

import (
	"math/rand"
	"strconv"
	"time"
)

type MessageSend struct {
	Id            int    `json:"id"`
	CreatedTime   int64  `json:"created_time"`
	MessageString string `json:"message_string"`
}

type MessageReceive struct {
	Id             int    `json:"id"`
	CreatedTime    int64  `json:"created_time"`
	MessageString  string `json:"message_string"`
	ResponseTime   int64  `json:"response_time"`
	ResponseString string `json:"response_string"`
}

func GenMessageToSend() *MessageSend {
	var message = new(MessageSend)
	message.Id = rand.Intn(99999-10000) + 10000
	message.CreatedTime = time.Now().Unix()
	message.MessageString = "Message of ID: " + strconv.Itoa(message.Id) + " created on " + strconv.Itoa(int(message.CreatedTime)) + "."
	return message
}
