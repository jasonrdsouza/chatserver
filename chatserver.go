/*
    Simple chat-roulette type server implemented in Golang.
    Supports tcp and websocket protocols
    Source: http://talks.golang.org/2012/chat.slide
*/
package chatserver

import (
    "io"
    "log"
    "fmt"
)

const listenAddr = "localhost:4000"

// Simple echo functionality
func echo(c io.ReadWriteCloser) {
    io.Copy(c, c)
}

// This channel is used to pair up chatters
var partner = make(chan io.ReadWriteCloser)

// Match function simultaneously tries to send and
// receive a connection on a channel. The select 
// statement allows only one case to succeed
func match(c io.ReadWriteCloser) {
    fmt.Fprint(c, "Waiting for a partner...")
    select {
    case partner <- c:
        // now handled by the other goroutine
    case p := <-partner:
        chat(p, c)
    }
}

// Allow for concurrent chatting, and keep track of
// errors/ chat members leaving
func chat(a, b io.ReadWriteCloser) {
    fmt.Fprintln(a, "Found one! Say hi.")
    fmt.Fprintln(b, "Found one! Say hi.")
    errc := make(chan error, 1)
    go cp(a, b, errc)
    go cp(b, a, errc)
    if err := <-errc; err != nil {
        log.Println(err)
    }
    a.Close()
    b.Close()
}

func cp(w io.Writer, r io.Reader, errc chan<- error) {
    _, err := io.Copy(w, r)
    errc <- err
}