package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

	"redis_app/command"
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
		fmt.Println("Ciao ciao!")
		return
	}
	newText := stdin.Text()
	if strings.HasSuffix(newText, "/") {
		text += "\n" + newText[:len(newText)-1] + "\n"
	} else {
		text += "\n" + newText
		sendMessage(connection, text)
		return
	}
	text = strings.Replace(text, "\n", "", 1)
	fmt.Print(text)

	_, err := connection.Write([]byte(text))

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

// convertRedisProtocol convert the raw string command to redis protocol string
func convertRedisProtocol(text string) string {
	args := strings.Split(text, " ")
	switch strings.ToLower(args[0]) {
	case command.Ping:
		// fmt.Println(args)
		return dealPing(args)
	default:
		return "BBB"
	}
}

func dealPing(args []string) string {
	if len(args) == 1 {
		return `
*1
$4
ping
`
	}
	if len(args) == 2 {
		return fmt.Sprintf(`
*2
$4
ping
$%d
%s
`, len(args[1]), args[1])
	}
	return InvalidArgLength()
}

func InvalidArgLength() string {
	return "wrong command"
}
