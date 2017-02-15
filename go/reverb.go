package main

import (
    "net"
    "log"
    "time"
    "fmt"
    "strings"
    "bufio"
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

func echo(c net.Conn, shout string, delay time.Duration) {
    fmt.Fprintln(c, "\t", strings.ToUpper(shout))
    time.Sleep(delay)
    fmt.Fprintln(c, "\t", shout)
    time.Sleep(delay)
    fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
    input := bufio.NewScanner(c)

    for input.Scan() {

        //echo(c, input.Text(), 1*time.Second)
        go echo(c, input.Text(), 1*time.Second)
    }

    c.Close()
}

