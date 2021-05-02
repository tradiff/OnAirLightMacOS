package agent

import (
	"log"
	"onair/camera"
	"time"
)

func Start() {

	i := 0
	for {
		isCameraOn := camera.InvokeGetCameraState()
		log.Println("camera on:", isCameraOn)

		i++
		log.Println("tick", i)
		time.Sleep(1 * time.Second)
	}
}
