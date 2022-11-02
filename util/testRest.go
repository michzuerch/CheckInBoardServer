package util

import (
	// "encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func TestRest() {
	currencies := "EURUSD,GBPUSD"
	apiKey := "P5WqKpA4crIZSsNf8llw"
	url := "https://marketdata.tradermade.com/api/v1/live?currency=" + currencies + "&api_key=" + apiKey

	resp, getErr := http.Get(url)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	fmt.Println(string(body))
}
