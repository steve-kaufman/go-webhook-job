package main

import (
	"bytes"
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

func getBody() string {
	body := os.Args[1]
	return body
}

func getContentType() string {
	contentType := os.Getenv("HOOK_CONTENT_TYPE")
	if contentType == "" {
		contentType = "text/plain"
	}
	return contentType
}

func main() {
	interval := getInterval()
	url := getURL()
	body := getBody()
	contentType := getContentType()

	loop := loop.New(interval, func() {
		webhook(url, body, contentType)
	})
	loop.Start()
}

func webhook(url string, body string, contentType string) {
	println("Sending POST to:", url)
	_, err := http.Post(url, contentType, bytes.NewBufferString(body))

	if err != nil {
		log.Println(err)
		return
	}

	println("Done")
}
