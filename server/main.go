package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"

	"redis_app/domain"
	"redis_app/parser"
	"redis_app/store"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:10000")
	if err != nil {
		panic(err)
	}
	db := map[string]map[string]domain.Single{}

	log.Println("Server running at localhost:10000")

	waitClient(listener, db)
}

func waitClient(listener net.Listener, db map[string]map[string]domain.Single) {
	connection, err := listener.Accept()
	if err != nil {
		panic(err)
	}
	sessionID := rand.Intn(1 << 20)
	db[string(sessionID)] = map[string]domain.Single{}

	go goEcho(connection, db[string(sessionID)])
	waitClient(listener, db)
}

func goEcho(connection net.Conn, d map[string]domain.Single) {
	fmt.Println(connection.LocalAddr())
	dddd := &store.DB{St: make(map[string]*domain.Single)}
	defer func() {
		fmt.Println("Closed")
		connection.Close()
	}()

	echo(connection, dddd)
}

func echo(connection net.Conn, d *store.DB) {
	var buf = make([]byte, 1024)

	_, err := connection.Read(buf)
	if err != nil {
		if err == io.EOF {
			return
		}
		panic(err)
	}
	fmt.Println(">", string(buf))

	resFormatted := parser.ParseCommand(string(buf), d)
	fmt.Println(resFormatted)
	resFormatted = "+OK"

	fmt.Println(resFormatted)
	_, err = connection.Write([]byte(resFormatted))
	if err != nil {
		panic(err)
	}

	echo(connection, d)
}
