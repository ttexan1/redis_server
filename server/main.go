package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"net/http"
	_ "net/http/pprof"
	"redis_app/parser"
	"redis_app/store"
	"redis_app/usecase"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	const port = 10000
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		panic(err)
	}
	db := store.NewDB()

	log.Println("Server running at localhost:10000")
	for {
		connection, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go goEcho(connection, db)
	}
}

func goEcho(connection net.Conn, db *store.DB) {
	fmt.Println(connection.LocalAddr())
	defer func() {
		connection.Close()
	}()
	uc := usecase.NewUsecase(db.Single, db.List)
	parsers := parser.InitParser(uc)

	for {
		// connection.Write([]byte("+OK\r\n"))
		var buf = make([]byte, 1024)
		_, err := connection.Read(buf)
		if err != nil {
			if err == io.EOF {
				return
			}
			panic(err)
		}
		resFormatted := parser.HandleRequest(string(buf), parsers)
		_, err = connection.Write([]byte(resFormatted))
		if err != nil {
			panic(err)
		}
	}
}
