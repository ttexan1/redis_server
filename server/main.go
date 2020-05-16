package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"

	"redis_app/parser"
	"redis_app/store"
	"redis_app/usecase"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:10000")
	if err != nil {
		panic(err)
	}
	db := map[string]*store.DB{}

	log.Println("Server running at localhost:10000")

	waitClient(listener, db)
}

func waitClient(listener net.Listener, db map[string]*store.DB) {
	connection, err := listener.Accept()
	if err != nil {
		panic(err)
	}
	sessionID := string(rand.Intn(1 << 20))
	go goEcho(connection, db, sessionID)
	waitClient(listener, db)
}

func goEcho(connection net.Conn, db map[string]*store.DB, sID string) {
	fmt.Println(connection.LocalAddr())
	db[sID] = store.NewDB()
	defer func() {
		db[sID] = nil
		connection.Close()
	}()
	uc := usecase.NewUseCase(db[sID].Single, db[sID].List)
	// parser :=InitParser

	var echo func(net.Conn, *store.DB, *usecase.UseCase)
	echo = func(connection net.Conn, d *store.DB, uc *usecase.UseCase) {
		var buf = make([]byte, 1024)
		_, err := connection.Read(buf)
		if err != nil {
			if err == io.EOF {
				return
			}
			panic(err)
		}
		fmt.Println(">", string(buf))

		resFormatted := parser.ParseCommand(string(buf), d, uc)
		fmt.Println(resFormatted)
		// resFormatted := "+OK"
		_, err = connection.Write([]byte(resFormatted))
		if err != nil {
			panic(err)
		}

		echo(connection, d, uc)
	}
	echo(connection, db[sID], uc)
}
