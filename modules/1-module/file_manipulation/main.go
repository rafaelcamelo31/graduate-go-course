package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("sample.txt")
	if err != nil {
		panic(err)
	}

	size, err := f.Write([]byte("I love basketball!\n"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Successfully created a file! Size: %d bytes \n", size)
	f.Close()

	file, err := os.ReadFile(("sample.txt"))
	if err != nil {
		panic(err)
	}
	fmt.Println(string(file))

	// Read a file from chunks of buffers
	// This is for when there are no enough memories to read whole file at once
	sampleFile, err := os.Open("sample.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(sampleFile)
	buffer := make([]byte, 10)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	err = os.Remove("sample.txt")
	if err != nil {
		panic(err)
	}
}
