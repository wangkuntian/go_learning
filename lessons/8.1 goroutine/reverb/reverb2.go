package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo2(c net.Conn, shout string, deley time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(deley)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(deley)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn2(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo2(c, input.Text(), time.Second)
	}
	c.Close()
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn2(conn) //并发处理连接
	}
}
