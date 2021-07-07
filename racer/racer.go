package racer

import (
	"errors"
	"net/http"
	"time"
)

func Race(a, b string, timeout time.Duration) (winner string, err error) {

	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", errors.New("请求超时")
	}
}
func ping(url string) chan bool {
	ret := make(chan bool)
	go func(u string) {
		http.Get(url)
		ret <- true
	}(url)
	return ret

}
