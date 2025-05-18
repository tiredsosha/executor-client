package web

import (
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/tiredsosha/executor-client/tools/logger"
)

func restart(c *gin.Context) {
	logger.Info.Println("restart request")

	if err := exec.Command("cmd", "/C", "shutdown", "/r", "/f").Run(); err != nil {
		logger.Error.Println("Failed to initiate shutdown:", err)
	}

	c.JSON(200, gin.H{
		"message": "OK",
	})
}
