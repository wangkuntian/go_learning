package main

import (
	"io"
	"log"
	"net"
	"os"
)

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8999")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("Done")
		done <- struct{}{}
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done
}
