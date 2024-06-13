package main

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/sultanowskii/repeater/pkg/repeater"
)

func importantFunction() error {
	if rand.Int()%2 == 0 {
		return errors.New("oh no")
	} else {
		return nil
	}
}

func main() {
	r := repeater.NewRepeater(
		repeater.WithMillisecondsInterval(100),
		repeater.WithRetries(5),
	)

	q, err := r.Run(importantFunction)
	if err == nil {
		fmt.Printf("successful run within %d attempt(s)\n", q)
	} else {
		fmt.Printf("'%v' after %d attempt(s)\n", err, q)
	}

	total, success := r.RunStats(importantFunction)
	fmt.Printf("%d out of %d attempt(s) were successful\n", success, total)
}
