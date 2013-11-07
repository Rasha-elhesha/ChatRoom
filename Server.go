package main
 
import (
    "net"
    "bufio"
    "fmt"
	"strconv"
)

//Constants
const PORT = 8080

// define new structure corresponding to each session
type session struct {
    // Registered connections.
    connections [] net.Conn
	// Corresponding names of registered connections
	names [] string
}
 

// define new structure corresponding to each session
type server struct {
	currentSession session
}
//global variables
var newSession session

func main() {
    server, _ := net.Listen("tcp", ":" + strconv.Itoa(PORT))
    if server == nil {
        panic("couldn't start listening....")
    }
	newSession = session {
			connections: []net.Conn{},
			names : []string{},
			}
    conns := clientConns(server)
    for {
        go handleConnection(<-conns)
    }
}

/*
 * This function is called by the main
 * The function listen for client connections, accept them
*/
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
 
/*
 * This function is called by the main to handle the chat functionality of new connection
 * The function reads the user name and saves it, then waits for user messages and broadcast them
*/
func handleConnection(client net.Conn) {
    reader := bufio.NewReader(client)
	//receive the user name
	buff := make([]byte, 512)
	clientNameb, _ := client.Read(buff)
	clientName := string(buff[0:clientNameb])
	newSession.names = append(newSession.names, clientName)
	newSession.connections = append(newSession.connections, client)
	
    for {
		//receive the user message
        line, err := reader.ReadBytes('\n')
        if err != nil {
            break
        }
		//broadcast client message
		message := clientName + ":" + string(line)
		for _, currentClient := range newSession.connections {
			if(currentClient != nil) {
				currentClient.Write([]byte(message))
			}
		}
        
    }
}
