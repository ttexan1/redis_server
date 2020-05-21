package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	connection, err := net.Dial("tcp", "localhost:10000")
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	sendMessage(connection)
}

func sendMessage(connection net.Conn) {
	// fmt.Print("> ")

	stdin := bufio.NewScanner(os.Stdin)
	if stdin.Scan() == false {
		return
	}
	text := stdin.Text()

	converted := simpleConvert(text)
	fmt.Print(converted)

	_, err := connection.Write([]byte(converted))

	if err != nil {
		panic(err)
	}

	var response = make([]byte, 4*1024)
	_, err = connection.Read(response)
	if err != nil {
		panic(err)
	}

	fmt.Printf(">> %s", response)
	text = ""
	sendMessage(connection)
}

func simpleConvert(text string) string {
	args := strings.Split(text, " ")
	directive := strings.ToLower(args[0])
	if directive == "ping" {
		return "*1\r\n$4\r\nping\r\n"
	}
	if directive == "echo" {
		return "*2\r\n$4\r\necho\r\n$4\r\ntest\r\n"
	}
	if directive == "get" {
		length := len(args[1])
		key := args[1]
		return fmt.Sprintf("*2\r\n$3\r\nget\r\n$%d\r\n%s\r\n", length, key)
	}
	if directive == "set" {
		keyLen := len(args[1])
		key := args[1]
		valLen := len(args[2])
		value := args[2]
		return fmt.Sprintf("*3\r\n$3\r\nset\r\n$%d\r\n%s\r\n$%d\r\n%s\r\n", keyLen, key, valLen, value)
	}
	return "*1\r\n$4\r\nping\r\n"
}
