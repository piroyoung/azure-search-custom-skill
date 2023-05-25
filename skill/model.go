package skill

type Record[T any] struct {
	RecordID string       `json:"recordId"`
	Data     map[string]T `json:"data"`
}

type Body[T any] struct {
	Values []Record[T] `json:"values"`
}
