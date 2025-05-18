package tray

import (
	"time"

	"github.com/tiredsosha/executor-client/tools/logger"

	"github.com/getlantern/systray"
)

var Conn bool = false
var icon bool

func onReady() {

	systray.SetIcon(grayIcon)
	systray.SetTitle("Executor Client")

	systray.SetTooltip("Executor")
	menuQuit := systray.AddMenuItem("QUIT", "Quit the whole app")

	go func() {
		<-menuQuit.ClickedCh
		systray.Quit()
	}()

	go func() {
		for {
			time.Sleep(3 * time.Second)

			if Conn == icon {
				continue
			}
			if Conn {
				systray.SetIcon(blueIcon)
			} else {
				systray.SetIcon(grayIcon)
			}

			icon = Conn
		}
	}()
}

func onExit() {
	logger.Error.Fatal("EXITING MANUALLY")
}

func TrayStart() {
	systray.Run(onReady, onExit)
}
