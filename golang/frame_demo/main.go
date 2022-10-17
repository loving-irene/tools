package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func main() {

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnection(conn)
	}
	fmt.Println("test")
}

func handleConnection(conn net.Conn) {
	br := bufio.NewReader(conn)
	for {
		data, err := br.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Printf("%s", data)
		fmt.Fprintf(conn, "OK\n")
	}
}
