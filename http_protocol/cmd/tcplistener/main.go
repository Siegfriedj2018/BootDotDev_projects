package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)


const port string = ":42069"
const protocol string = "tcp"

func getLinesChannel(f io.ReadCloser) <-chan string {
	str := make(chan string)

	go func() {
		defer f.Close()
		defer close(str)
		var strBuffer string
		for {	
			buffer := make([]byte, 8)
			var parts []string
			
			redData, err := f.Read(buffer)
			lineBuffer := string(buffer[:redData])
			parts = strings.Split(lineBuffer, "\n")
			strBuffer += parts[0]
			if len(parts) > 1 {
				str <-strBuffer
				strBuffer = ""
				strBuffer += parts[len(parts)-1]
			}
			if err != nil {
				if err == io.EOF {
					strBuffer += parts[0]
					str <-strBuffer
					
					fmt.Println("Finished reading file, Goodbye!")
					break
				}
				
				fmt.Printf("Error reading the file: %v\n", err)
				return
			}
		}
	}()
	return str
}


func main() {
	currentFile, err := net.Listen(protocol, port)
	if err != nil {
		log.Fatalf("Error opening: %v\n", err)
	}
	defer currentFile.Close()

	fmt.Printf("Listening for TCP traffic: %s\n", currentFile.Addr())
	fmt.Println("=======================================")

	for {
		conn, err := currentFile.Accept()
		if err != nil {
			fmt.Printf("Error accecpting connections: %v\n", err)
			currentFile.Close()
			os.Exit(2)
		}
		fmt.Println("Connection accepted from", conn.RemoteAddr())

		stringChan := getLinesChannel(conn)
		
		for line := range stringChan {
			fmt.Println(line)
		}
		fmt.Println("Connection to ", conn.RemoteAddr(), "closed")
	}
}
