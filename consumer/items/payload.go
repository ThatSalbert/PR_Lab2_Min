package items

import "time"

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

func GenMessageReceive(messageSend *MessageSend) *MessageReceive {
	var message = new(MessageReceive)
	message.Id = messageSend.Id
	message.CreatedTime = messageSend.CreatedTime
	message.MessageString = messageSend.MessageString
	message.ResponseTime = time.Now().Unix()
	message.ResponseString = "Hello there."
	return message
}
