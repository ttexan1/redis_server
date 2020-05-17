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

	sendMessage(connection, "")
}

func sendMessage(connection net.Conn, text string) {
	fmt.Print("> ")

	stdin := bufio.NewScanner(os.Stdin)
	if stdin.Scan() == false {
		fmt.Println("-Error")
		return
	}
	converted := simpleConvert(text)
	fmt.Println(converted)

	_, err := connection.Write([]byte(converted))

	if err != nil {
		panic(err)
	}

	var response = make([]byte, 4*1024)
	_, err = connection.Read(response)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Server> %s \n", response)
	text = ""
	sendMessage(connection, text)
}

func simpleConvert(text string) string {
	if text == "" {
		return "*1\r\n$4\r\nping\r\n"
	}
	if strings.ToLower(text) == "ping" {
		return "*1\r\n$4\r\nping\r\n"
	}
	if strings.ToLower(text) == "echo test" {
		return "*2\r\n$4\r\necho\r\n$4\r\ntest\r\n"
	}

	args := strings.Split(text, " ")
	if args[0] == "get" {
		if strings.ToLower(text) == "get key" {
			length := len(args[1])
			key := args[1]
			return fmt.Sprintf("*2\r\n$3\r\nget\r\n$%d\r\n%s\r\n", length, key)
		}
	}
	// set key value
	keyLen := len(args[1])
	key := args[1]
	valLen := len(args[2])
	value := args[2]
	return fmt.Sprintf("*3\r\n$3\r\nset\r\n$%d\r\n%s\r\n$%d\r\n%s\r\n", keyLen, key, valLen, value)
}
