package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var f *os.File
	f = os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		_, err := strconv.ParseFloat(text, 64)
		if err != nil {
			if text == "STOP" {
				return
			} else {
				fmt.Println("Not a Float")
			}
		}
	}
}
