package main

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	MainURL     string = "https://computingsociety.co.uk"
	SurpriseURL string = "https://www.youtube.com/watch?v=mx86-rTclzA"
)

func main() {
	r := gin.Default()
	r.GET("/roulette", roulette)
	r.Run("localhost:4444")
}

func roulette(c *gin.Context) {
	n := rand.Intn(9)
	// 1/10 chance of Rick Roll for great results
	if n == 4 {
		c.Redirect(http.StatusTemporaryRedirect, SurpriseURL)
	} else {
		c.Redirect(http.StatusTemporaryRedirect, MainURL)
	}
}
