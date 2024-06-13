package repeater

import "errors"

var (
	errRetriesExceeded = errors.New("retries exceeded")
)
