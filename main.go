package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	banner()
	for {
		options()
		var input int
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Invalid input:", err)
			continue
		}
		fmt.Scanln()
		if input == 1 {
			parserTLE()
		} else if input == 2 {
			satDB()
		} else if input == 3 {
			satPos()
		} else if input == 4 {
			visualPass()
		} else if input == 5 {
			populateDB()
		} else if input == 6 {
			fmt.Println("Exited.")
			os.Exit(0)
		} else {
			fmt.Println("Invalid input")
		}
		time.Sleep(2 * time.Second)
	}
}
func banner() {
	banner, err := os.ReadFile("banners/start.txt")
	if err != nil {
		fmt.Println("Cannot open file.")
		return
	}
	fmt.Println(string(banner))
}
func options() {
	options, err := os.ReadFile("banners/options.txt")
	if err != nil {
		fmt.Println("Cannot open file.")
		return
	}
	fmt.Println(string(options))
}
