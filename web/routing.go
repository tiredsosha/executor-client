package web

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tiredsosha/admin/tray"
)

func router(server *gin.Engine) {
	server.GET("/on", powerPark)
	server.GET("/off", powerZone)
	server.GET("/restart", powerPc)
	server.POST("/sound", powerProjector)
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
