package web

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/itchyny/volume-go"

	"github.com/tiredsosha/executor-client/tools/logger"
)

func soundChange(c *gin.Context) {
	var data JsonNoID

	// Bind JSON and validate
	if err := c.ShouldBindJSON(&data); err != nil {
		logger.Error.Println("Invalid input:", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	logger.Info.Println("sound change request -", data)

	volumeInt, _ := strconv.Atoi(data.Command)
	volume.SetVolume(volumeInt)

	c.JSON(200, gin.H{
		"message": "OK",
	})
}
