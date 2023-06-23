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
		if input == 1 {
			getTLE()
		} else if input == 2 {
			satPos()
		} else if input == 3 {
			visualPass()
		} else if input == 4 {
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
