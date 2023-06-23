package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

type visual struct {
	Info struct {
		Satname string `json:"satname"`
	}
	Pass []getPass `json:"passes"`
}
type getPass struct {
	StartAz     float64 `json:"maxAz"`
	StartAzComp string  `json:"maxAzCompass"`
	StartUTC    int     `json:"maxUTC"`
	EndUTC      int     `json:"endUTC"`
	StartEl     float64 `json:"maxEl"`
	EndAz       float64 `json:"endAz"`
	EndAzComp   string  `json:"endAzCompass"`
}

func visualPass() {
	var id int
	fmt.Println("Enter NORAD ID: ")
	fmt.Scan(&id)

	apiKey := setKey()
	encoded := url.QueryEscape(apiKey)

	url := fmt.Sprintf("https://api.n2yo.com/rest/v1/satellite/visualpasses/%d/41.702/-76.014/100/1/300/&apiKey=%s", id, encoded)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Error requesting Sat data:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Unexpected response status: %s", resp.Status)
	}

	var passInfo visual
	err = json.NewDecoder(resp.Body).Decode(&passInfo)
	if err != nil {
		log.Fatal("Error parsing JSON:", err)
	}
	fmt.Println("Satellite Name:", passInfo.Info.Satname)
	fmt.Println()
	for _, i := range passInfo.Pass {
		fmt.Println("Starting Azimuth of Pass (degrees):", i.StartAz)
		fmt.Println("Direction:", i.StartAzComp)
		fmt.Println("Elevation (degrees):", i.StartEl)
		fmt.Println(time.Unix(int64(i.StartUTC), 0))
		fmt.Println()

		fmt.Println("Ending Azimuth of Pass (degrees):", i.EndAz)
		fmt.Println("Direction at End of Pass:", i.EndAzComp)
		fmt.Println(time.Unix(int64(i.EndUTC), 0))
	}
}
