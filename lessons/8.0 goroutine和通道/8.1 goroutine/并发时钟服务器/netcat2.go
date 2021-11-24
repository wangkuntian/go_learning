package main

import (
	"io"
	"log"
	"net"
	"os"
)

func mustCopy2(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go mustCopy2(os.Stdout, conn)
	mustCopy2(conn, os.Stdin)
}
