package main
 
import (
    "net"
    "bufio"
    "fmt"
	"strconv"
)
 
const PORT = 8080
 
func main() {
    server, _ := net.Listen("tcp", ":" + strconv.Itoa(PORT))
    if server == nil {
        panic("couldn't start listening....")
    }
    conns := clientConns(server)
    for {
        go handleConn(<-conns)
    }
}
 
func clientConns(listener net.Listener) chan net.Conn {
    channel := make(chan net.Conn)
    go func() {
        for {
            client, _ := listener.Accept()
            if client == nil {
                fmt.Printf("couldn't accept client connection")
                continue
            }
		buff := make([]byte, 512)
		nr, err := client.Read(buff)
        if err != nil {
            return
        }

        messageFromClient := buff[0:nr]
        fmt.Printf("Received: %v", string(messageFromClient))
            channel <- client
        }
    }()
    return channel
}
 
func handleConn(client net.Conn) {
    b := bufio.NewReader(client)
    for {
        line, err := b.ReadBytes('\n')
        if err != nil { // EOF, or worse
            break
        }
        client.Write(line)
    }
}
