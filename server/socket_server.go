package main

import (
    "bytes"
    "fmt"
    "net"
    "os"
    "sync"
)

func main() {
	if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
        os.Exit(1)
    }
    service := os.Args[1]
    
    var wg sync.WaitGroup
    wg.Add(2)

    go startTCPServer(service)
    go startUDPServer(service)

    wg.Wait()
}

func startUDPServer(service string) {
	udpAddr, err := net.ResolveUDPAddr("", service)
    checkError(err)

    conn, err := net.ListenUDP("udp", udpAddr)
    checkError(err)

    for {
        handleClientUDP(conn)
    }
}

func handleClientUDP(conn *net.UDPConn) {
	var buf [512]byte
    n, addr, err := conn.ReadFromUDP(buf[0:])
    if err != nil {
        return
    }
    
    fmt.Println("Received UDP Request is: " + string(buf[0:]))
    
    _, err2 := conn.WriteToUDP(bytes.ToUpper(buf[0:n]), addr)
    if err2 != nil {
        return
    }
}

func startTCPServer(service string) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
    checkError(err)

    listener, err := net.ListenTCP("tcp", tcpAddr)
    checkError(err)

    for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }
        handleClientTCP(conn)
        conn.Close()
    }
}

func handleClientTCP(conn net.Conn) {
    var buf [512]byte
    for {
        n, err := conn.Read(buf[0:])
        if err != nil {
            return
        }
    
        fmt.Println("Received TCP Request is: " + string(buf[0:]))
    
        _, err2 := conn.Write(bytes.ToUpper(buf[0:n]))
        if err2 != nil {
            return
        }
    }
}

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}
