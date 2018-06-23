package main

import (
    "bufio"
    "net"
    "os"
    "fmt"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
        os.Exit(1)
    }
    service := os.Args[1]

    sendUDPRequest(service)
    sendTCPRequest(service)

    os.Exit(0)
}

func input(requestType string) string {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Enter the " + requestType + " request body: ")
    input, err := reader.ReadString('\n')
    checkError(err)
    return input
}

func sendUDPRequest(service string) {
    udpAddr, err := net.ResolveUDPAddr("", service)
    checkError(err)

    conn, err := net.DialUDP("udp", nil, udpAddr)
    checkError(err)

    _, err = conn.Write([]byte(input("UDP")))
    checkError(err)

    var buf [512]byte
    n, err := conn.Read(buf[0:])
    checkError(err)

    fmt.Println("Response from UDP Server: " + string(buf[0:n]))
}

func sendTCPRequest(service string) {
    tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
    checkError(err)

    conn, err := net.DialTCP("tcp", nil, tcpAddr)
    checkError(err)

    _, err = conn.Write([]byte(input("TCP")))
    checkError(err)

    var buf [512]byte
    n, err := conn.Read(buf[0:])
    checkError(err)

    fmt.Println("Response from TCP Server: " + string(buf[0:n]))
}

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
        os.Exit(1)
    }
}
