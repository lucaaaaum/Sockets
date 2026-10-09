package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("cliente iniciado, tentando conectar")
	connection, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Printf("problema ao estabelecer conexão %v\n", err)
	}
	handleConnection(connection)
}

func handleConnection(connection net.Conn) {
	defer connection.Close()
	fmt.Println("conexão estabelecida")

	connectionReader := bufio.NewReader(connection)
	terminalReader := bufio.NewReader(os.Stdin)
	
	fmt.Print("digite seu usuário: ")
	user, err := terminalReader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(connection, "LOGIN %s\n", user)
	loginAnswer, err := connectionReader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	fmt.Println(loginAnswer)
}