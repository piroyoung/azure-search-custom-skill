package main

import (
	"github.com/gin-gonic/gin"

	"github.com/piroyoung/azure-search-custom-skill/skill"
)

func main() {
	stopWords := []string{"a", "b", "c"}
	countSkill := skill.NewWordCountSkill(stopWords)

	r := gin.Default()
	r.POST("/v1/skills/count", func(c *gin.Context) {
		var body skill.Body[map[string]string]
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		result := countSkill.Apply(body)
		c.JSON(200, result)
	})
	r.Run()
}
