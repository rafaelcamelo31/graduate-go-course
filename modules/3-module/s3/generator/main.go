package main

import (
	"fmt"
	"os"
)

func main() {
	i := 0

	for i < 1000 {
		dir, err := os.Create(fmt.Sprintf("./tmp/file%d.txt", i))
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer dir.Close()

		_, err = dir.WriteString(fmt.Sprintf("This is file number %d\n", i))
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
		i++
	}
}
