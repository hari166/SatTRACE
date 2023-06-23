package main

import (
	"os"
)

var apiKey string

func setKey() string {
	apiKey = os.Getenv("N2YO")
	return apiKey
}
