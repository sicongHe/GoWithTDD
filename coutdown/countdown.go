package coutdown

import (
	"fmt"
	"io"

	"time"
)

var finalWord = "GO!"
var count = 3

type SpyCountdownWriter struct {
	record string
}

type SpySleeper struct {
	Calls int
}
type ConfigurableSleeper struct {
	Duration time.Duration
}

type Sleeper interface {
	Sleep()
}

func (s *SpySleeper) Sleep() {
	s.Calls += 1
}
func (c *ConfigurableSleeper) Sleep() {
	time.Sleep(c.Duration)
}
func (s *SpyCountdownWriter) Sleep() {
	s.record += "sleep\n"
}

func (s *SpyCountdownWriter) Write(p []byte) (n int, err error) {
	s.record += "write\n"
	return
}

func Countdown(w io.Writer, s Sleeper) {
	for i := count; i > 0; i-- {
		s.Sleep()
		fmt.Fprintln(w, i)
	}
	s.Sleep()
	fmt.Fprint(w, finalWord)
}
