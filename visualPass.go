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
	Mag         float32 `json:"mag"`
}

func visualPass() {
	var id int
	fmt.Println("Enter NORAD ID: ")
	fmt.Scan(&id)

	apiKey := setKey()
	encoded := url.QueryEscape(apiKey)

	url := fmt.Sprintf("https://api.n2yo.com/rest/v1/satellite/visualpasses/%d/41.702/-76.014/100/2/300/&apiKey=%s", id, encoded)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Error requesting Sat data:", err)
	}
	resp.Close = true
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
	if len(passInfo.Pass) == 0 {
		fmt.Println("Satellite not visible from your location for next 2 days.")
	}
	for _, i := range passInfo.Pass {
		fmt.Println("Starting Azimuth of Pass (degrees):", i.StartAz)
		fmt.Println("Direction:", i.StartAzComp)
		fmt.Println("Elevation (degrees):", i.StartEl)
		fmt.Println(time.Unix(int64(i.StartUTC), 0))
		fmt.Println()

		fmt.Println("Ending Azimuth of Pass (degrees):", i.EndAz)
		fmt.Println("Direction at End of Pass:", i.EndAzComp)
		fmt.Println(time.Unix(int64(i.EndUTC), 0))
		if i.Mag == 100000 {
			fmt.Println("Magnitude cannot be determined.")
		} else {
			fmt.Println("Apparent Magnitude:", i.Mag)
		}
	}
}
