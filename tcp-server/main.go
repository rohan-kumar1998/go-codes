package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func serve(conn net.Conn) {
	scanner := bufio.NewScanner(conn)

	defer conn.Close()

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(string(ln))
	}
}

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Panic(err)
		}
		go serve(conn)
	}
}
