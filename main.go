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

	r.GET("/hello", func(c *gin.Context) {
		fmt.Println("-- " + helloMsg)
		// 打印header信息
		// for k, v := range c.Request.Header {
		// 	fmt.Println(k, v)
		// }
		c.JSON(200, gin.H{
			"message":  helloMsg,
			"action":   "/hello",
			"agentIp":  c.RemoteIP(),
			"clientIp": c.ClientIP(),
		})
	})
	r.POST("/hello", func(c *gin.Context) {
		fmt.Println("-- " + helloMsg)
		c.JSON(200, gin.H{
			"message": helloMsg,
			"action":  "/hello",
		})
	})
	r.GET("/test", func(c *gin.Context) {
		fmt.Println("-- " + helloMsg)
		c.JSON(200, gin.H{
			"message": helloMsg,
			"action":  "/test",
		})
	})

	// welcome message
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("  Listen on ", svcListen)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	r.Run(svcListen)

}
