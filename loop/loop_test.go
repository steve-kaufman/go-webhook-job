package loop_test

import (
	"testing"
	"time"

	"github.com/steve-kaufman/go-webook-job/loop"
)

type LoopTest struct {
	name string

	interval time.Duration
	waitTime time.Duration
}

var loopTests = []LoopTest{
	{
		name: "10 Milliseconds",

		interval: time.Millisecond,

		waitTime: time.Millisecond * 10,
	},
	{
		name: "100 Milliseconds",

		interval: time.Millisecond,

		waitTime: time.Millisecond * 100,
	},
	{
		name: "1 second",

		interval: time.Millisecond * 10,

		waitTime: time.Second,
	},
	{
		name: "500 ms",

		interval: time.Millisecond * 100,

		waitTime: time.Millisecond * 500,
	},
}

func TestLoop(t *testing.T) {
	for _, tc := range loopTests {
		t.Run(tc.name, func(t *testing.T) {
			for i := 1; i < 10; i += 1 {
				runTest(t, tc)
			}
		})
	}
}

func runTest(t *testing.T, tc LoopTest) {
	timesCalled := 0
	callback := func() {
		timesCalled += 1
	}

	loop := loop.New(tc.interval, callback)

	before := time.Now()

	go loop.Start()
	time.Sleep(tc.waitTime)
	loop.Stop()

	deltaTime := time.Since(before)

	expectedTimesCalled := int(deltaTime / tc.interval)

	upperLimit := int(float64(expectedTimesCalled)*1.05) + 1
	lowerLimit := int(float64(expectedTimesCalled)*0.95) - 2

	if timesCalled > upperLimit {
		t.Fatalf("Should've been called at most %d times; Was called: %d times", upperLimit, timesCalled)
	}
	if timesCalled < lowerLimit {
		t.Fatalf("Should've been called at least %d times; Was called: %d times", lowerLimit, timesCalled)
	}
}
