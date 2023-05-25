package main

import (
	"github.com/gin-gonic/gin"
	"strings"

	"github.com/piroyoung/azure-search-custom-skill/skill"
)

func main() {
	lowerSkill := skill.NewSkill(func(s string) (string, error) {
		return strings.ToLower(s), nil
	})

	r := gin.Default()
	r.POST("/v1/skills/lower", func(c *gin.Context) {
		var body skill.Body[string]
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		result := lowerSkill.Apply(body)
		c.JSON(200, result)
	})
	r.Run()
}
