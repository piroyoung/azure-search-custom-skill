package skill

import "strings"

type Skill[S, T any] interface {
	Apply(Body[S]) Body[T]
}

type WordCountSkill struct {
	stopWords []string
}

func NewWordCountSkill(stopWords []string) *WordCountSkill {
	return &WordCountSkill{stopWords: stopWords}
}

func (w WordCountSkill) Apply(body Body[map[string]string]) Body[map[string]map[string]int] {
	result := make([]Record[map[string]map[string]int], len(body.Values))
	for i, record := range body.Values {
		result[i] = Record[map[string]map[string]int]{
			RecordID: record.RecordID,
			Data:     w.countWords(record.Data),
		}
	}
	return Body[map[string]map[string]int]{Values: result}
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
