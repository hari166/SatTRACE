package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

const (
	Host     = "localhost"
	Port     = 5432
	User     = "postgres"
	Password = "root"
	Dbname   = "sat_db"
)

func populateDB() {
	psConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", Host, Port, User, Password, Dbname)
	db, err := sql.Open("postgres", psConn)
	check(err)

	err = db.Ping()
	check(err)

	key := os.Getenv("N2YO")
	encoded := url.QueryEscape(key)

	defer db.Close()

	count := 0

	//28000->28999
	//28000->29999
	//30000->30999
	//31000->31999

	for id := 28000; id <= 28999; id++ {

		url := fmt.Sprintf("https://api.n2yo.com/rest/v1/satellite/tle/%d&apiKey=%s", id, encoded)
		res, err := http.Get(url)
		check(err)

		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			log.Fatalf("Error: %d", res.StatusCode)
		}

		var response struct {
			Info Satellite `json:"info"`
		}

		err = json.NewDecoder(res.Body).Decode(&response)
		check(err)

		valID := response.Info.SatID
		valName := response.Info.Satname

		stmt := `insert into "satellites" values($1,$2)`

		_, err = db.Exec(stmt, valID, valName)
		check(err)

		count++
		fmt.Printf("\rProcessing: %d entries", count)

	}
	fmt.Println("\nSuccess")
}
