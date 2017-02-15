package main

import (
    "io"
    "log"
    "net"
    "time"
)

func main() {
    listener, err := net.Listen("tcp", "0.0.0.0:8080")

    if err != nil {
        log.Fatal(err)
    }

    for {

        // `listener.Accept()` Blocks until a new connection:
        conn, err := listener.Accept()

        if err != nil {
            log.Print(err)

            continue
        }

        //handleConn(conn)

        go handleConn(conn)
    }
}

func handleConn(c net.Conn) {
    defer c.Close()

    for {
        _, err := io.WriteString(c, time.Now().Format("15:04:05\n"))

        if err != nil {
            return
        }

        time.Sleep(1 * time.Second)
    }
}