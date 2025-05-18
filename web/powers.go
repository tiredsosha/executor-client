package web

import (
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/tiredsosha/executor-client/tools/logger"
)

func powerOff(c *gin.Context) {
	if err := exec.Command("cmd", "/C", "shutdown", "/s", "/f").Run(); err != nil {
		logger.Error.Println("Failed to initiate shutdown:", err)
	}

	c.JSON(200, gin.H{
		"message": "OK",
	})
}
