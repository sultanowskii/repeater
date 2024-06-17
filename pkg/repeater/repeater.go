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

// Run runs the f() until attempts exceed or f() succeeds.
// Repeater.checkResult tells whether or not f() result is successful
//
// Returns a number of attempts and an error:
// - nil on first successful f() call
// - errRetriesExceeded if number of attempts reaches the limit.
func (r *Repeater) Run(f func() error) (int, error) {
	attempts := 0

	for attempts <= r.retries {
		err := f()
		attempts++
		if r.checkResult(err) {
			return attempts, nil
		}

		time.Sleep(r.interval)
	}

	return attempts, errRetriesExceeded
}

// RunStats runs the f() retries times.
// Returns number of total and successful attempts.
func (r *Repeater) RunStats(f func() error) (int, int) {
	totalAttempts := 0
	successfulAttempts := 0

	for totalAttempts <= r.retries {
		err := f()
		totalAttempts++
		if r.checkResult(err) {
			successfulAttempts++
		}

		time.Sleep(r.interval)
	}

	return totalAttempts, successfulAttempts
}
