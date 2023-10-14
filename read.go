package main

import (
	"fmt"
	"log"
	"os"
)

func MainRead() {

	file, err := os.Open("read_file/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	var fileSlice []string

	for i := 0; i < 5; i++ {
		x := 43
		fileData := make([]byte, x)
		count, _ := file.ReadAt(fileData, int64(x*i))
		fileSlice = append(fileSlice, string(fileData[:count]))
	}
	fmt.Print(fileSlice)
}
