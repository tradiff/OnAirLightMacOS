package light

import (
	"log"
	"onair/util"
	"time"
)

var colorCycleRunning = false
var colorTableIdx = 0
var colorCycleQuitChannel = make(chan bool)

var ColorTable = []struct {
	Label     string
	Hue       int
	Immediate bool
}{
	{Label: "Red(0)", Hue: 0, Immediate: true},
	{Label: "Green", Hue: 17408, Immediate: false},
	{Label: "Cyan", Hue: 37120, Immediate: false},
	{Label: "Blue", Hue: 41728, Immediate: false},
	{Label: "Pink", Hue: 57344, Immediate: false},
	{Label: "Red(65535)", Hue: 65535, Immediate: false},
}

func ColorCycleStart(lightNumber string, startHue int) {
	log.Printf("Starting color cycle for light:%s", lightNumber)
	if colorCycleRunning {
		log.Printf("already running!")
		return
	}
	colorCycleQuitChannel = make(chan bool, 1)
	go colorCycle(lightNumber, startHue)
}

func ColorCycleStop(lightNumber string) {
	log.Printf("Stopping color cycle for light:%s", lightNumber)
	log.Printf("Setting light:%s OFF", lightNumber)
	SetState(lightNumber, false, 0, 0, 0, 0)
	colorCycleQuitChannel <- true
}

func colorCycle(lightNumber string, startColor int) {
	colorCycleRunning = true
	colorTableIdx = startColor
	defer func() {
		colorCycleRunning = false
	}()

	for {
		colorTableIdx++
		if colorTableIdx >= len(ColorTable) {
			colorTableIdx = 0
		}
		var color = ColorTable[colorTableIdx]

		transitionTime := 0
		if color.Immediate {
			transitionTime = 0
		} else {
			transitionTime = 500
		}

		log.Printf("Setting light:%s color:%s tt:%d", lightNumber, color.Label, transitionTime)
		SetState(lightNumber, true, color.Hue, util.Config.ColorSaturation, util.Config.ColorBrightness, transitionTime)

		select {
		case <-colorCycleQuitChannel:
			// quit the goroutine
			return
		case <-time.After(time.Duration(util.Config.ColorCycleInterval) * time.Second):
			// continue to the next iteration
			break
		}

	}
}
