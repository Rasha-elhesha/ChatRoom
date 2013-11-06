package main

import (
    "net"
	"fmt"
	"bufio"
	"os"
)



func main() {
    connection,err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        panic("Error")
    }
    defer connection.Close()
	//send the user name
	fmt.Printf("Welcome to chat rooms. Please enter your name.\n")
	inputReader := bufio.NewReader(os.Stdin)
	name, _ , _:= inputReader.ReadLine()
	userName := string(name)
    connection.Write([]byte(userName))
	
	fmt.Printf("*****************Go Chat*****************\n")
	//run the writer to read messages from console and send them to the server
	go messageWriter(connection)
	
    //run the reader to read any messages received from the server
	messageReader(connection)
	
}

func messageReader(client net.Conn) {
	inputReader := bufio.NewReader(os.Stdin)
	for {
		message, _ := inputReader.ReadString('\n')
		client.Write([]byte(message))
	}
}

func messageWriter(client net.Conn) {
	reader := bufio.NewReader(client)
	for {
		message, err := reader.ReadBytes('\n')
        if err != nil { // EOF, or worse
		fmt.Printf("Error read from server")
            break
        }
		//write the received message
		fmt.Printf(string(message))
	}
}
