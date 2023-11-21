package main

import "flag"

func main() {
	flag.String("p", "8005", "Enter the Port Number")
	flag.String("f", "", "File Path for file send")
}
