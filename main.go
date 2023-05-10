package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Joke struct {
	Joke string `json:"joke"`
}

func fetchRandomJoke() (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://icanhazdadjoke.com/", nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var joke Joke
	err = json.NewDecoder(resp.Body).Decode(&joke)
	if err != nil {
		return "", err
	}
	return joke.Joke, nil
}

func main() {
	r := gin.Default()

	r.GET("/joke", func(c *gin.Context) {
		joke, err := fetchRandomJoke()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetch joke",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"joke": joke,
		})
	})
	port := 8080
	fmt.Printf("Joke Generator API is running on port %d\n", port)
	r.Run(":8080")
}
