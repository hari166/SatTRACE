package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type TLEData struct {
	Data string `json:"tle"`
	Info struct {
		Satname string `json:"satname"`
	}
}

func parserTLE() {
	var id int
	fmt.Println("Enter NORAD ID: ")
	fmt.Scan(&id)
	apiKey := setKey()
	encoded := url.QueryEscape(apiKey)
	getTLE(encoded, id)
}
func getTLE(encoded string, id int) {
	url := fmt.Sprintf("https://api.n2yo.com/rest/v1/satellite/tle/%d&apiKey=%s", id, encoded)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Error requesting TLE data:", err)
	}
	resp.Close = true
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Unexpected response status: %s", resp.Status)
	}

	var tleData struct {
		TLEData
	}
	err = json.NewDecoder(resp.Body).Decode(&tleData)
	if err != nil {
		log.Fatal("Error parsing JSON:", err)
	}
	fmt.Println("Satellite Name:", tleData.Info.Satname)
	fmt.Println(tleData.Data)
	fmt.Print("\n")
	parsed := strings.Split(tleData.Data, "\n")

	if len(parsed) < 2 {
		fmt.Println("Decomissioned satellite. No longer in orbit.")
		return
	}

	line1 := parsed[0]
	line2 := parsed[1]

	eccentricityStr := strings.TrimSpace(line2[26:34])
	eccentricity, err := strconv.ParseFloat(fmt.Sprintf("0.%s", eccentricityStr), 64)
	if err != nil {
		fmt.Println("Error parsing eccentricity:", err)
		return
	}
	fmt.Printf("Eccentricity: %.7f\n", eccentricity)

	year := strings.TrimSpace(line1[9:11])
	if year >= `70` && year <= `99` {
		year = strconv.Itoa(19) + year
	} else {
		year = strconv.Itoa(20) + year
	}
	fmt.Println("Year of Launch:", year)

	rightAscension := strings.TrimSpace(line2[17:26])
	fmt.Println("Right Ascension:", rightAscension+"°")

	meanAnomaly := strings.TrimSpace(line2[53:58])
	revPerDay, err := strconv.ParseFloat(meanAnomaly, 64)
	if err != nil {
		fmt.Println("Error parsing Mean Anomaly:", err)
		return
	}
	fmt.Println("Revolutions/Day:", revPerDay)

	epochYear := strings.TrimSpace(line1[18:20])
	if year >= `70` && year <= `99` {
		epochYear = strconv.Itoa(19) + epochYear
	} else {
		epochYear = strconv.Itoa(20) + epochYear
	}
	fmt.Println("Epoch Year:", epochYear)
	fmt.Println("Epoch:", strings.TrimSpace(line1[20:30]))

	argOfPeriapsis := strings.TrimSpace(line2[42:52])
	fmt.Println("Argument of Periapsis:", argOfPeriapsis+"°")

	fmt.Println("Inclination:", strings.TrimSpace(line2[10:16])+"°")

}
