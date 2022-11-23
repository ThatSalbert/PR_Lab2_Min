package items

var DataList = &Queue{}

type Queue struct {
	Elements []MessageSend
}

func (q *Queue) Enqueue(message MessageSend) {
	q.Elements = append(q.Elements, message)
}

func (q *Queue) Dequeue() *MessageSend {
	if q.IsEmpty() {
		return nil
	}
	message := q.Elements[0]
	if q.GetSize() == 1 {
		q.Elements = nil
		return &message
	}
	q.Elements = q.Elements[1:]
	return &message
}

func (q *Queue) SearchFor(MessageId int) *MessageSend {
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

func (q *Queue) GetSize() int {
	return len(q.Elements)
}
