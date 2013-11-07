package main

import (
    "net"
	"fmt"
	"bufio"
	"os"
)



func main() {
    connection,err := net.Dial("tcp", "localhost:8080")
	//Close connection and clean up when the client finish chat
    if err != nil {
        fmt.Println("Can not connect to the server")
		return
    }
	defer cleanUp(connection)
	//send the user name to the server
	sendUserName(connection)
	
	fmt.Printf("*****************Go Chat*****************\n")
	//run the writer to read messages from console and send them to the server
	go messageWriter(connection)
	
    //run the reader to read any messages received from the server
	messageReader(connection)
	
}

/*
 * This function is called when error occurred during runtime
 * The function closes the connection made with the server
*/
func cleanUp(clientConnection net.Conn) {
	clientConnection.Close()
	os.Exit(0)
}

/*
 * This function is called by the main after the connection made to the server
 * The function request the user to enter his name and sends this name to the server
*/
func sendUserName(client net.Conn) {
	fmt.Printf("Welcome to chat rooms. Please enter your name.\n")
	inputReader := bufio.NewReader(os.Stdin)
	name, _ , error:= inputReader.ReadLine()
	if error != nil {
        fmt.Println("Can not read user name")
		cleanUp(client)
		return
    }
	userName := string(name)
    client.Write([]byte(userName))
}

/*
 * This function is called by the main as part of chatting functionality
 * The function waits for user input message  and sends this message to the server to broadcast it
*/
func messageReader(client net.Conn) {
	inputReader := bufio.NewReader(os.Stdin)
	for {
		message, error := inputReader.ReadString('\n')
		if error != nil {
			fmt.Println("Can not read user message")
			cleanUp(client)
			break
		}
		//Send the user message to the server
		client.Write([]byte(message))
	}
}

/*
 * This function is called by the main as part of chatting functionality
 * The function waits the server to send a broadcasted message and print this message in the terminal
*/
func messageWriter(client net.Conn) {
	reader := bufio.NewReader(client)
	for {
		message, err := reader.ReadBytes('\n')
        if err != nil {
			fmt.Println("Error read from server")
			cleanUp(client)
            break
        }
		//write the received message
		fmt.Printf(string(message))
	}
}
