package main

import (
	"fmt"
	"main/Handle"
	"net"
	"time"
)

func main() {

	dial, err := net.Dial("tcp", "localhost:5678")
	Handle.HandleError(err)
	defer dial.Close()

	data := []byte("Hello mapren!")
	_, err = dial.Write(data)
	Handle.HandleError(err)

	dial.SetWriteDeadline(time.Now().Add(5 * time.Second))

	if err != nil {
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			fmt.Println("Time Out!")
			return
		} else {
			Handle.HandleError(err)
		}
	}

}
