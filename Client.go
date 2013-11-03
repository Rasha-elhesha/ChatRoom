package main

import (
    "net"
	"time"
)

func main() {
    c,err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        panic("Error")
    }
    defer c.Close()
    _,err := c.Write([]byte("hi, I am the client"))
    if err != nil {
        println("Error write")
    }
    time.Sleep(10)
}
