package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/steve-kaufman/go-webook-job/loop"
)

func getInterval() time.Duration {
	intervalStr := os.Getenv("JOB_INTERVAL")
	if intervalStr == "" {
		panic("$JOB_INTERVAL must be set")
	}
	interval, err := time.ParseDuration(intervalStr)
	if err != nil {
		panic("$JOB_INTERVAL must be an int")
	}

	return interval
}

func getURL() string {
	url := os.Getenv("HOOK_URL")
	if url == "" {
		panic("$HOOK_URL must be set")
	}
	return url
}

func main() {
	interval := getInterval()
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
