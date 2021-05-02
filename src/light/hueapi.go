package light

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// todo: configuration
const (
	hueUsername = "kfZqr9nqnhINjSKYlPXQ4R6TacR9nPE5Q9UOOC14"
	hueBridgeIp = "10.0.0.112"
)

func SetState(lightNumber string, on bool, hue int, sat int, bri int, transitionTime int) {
	url := fmt.Sprintf("http://%s/api/%s/lights/%s/state", hueBridgeIp, hueUsername, lightNumber)

	dto := map[string]interface{}{
		"on":             on,
		"hue":            hue,
		"sat":            sat,
		"bri":            bri,
		"transitiontime": transitionTime,
	}

	jsonDto, err := json.Marshal(dto)
	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonDto))
	if err != nil {
		log.Fatalln(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
}
