package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	io.Copy(logWriter{}, resp.Body)
}

func (l logWriter) Write(bs []byte) (int, error) {
	fmt.Println("In logWriter.Write")
	fmt.Println(string(bs))
	fmt.Println("End of logWrite.Write")
	return len(bs), nil
}
