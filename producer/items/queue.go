package items

var DataList = &Queue{}

type Queue struct {
	Elements []MessageReceive
}

func (q *Queue) Enqueue(order MessageReceive) {
	q.Elements = append(q.Elements, order)
}

func (q *Queue) SearchFor(MessageId int) *MessageReceive {
	for _, message := range q.Elements {
		if MessageId == message.Id {
			return &message
		}
	}
	return nil
}

func (q *Queue) IsEmpty() bool {
	return len(q.Elements) == 0
}
