package main

import (
	"os"
	"strconv"
	"time"
)

func main() {
	errorCode := os.Getenv("DESIRED_EXIT_CODE")

	numericCode, err := strconv.Atoi(errorCode)
	if err != nil {
		panic("Unable to parse exit code.")
	}

	time.Sleep(2 * time.Second)

	os.Exit(numericCode)
}
