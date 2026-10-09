package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	fmt.Println("servidor iniciado, aguardando conexões")

	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Printf("problema ao estabelecer conexão %v", err)
			continue
		}
		go handleConnection(connection)
	}
}

func handleConnection(connection net.Conn) {
	defer connection.Close()
	fmt.Println("conexão estabelecida")
}