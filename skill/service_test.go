package skill

import (
	"reflect"
	"testing"
)

func TestNewWordCountSkill(t *testing.T) {
	type args struct {
		stopWords []string
	}
	tests := []struct {
		name string
		args args
		want *WordCountSkill
	}{
		{
			name: "TestNewWordCountSkill",
			args: args{
				stopWords: []string{"a", "b", "c"},
			},
			want: &WordCountSkill{
				stopWords: []string{"a", "b", "c"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWordCountSkill(tt.args.stopWords); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWordCountSkill() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWordCountSkill_Apply(t *testing.T) {
	type fields struct {
		stopWords []string
	}
	type args struct {
		body Body[map[string]string]
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Body[map[string]map[string]int]
	}{
		{
			name: "TestWordCountSkill_Apply",
			fields: fields{
				stopWords: []string{"a", "b", "c"},
			},
			args: args{
				body: Body[map[string]string]{
					Values: []Record[map[string]string]{
						{
							RecordID: "1",
							Data:     map[string]string{"content": "a b c d e f"},
						},
					},
				},
			},
			want: Body[map[string]map[string]int]{
				Values: []Record[map[string]map[string]int]{
					{
						RecordID: "1",
						Data: map[string]map[string]int{
							"content": {
								"d": 1,
								"e": 1,
								"f": 1,
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := WordCountSkill{
				stopWords: tt.fields.stopWords,
			}
			if got := w.Apply(tt.args.body); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Apply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWordCountSkill_countWords(t *testing.T) {
	type fields struct {
		stopWords []string
	}
	type args struct {
		data map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]map[string]int
	}{
		{
			name: "TestWordCountSkill_countWords",
			fields: fields{
				stopWords: []string{"a", "b", "c"},
			},
			args: args{
				data: map[string]string{"content": "a b c d e f"},
			},
			want: map[string]map[string]int{
				"content": {
					"d": 1,
					"e": 1,
					"f": 1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := WordCountSkill{
				stopWords: tt.fields.stopWords,
			}
			if got := w.countWords(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWordCountSkill_isStopWord(t *testing.T) {
	type fields struct {
		stopWords []string
	}
	type args struct {
		word string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "TestWordCountSkill_isStopWord",
			fields: fields{
				stopWords: []string{"a", "b", "c"},
			},
			args: args{
				word: "a",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := WordCountSkill{
				stopWords: tt.fields.stopWords,
			}
			if got := w.isStopWord(tt.args.word); got != tt.want {
				t.Errorf("isStopWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
