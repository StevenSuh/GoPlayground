package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		log.Fatal("Error:", err)
	}

	lw := logWriter{}

	io.Copy(lw, resp.Body)

	// data, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal("Error:", err)
	// }

	// fmt.Println(string(data))
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
