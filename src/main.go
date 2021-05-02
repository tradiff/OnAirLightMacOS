package main

import (
	"log"
	"onair/agent"
	"onair/camera"
	"onair/util"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	handleGetCameraState()

	_, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	startAgent()
}

func handleGetCameraState() {
	for _, arg := range os.Args[1:] {
		if arg == "--get-camera-state" {
			if camera.GetCameraState() {
				os.Exit(0)
			} else {
				os.Exit(1)
			}
		}
	}
}

func startAgent() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		os.Exit(0)
	}()

	agent.Start()
}
