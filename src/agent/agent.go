package agent

import (
	"log"
	"onair/camera"
	"onair/light"
	"onair/util"
	"time"
)

var isCameraOnPrev = false

func Start() {
	log.Println("Hello world!")
	for {
		isCameraOn := camera.InvokeGetCameraState()
		if isCameraOn != isCameraOnPrev {
			log.Printf("Camera:%t", isCameraOn)
			isCameraOnPrev = isCameraOn
			if isCameraOn {
				light.ColorCycleStart(util.Config.HueLightNumber, 0)
			} else {
				light.ColorCycleStop(util.Config.HueLightNumber)
			}
		}

		time.Sleep(10 * time.Second)
	}
}
