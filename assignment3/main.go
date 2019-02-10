package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	filename := os.Args[1]
	byteSlice, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error occured while reading file:", err)
		os.Exit(1)
	}
	fmt.Println("Contents read from", filename+":")
	fmt.Println(string(byteSlice))
}
