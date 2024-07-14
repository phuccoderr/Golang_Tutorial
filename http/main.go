package main

import (
	"fmt"
	"io"
	"os"
)

type logWriter struct{}

func main() {
	// resp, err := http.Get("https://google.com")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	os.Exit(1)
	// }
	
	// lw := logWriter{}
	// io.Copy(lw, resp.Body)

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout,f)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs),nil
}