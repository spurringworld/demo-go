package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Logs struct {
	Message string
}

func main() {
	//load ENV
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	helloMsg := os.Getenv("HELLOMSG")
	svcListen := os.Getenv("SERVER_LISTENING")
	// gin-web
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// save logs func
	r.GET("/hello", func(c *gin.Context) {
		fmt.Println("-- " + helloMsg)
		c.JSON(200, gin.H{
			"message": helloMsg,
		})
	})

	// welcome message
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("  Listen on ", svcListen)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	r.Run(svcListen)

}
