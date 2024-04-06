package main

import (
	"fmt"
	"main/Handle"
	"net"
	"time"
)

func main() {

	listener, err := net.Listen("tcp", "localhost:1234")
	Handle.HandleError(err)
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		Handle.HandleError(err)

		go handleClient(conn)

		conn.SetReadDeadline(time.Now().Add(5 * time.Second))

		if err != nil{
			if netErr, ok := err.(net.Error); ok && netErr.Timeout(){
		fmt.Println("Time Out!")
		return
		} else{
			Handle.HandleError(err)
		}
	}
	}

}

func handleClient(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer)
	Handle.HandleError(err)

	fmt.Printf("Received: %s\n", buffer[:n])
}
