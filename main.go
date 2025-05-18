package main

import (
	"github.com/tiredsosha/executor-client/mosquitto"
	config "github.com/tiredsosha/executor-client/tools/configurator"
	"github.com/tiredsosha/executor-client/tools/logger"
	"github.com/tiredsosha/executor-client/tray"
	"github.com/tiredsosha/executor-client/web"
)

const (
	version = "1.0.0"
)

func main() {
	logger.LogInit(true)

	cfg := config.ConfInit()

	hostname := "client"
	topicPrefix := "executor/"

	logger.DebugLog(version, true, cfg.MqttOn, hostname, cfg.Broker, cfg.Username, cfg.Password, cfg.Port)

	// вот это нужно, но мы тоже будем с ним разбираться в конце
	go tray.TrayStart()

	if cfg.MqttOn {
		// так sub топик нам не нужен вообще, нам нужен только пуб топик. Мы наш стстум отправлять не будет. Так же надо убрать apps
		mqttData := mosquitto.MqttConf{
			ID:       hostname,
			Broker:   cfg.Broker,
			Username: cfg.Username,
			Password: cfg.Password,
			SubTopic: topicPrefix + hostname + "/commands/",
			PubTopic: topicPrefix,
			Icon:     &tray.Conn,
		}
		go mosquitto.StartBroker(mqttData)
	}

	web.StartServer(cfg.Port)
}
