package main

import (
	"github.com/gin-gonic/gin"
	"io"
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

	r := gin.Default()

	r.POST("/v1/skills/:name", func(c *gin.Context) {
		name := c.Param("name")
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		result, err := book.Apply(name, body)
		if err == skill.ErrNotFound {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		} else if err == skill.ErrParse {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		} else if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.Data(200, "application/json", result)
	})
	r.Run()
}
