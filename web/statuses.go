package web

import (
	"github.com/gin-gonic/gin"
	"github.com/tiredsosha/executor-client/tools/logger"
)

func status(c *gin.Context) {
	logger.Info.Println("status request")

	c.JSON(200, gin.H{
		"message": "OK",
	})
}
