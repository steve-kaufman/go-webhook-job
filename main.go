package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/steve-kaufman/go-webook-job/loop"
)

func getInterval() int {
	intervalStr := os.Getenv("JOB_INTERVAL")
	if intervalStr == "" {
		panic("$JOB_INTERVAL must be set")
	}
	interval, err := strconv.Atoi(intervalStr)
	if err != nil {
		panic("$JOB_INTERVAL must be an int")
	}

	return interval
}

func getURL() string {
	url := os.Getenv("JOB_URL")
	if url == "" {
		panic("$JOB_URL must be set")
	}
	return url
}

func main() {
	interval := time.Duration(getInterval())
	url := getURL()

	loop := loop.New(interval, func() {
		webhook(url)
	})
	loop.Start()
}

func webhook(url string) {
	println("Sending POST to:", url)
	_, err := http.Post(url, "", nil)

	if err != nil {
		log.Println(err)
		return
	}

	println("Done")
}
