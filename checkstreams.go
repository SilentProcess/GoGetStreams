package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Little program for checking your favorite streamers

var client = &http.Client{Timeout: 10 * time.Second}

type ReturnData struct {
	Streams []streams
}

type streams struct {
	StreamType  string  `json:"stream_type"`
	ChannelName channel `json:"channel"`
}

type channel struct {
	StreamerName string `json:"name"`
}

func main() {
	start := time.Now()

	MyData := new(ReturnData)
	getJson("https://api.twitch.tv/kraken/streams?channel=boxbox,lirik,imaqtpie&stream_type=live&client_id=InsertClientIDHere", MyData)
	fmt.Println(MyData.Streams)

	elapsed := time.Since(start)
	fmt.Printf("Total execution time: %s", elapsed)
}

// fucntion for deconding json
func getJson(url string, target interface{}) error {
	r, err := client.Get(url)
	if err != nil {
		return err
	}
	// fmt.Println(*r) // For HTTP status messages
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
