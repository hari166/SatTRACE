package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type post struct {
	Pos  []getPos `json:"positions"`
	Info struct {
		Satname string `json:"satname"`
	}
}
type getPos struct {
	Lat  float64 `json:"satlatitude"`
	Long float64 `json:"satlongitude"`
	Alt  float64 `json:"sataltitude"`
}

func satPos() {
	var id int
	fmt.Println("Enter NORAD ID: ")
	fmt.Scan(&id)

	apiKey := setKey()
	encoded := url.QueryEscape(apiKey)

	url := fmt.Sprintf("https://api.n2yo.com/rest/v1/satellite/positions/%d/28.613/77.209/200/1/&apiKey=%s", id, encoded)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Error requesting Sat data:", err)
	}
	resp.Close = true
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Unexpected response status: %s", resp.Status)
	}

	var position post
	err = json.NewDecoder(resp.Body).Decode(&position)
	if err != nil {
		log.Fatal("Error parsing JSON:", err)
	}
	fmt.Println("Satellite Name:", position.Info.Satname)
	for _, v := range position.Pos {
		fmt.Println("Latitude:", v.Lat)
		fmt.Println("Longitude:", v.Long)
		fmt.Println("Altiude:", v.Alt, "km")
	}
}
