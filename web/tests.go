package web

import (
	"github.com/gin-gonic/gin"
	"github.com/tiredsosha/executor-client/tools/logger"
)

// func testFunctionality() {
// 	var testJson protocols.UserCommand
// 	testJson.Req = "test"
// 	protocols.SendUdp("127.0.0.1", 8090, "restart")
// 	protocols.SendOsc("127.0.0.1", 8091, "/test", "restart")
// 	protocols.SendGet("http://127.0.0.1:8092")
// 	// protocols.SendPost("http://127.0.0.1:8092", testJson)
// }

func testGet(c *gin.Context) {
	// testFunctionality()

	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func testPost(c *gin.Context) {
	var data JsonNoID

	// Bind JSON and validate
	if err := c.ShouldBindJSON(&data); err != nil {
		logger.Error.Println("Invalid input:", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// testFunctionality()

	c.JSON(200, gin.H{
		"command": data.Command,
	})
}
