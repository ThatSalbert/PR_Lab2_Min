package items

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
