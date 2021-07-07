package di

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	ret := "hello " + name
	fmt.Fprint(writer, ret)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, r.URL.String())
}
