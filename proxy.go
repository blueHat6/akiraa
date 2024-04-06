package main

import (
	"net"
	"io"
	"main/Handle"
)

func proxyForward(from io.Reader, to io.Writer) error {
	_, err := io.Copy(to, from)

	return err
}

func main(){
	
	listener, err := net.Listen("tcp", "localhost:5678")
	Handle.HandleError(err)
	defer listener.Close()

	for{
		conn, err := listener.Accept()
		Handle.HandleError(err)

		go handleServer(conn)
	}
}

func handleServer(from net.Conn){
	defer from.Close()

	to, err := net.Dial("tcp", "localhost:1234")
	Handle.HandleError(err)
	defer to.Close()

	err = proxyForward(from, to)
	Handle.HandleError(err)

}