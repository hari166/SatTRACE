package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"

	_ "github.com/lib/pq"
)

type Satellite struct {
	SatID   int    `json:"satid"`
	Satname string `json:"satname"`
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "****"
	dbname   = "sat_db"
)

var userInput string

func satDB() {
	psConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psConn)
	check(err)

	err = db.Ping()
	check(err)

	fmt.Println("Enter Satellite Name")
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	userInput = scan.Text()

	key := os.Getenv("N2YO")
	encoded := url.QueryEscape(key)

	defer db.Close()

	var value int
	stmt := `SELECT "satelliteID" FROM satellites WHERE LOWER("satelliteName") = LOWER($1) limit 1;`
	err = db.QueryRow(stmt, userInput).Scan(&value)
	check(err)
	fmt.Printf("NORAD ID: %d\n", value)
	getTLE(encoded, value)

}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
