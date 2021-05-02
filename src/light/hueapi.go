package light

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"onair/util"
)

func SetState(lightNumber string, on bool, hue int, sat int, bri int, transitionTime int) {
	url := fmt.Sprintf("http://%s/api/%s/lights/%s/state", util.Config.HueBridgeIp, util.Config.HueUsername, lightNumber)

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
