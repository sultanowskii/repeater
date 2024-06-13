package repeater

import (
	"math"
	"time"
)

type Repeater struct {
	interval    time.Duration
	retries     int
	checkResult resultChecker
}

func NewRepeater(opts ...repeaterOption) *Repeater {
	r := &Repeater{}

	r.setDefaults()

	for _, opt := range opts {
		opt(r)
	}

	return r
}

func (r *Repeater) setDefaults() {
	r.interval = 0
	r.retries = math.MaxInt - 1
	r.checkResult = IsNoError
}

// Run runs the f() until tries exceed or f() "succeeds".
// Repeater.checkResult tells whether or not f() result is successful
//
// Returns a number of tries and an error:
// - nil on first successful f() call
// - errRetriesExceeded if number of tries reaches the limit.
func (r *Repeater) Run(f func() error) (int, error) {
	tries := 0

	for tries <= r.retries {
		err := f()
		tries++
		if r.checkResult(err) {
			return tries, nil
		}

		time.Sleep(r.interval)
	}

	return tries, errRetriesExceeded
}

// RunStats runs the f() retries times.
// Returns number of total tries and successful tries.
func (r *Repeater) RunStats(f func() error) (int, int) {
	totalTries := 0
	successfulTries := 0

	for totalTries <= r.retries {
		err := f()
		totalTries++
		if r.checkResult(err) {
			successfulTries++
		}

		time.Sleep(r.interval)
	}

	return totalTries, successfulTries
}
