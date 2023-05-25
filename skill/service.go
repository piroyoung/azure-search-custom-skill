package skill

type Skill[S, T any] struct {
	mutation func(S) (T, error)
}

func NewSkill[S, T any](mutation func(S) (T, error)) *Skill[S, T] {
	return &Skill[S, T]{mutation: mutation}
}

func NewSkillNoErr[S, T any](mutation func(S) T) *Skill[S, T] {
	return &Skill[S, T]{mutation: func(s S) (T, error) {
		return mutation(s), nil
	}}
}

func (g Skill[S, T]) Apply(body Body[S]) Body[T] {
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
