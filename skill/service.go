package skill

import "strings"

type Skill[S, T any] interface {
	Apply(Body[S]) Body[T]
}

type GeneralSkill[S, T any] struct {
	mutation func(S) (T, error)
}

func NewGeneralSkill[S, T any](mutation func(S) (T, error)) *GeneralSkill[S, T] {
	return &GeneralSkill[S, T]{mutation: mutation}
}

func (g GeneralSkill[S, T]) Apply(body Body[S]) Body[T] {
	result := make([]Record[T], len(body.Values))
	for i, record := range body.Values {
		result[i].RecordID = record.RecordID
		result[i].Data = make(map[string]T)
		for k, v := range record.Data {
			value, err := g.mutation(v)
			if err != nil {
				result[i].Errors = append(result[i].Errors, NewMessage(err.Error()))
			} else {
				result[i].Data[k] = value
			}
		}
	}
	return Body[T]{Values: result}
}

type WordCountSkill struct {
	stopWords []string
}

func NewWordCountSkill(stopWords []string) *WordCountSkill {
	return &WordCountSkill{stopWords: stopWords}
}

func (w WordCountSkill) Apply(body Body[string]) Body[map[string]int] {
	result := make([]Record[map[string]int], len(body.Values))
	for i, record := range body.Values {
		result[i] = Record[map[string]int]{
			RecordID: record.RecordID,
			Data:     w.countWords(record.Data),
		}
	}
	return Body[map[string]int]{Values: result}
}

func (w WordCountSkill) countWords(data map[string]string) map[string]map[string]int {
	result := make(map[string]map[string]int)
	for k, v := range data {
		result[k] = make(map[string]int)
		for _, word := range strings.Split(v, " ") {
			if !w.isStopWord(word) {
				result[k][word]++
			}
		}
	}
	return result
}

func (w WordCountSkill) isStopWord(word string) bool {
	for _, stopWord := range w.stopWords {
		if word == stopWord {
			return true
		}
	}
	return false
}
