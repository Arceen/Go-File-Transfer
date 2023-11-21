package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
)

func main() {
	port := flag.String("port", "8080", "Port to listen on")
	filePath := flag.String("file", "", "File path to for file to be send")
	flag.Parse()

	if *filePath == "" {
		fmt.Println("Please provide a file path using -file flag.")
		return
	}

	ln, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer ln.Close()

	fmt.Println("Listening on port", *port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// Send filename first
		_, err = conn.Write([]byte(filepath.Base(*filePath)))
		if err != nil {
			fmt.Println("Error sending filename:", err)
			conn.Close()
			continue
		}

		// Then send the file
		file, err := os.Open(*filePath)
		if err != nil {
			fmt.Println("Error opening file:", err)
			conn.Close()
			continue
		}
		defer file.Close()

		_, err = io.Copy(conn, file)
		if err != nil {
			fmt.Println("Error sending file:", err)
			conn.Close()
			continue
		}

		fmt.Println("File sent to", conn.RemoteAddr())
	}
}
