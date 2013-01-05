package chatserver

import (
    "log"
    "net"
)

func Tcp_main() {
    l, err := net.Listen("tcp", listenAddr)
    if err != nil {
        log.Fatal(err)
    }
    for {
        c, err := l.Accept()
        if err != nil {
            log.Fatal(err)
        }
        //go echo(c)
        go match(c)
    }
}