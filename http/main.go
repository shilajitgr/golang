package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}

	// bs := make([]byte, 99999) // empty byte slice with a capacity of 99999
	// resp.Body.Read(bs)	:	resp.Body is a io.ReadCloser type object, which is calling Read function defined
	// within Reader interface, which itself is an honorary member of "io.ReadCloser" interface
	// fmt.Println(string(bs))

	io.Copy(logWriter{}, resp.Body) // user defined Writer object but will not write anything on the cmdline
	// as the Write func is not writing anything to the cmdline
	// func Copy(dst Writer, src Reader) (written int64, err error) 	:	its converting the resp.Body into slice
	// the writing it to the stdout
}

// Old implementation
// func (logWriter) Write(bs []byte) (int, error) {
// 	return 1, nil
// }

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("just wrote", len(bs), "bytes")

	return len(bs), nil
}
