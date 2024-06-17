package repeater

import "time"

type repeaterOption func(r *Repeater)

// WithInterval sets interval (in seconds) between attempts.
// Defaults to 0.
func WithInterval(interval time.Duration) repeaterOption {
	return func(r *Repeater) {
		r.interval = interval
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
