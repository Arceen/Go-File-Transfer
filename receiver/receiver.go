package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	port := flag.String("port", "8080", "Port to connect to")
	flag.Parse()

	conn, err := net.Dial("tcp", "localhost:"+*port)
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	// Receive filename first
	filenameBuffer := make([]byte, 256) // Adjust the buffer size accordingly
	n, err := conn.Read(filenameBuffer)
	if err != nil {
		fmt.Println("Error receiving filename:", err)
		return
	}
	filename := string(filenameBuffer[:n])

	// Create a file with the received filename
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Receive the file content
	_, err = io.Copy(file, conn)
	if err != nil {
		fmt.Println("Error receiving file:", err)
		return
	}

	fmt.Println("File received and saved as", filename)
}
