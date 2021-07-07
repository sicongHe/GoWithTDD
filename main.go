package main

import (
	"os"
	"time"

	"github.com/siconghe/hello/coutdown"
)

func main() {
	sleeper := &coutdown.ConfigurableSleeper{Duration: 1 * time.Second}
	coutdown.Countdown(os.Stdout, sleeper)
}
