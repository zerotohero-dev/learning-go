package main

// use with reverb.go

import (
    "net"
    "os"
    "log"
    "io"
)

func main() {
    conn, err := net.Dial("tcp", "0.0.0.0:8080")

    if err != nil {
        log.Fatal(err)
    }

    defer conn.Close()

    done := make(chan struct{})

    go func() {
        io.Copy(os.Stdout, conn)
        log.Println("Done")
        done <- struct{}{}
    }()

    mustCopy(conn, os.Stdin)

    <-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}