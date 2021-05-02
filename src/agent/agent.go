package agent

import (
	"log"
	"onair/camera"
	"onair/light"
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
				light.ColorCycleStart("8", 0, 255, 255)
			} else {
				light.ColorCycleStop("8")
			}
		}

		time.Sleep(10 * time.Second)
	}
}
