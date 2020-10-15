package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/project/:id", func(c *gin.Context) {
		id := c.Param("id")

		var body interface{}
		if err := c.BindJSON(&body); err != nil {
			c.AbortWithStatus(400)
			return
		}

		c.Status(200)
	})

	log.Fatal(r.Run())
}
