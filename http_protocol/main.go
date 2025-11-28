package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)
func main() {
	currentFile, err := os.Open("./messages.txt")
	if err != nil {
		log.Fatalf("Error opening file: %v\n", err)
	}

	buffer := make([]byte, 8)
	var lineBuffer string

	readData, err := currentFile.Read(buffer)
	if err != nil {
		log.Fatalf("Error reading the file\n")
	}
	
	for readData > 0 {
		parts := strings.Split(string(buffer), "\n")
		lineBuffer = lineBuffer + parts[0]
		if len(parts) > 1 {
			fmt.Printf("All Parts: %+v\n", parts)
			// fmt.Printf("read: %s\n", lineBuffer)
			lineBuffer = ""
			if parts[1] == "" {
				lineBuffer += ""
			} else {
				lineBuffer += parts[1]
			}
		}
		
		readData, err = currentFile.Read(buffer)
		if err != nil {
			if err == io.EOF {
				log.Println("Finished reading file, Goodbye!")
				os.Exit(0)
			}

			log.Fatalf("Error reading the file: %v\n", err)
		}
		
	}
	
	

}