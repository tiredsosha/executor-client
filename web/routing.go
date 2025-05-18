package web

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tiredsosha/executor-client/tray"
)

func router(server *gin.Engine) {
	server.GET("/off", powerOff)
	server.GET("/restart", restart)
	server.POST("/sound", soundChange)
	server.GET("/getTest", testGet)
	server.POST("/postTest", testPost)
}

func StartServer(port int) {
	portSrt := ":" + strconv.Itoa(port)
	route := gin.Default()
	router(route)

	server := &http.Server{
		Addr:           portSrt,
		Handler:        route,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	tray.Conn = true

	server.ListenAndServe()
}
