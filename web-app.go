package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type day int

const (
	_ = iota
	Monday
	Tuesday
)

func main() {
	var asa string = "a"
	fmt.Printf(asa)
	a := http.StatusBadRequest
	fmt.Printf("%d\n", a)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
