package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	intervalStr := os.Getenv("JOB_INTERVAL")
	if intervalStr == "" {
		panic("$JOB_INTERVAL must be set")
	}
	interval, err := strconv.Atoi(intervalStr)
	if err != nil {
		panic("$JOB_INTERVAL must be an int")
	}

	url := os.Getenv("JOB_URL")
	if url == "" {
		panic("$JOB_URL must be set")
	}

	loop(interval, url)
}

func loop(interval int, url string) {
	lastTime := time.Now().Unix()
	elapsedMinutes := 0

	for {
		now := time.Now().Unix()

		secSinceLastTime := now - lastTime
		elapsedMinutes += int(secSinceLastTime) / 60

		if elapsedMinutes >= interval {
			webhook(url)
			elapsedMinutes = 0
		}

		time.Sleep(time.Minute)
		lastTime = now
	}
}

func webhook(url string) {
	_, err := http.Post(url, "", nil)

	if err != nil {
		log.Println(err)
	}
}
