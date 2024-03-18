package main

import (
	"fmt"
	"net"
	"os"
	"bufio"
	"strings"
)

func main() {
	args := os.Args
	if len(args) != 4 {
		return
	}

	name := args[1]
	address := args[2]
	port := args[3]

	var server, err = net.Dial("tcp", address + ":" + port)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Failed to connect - Dial error")
	// 	return
	// }
	for err != nil {
		server, err = net.Dial("tcp", address + ":" + port)
	}
	fmt.Fprintf(server, name + " ")

	for {
		event := bufio.NewReader(os.Stdin)
		bytes_read, _ := event.ReadString('\n')
		// fmt.Fprintf(os.Stdout, bytes_read)
		res := strings.Split(bytes_read, " ")
		time := res[0]
		message := res[1]
		fmt.Fprintf(server, time + " " + name + " " + message)
	}
}
