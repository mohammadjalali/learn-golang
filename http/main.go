package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type Reader interface {
	Read(p []byte) (n int, err error)
}

type logWriter struct{}

func main() {
	res, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lw := logWriter{}
	io.Copy(lw, res.Body)
	// body, err := io.ReadAll(resp.Body)
	// resp.Body.Close()
	// if resp.StatusCode > 299 {
	// 	log.Fatalf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)
	// }
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s", body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}

func (logWriter) ReaderFrom(i int) (int64, error) {
	fmt.Println("*******************************************************")
	return 2, nil
}
