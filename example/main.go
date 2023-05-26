package main

import (
	"github.com/piroyoung/azure-search-custom-skill/service"
	"strings"

	"github.com/piroyoung/azure-search-custom-skill/skill"
)

func main() {
	lowerSkill := skill.NewSkillNoErr(strings.ToLower)
	upperSkill := skill.NewSkillNoErr(strings.ToUpper)
	splitSkill := skill.NewSkillNoErr(func(s string) []string {
		return strings.Split(s, " ")
	})
	wordCountSkill := skill.NewSkillNoErr(func(s string) map[string]int {
		result := make(map[string]int)
		for _, word := range strings.Split(s, " ") {
			result[word]++
		}
		return result
	})

	book := skill.NewBook()
	book.Register("lower", lowerSkill.Flatten())
	book.Register("upper", upperSkill.Flatten())
	book.Register("split", splitSkill.Flatten())
	book.Register("wordcount", wordCountSkill.Flatten())

	svc := service.NewCustomSkillService(book)
	svc.Run()
}
