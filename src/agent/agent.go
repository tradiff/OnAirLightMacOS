package agent

import (
	"log"
	"math/rand"
	"onair/camera"
	"onair/light"
	"onair/util"
	"time"
)

var isCameraOnPrev = false

func Start() {
	log.Println("Hello world!")
	rand.Seed(time.Now().UnixNano())
	for {
		isCameraOn := camera.InvokeGetCameraState()
		if isCameraOn != isCameraOnPrev {
			log.Printf("Camera:%t", isCameraOn)
			isCameraOnPrev = isCameraOn
			if isCameraOn {
				startColor := rand.Intn(len(light.ColorTable) - 1)
				light.ColorCycleStart(util.Config.HueLightNumber, startColor)
			} else {
				light.ColorCycleStop(util.Config.HueLightNumber)
			}
		}

		time.Sleep(10 * time.Second)
	}
}
