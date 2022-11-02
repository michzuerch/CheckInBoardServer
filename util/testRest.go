package util

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type data struct {
	Endpoint       string                   `json:"endpoint"`
	Quotes         []map[string]interface{} `json:"quotes"`
	Requested_time string                   `json:"requested_time"`
	Timestamp      int32                    `json:"timestamp"`
}

func TestRest() {
	currencies := "CHFEUR,EURUSD,GBPUSD"
	api_key := "P5WqKpA4crIZSsNf8llw"
	url := "https://marketdata.tradermade.com/api/v1/live?currency=" + currencies + "&api_key=" + api_key

	resp, getErr := http.Get(url)

	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	fmt.Println(string(body))
	data_obj := data{}

	jsonErr := json.Unmarshal(body, &data_obj)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println("endpoint", data_obj.Endpoint, "requested time", data_obj.Requested_time, "timestamp", data_obj.Timestamp)

	for key, value := range data_obj.Quotes {
		fmt.Println(key)
		fmt.Println(value)
		fmt.Println("symbol", value["base_currency"], value["quote_currency"], "bid", value["bid"], "ask", value["ask"], "mid", value["mid"])
	}
}
