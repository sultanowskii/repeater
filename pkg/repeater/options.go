package repeater

import "time"

type repeaterOption func(r *Repeater)

// WithInterval sets interval (in seconds) between tries.
// Defaults to 0.
func WithSecondsInterval(n int) repeaterOption {
	return func(r *Repeater) {
		r.interval = time.Duration(n) * time.Second
	}
}

// WithInterval sets interval (in milliseconds) between tries.
// Defaults to 0.
func WithMillisecondsInterval(n int) repeaterOption {
	return func(r *Repeater) {
		r.interval = time.Duration(n) * time.Millisecond
	}
}

// WithRetries sets number of retries.
// Defaults to infinite (?)
func WithRetries(n int) repeaterOption {
	return func(r *Repeater) {
		r.retries = n
	}
}

// WithResultChecker sets result checker that tells whether
// function call was successful or not.
// Defaults to repeater.IsNoError
func WithResultChecker(checker resultChecker) repeaterOption {
	return func(r *Repeater) {
		r.checkResult = checker
	}
}
