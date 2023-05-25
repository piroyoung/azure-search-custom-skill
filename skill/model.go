package skill

type Message struct {
	Value string `json:"message"`
}

func NewMessage(value string) Message {
	return Message{Value: value}
}

type Record[T any] struct {
	RecordID string       `json:"recordId"`
	Data     map[string]T `json:"data"`
	Errors   []Message    `json:"errors"`
	Warnings []Message    `json:"warnings"`
}

type Body[T any] struct {
	Values []Record[T] `json:"values"`
}
