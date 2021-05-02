package agent

import (
	"log"
	"time"
)

func Start() {

	i := 0
	for {
		i++
		log.Println("tick", i)
		time.Sleep(1 * time.Second)
	}
}
