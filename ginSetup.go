package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func ServerSetup() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "ok")
	})

	if err := r.Run(); err != nil {
		log.Fatalln(err)
	}
}
