package web

import (
	"github.com/gin-gonic/gin"
	"github.com/tiredsosha/admin/protocols"
	"github.com/tiredsosha/admin/tools/formater"
	"github.com/tiredsosha/admin/tools/logger"

	config "github.com/tiredsosha/admin/tools/configurator"
)

func powerPc(c *gin.Context) {
	var data JsonCommand

	// Bind JSON and validate
	if err := c.ShouldBindJSON(&data); err != nil {
		logger.Error.Println("Invalid input:", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	logger.Info.Println("request data -", data)

	if data.Command == "on" {
		go protocols.SendWOL(config.FindPC(data.Zone, "mac"))
	} else {
		go protocols.SendGet(formater.CustomStr(
			"http://{ip}:3001/off",
			map[string]any{"ip": config.FindPC(data.Zone, "ip")}),
		)
	}

	c.JSON(200, gin.H{
		"message": "OK",
	})
}
