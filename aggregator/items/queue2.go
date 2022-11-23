package items

var DataList2 = &Queue2{}

type Queue2 struct {
	Elements []MessageReceive
}

func (q *Queue2) Enqueue(message MessageReceive) {
	q.Elements = append(q.Elements, message)
}

func (q *Queue2) Dequeue() *MessageReceive {
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

func (q *Queue2) SearchFor(MessageId int) *MessageReceive {
	for _, message := range q.Elements {
		if MessageId == message.Id {
			return &message
		}
	}
	return nil
}

func (q *Queue2) IsEmpty() bool {
	return len(q.Elements) == 0
}

func (q *Queue2) GetSize() int {
	return len(q.Elements)
}
