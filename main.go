package main

import (
	"fmt"
	"os"
)

func main() {
	jsonFile, err := os.Open("known_ports.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully opened known_ports.json")

	defer jsonFile.Close()

	type Ports struct {
		Ports []Port `json:"ports"`
	}

	type Port struct {
		Number int
	}
	// fmt.Println("Enter the port number you wish to check:")
	// var port_num int

	// fmt.Scanln(&port_num)

	// fmt.Printf("Your port number is: %d", port_num)
}
