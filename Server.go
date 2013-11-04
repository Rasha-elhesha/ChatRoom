package main
 
import (
    "net"
    "bufio"
    "fmt"
	"strconv"
)

// define new structure corresponding to each session
type session struct {
    // Registered connections.
    connections [] net.Conn
	// Corresponding names of registered connections
	names [] string
    // The session administrator
	administrator string
}
 
const PORT = 8080

// define new structure corresponding to each session
type server struct {
	currentSession session
}

var newSession session

func main() {
    server, _ := net.Listen("tcp", ":" + strconv.Itoa(PORT))
    if server == nil {
        panic("couldn't start listening....")
    }
    conns := clientConns(server)
    for {
        go handleConnection(<-conns)
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
            channel <- client
        }
    }()
    return channel
}
 
func handleConnection(client net.Conn) {
    b := bufio.NewReader(client)
	//receive the user name
	buff := make([]byte, 512)
	clientNameb, _ := client.Read(buff)
	clientName := string(buff[0:clientNameb])
	
	//check if the user wants to start new session or join existing
	nr, _ := client.Read(buff)
	messageFromClient := string(buff[0:nr][0])
	clientOption, _:= strconv.Atoi(messageFromClient)
	if clientOption == 1 {
		//start a new session
		fmt.Println(clientName + ", you choose to Start a new session")	
		newSession = session {
			connections: []net.Conn{client},
			names : []string{clientName},
			administrator:   clientName,
		}
	} else {
		//join an existing session
		fmt.Println(clientName + ", you choose to Join an existing session")
		newSession.names = append(newSession.names, clientName)
		newSession.connections = append(newSession.connections, client)
	}
    for {
        line, err := b.ReadBytes('\n')
        if err != nil { // EOF, or worse
            break
        }
        client.Write(line)
    }
}
