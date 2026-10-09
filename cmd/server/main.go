package main

import (
	"bufio"
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
			fmt.Printf("problema ao estabelecer conexão %v\n", err)
			continue
		}
		go handleConnection(connection)
	}
}

func handleConnection(connection net.Conn) {
	defer closeConnection(connection)
	fmt.Println("conexão estabelecida")

	scanner := bufio.NewScanner(connection)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			break
		}
		line := scanner.Text()
		fmt.Printf("Recebida linha:\n%s\n", line)
	}
}

func closeConnection(connection net.Conn) {
	connection.Close()
	fmt.Println("conexão encerrada")
}