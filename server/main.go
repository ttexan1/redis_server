package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"

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
	listener, err := net.Listen("tcp", "localhost:10000")
	if err != nil {
		panic(err)
	}
	db := store.NewDB()

	log.Println("Server running at localhost:10000")
	waitClient(listener, db)
}

func waitClient(listener net.Listener, db *store.DB) {
	connection, err := listener.Accept()
	if err != nil {
		panic(err)
	}
	go goEcho(connection, db)
	waitClient(listener, db)
}

func goEcho(connection net.Conn, db *store.DB) {
	fmt.Println(connection.LocalAddr())
	// db = store.NewDB()
	defer func() {
		connection.Close()
	}()
	uc := usecase.NewUsecase(db.Single, db.List)
	parsers := parser.InitParser(uc)

	var echo func(net.Conn, *usecase.Usecase)
	echo = func(connection net.Conn, uc *usecase.Usecase) {
		time.Sleep(time.Second)
		// connection.Write([]byte("+OK\r\n"))
		var buf = make([]byte, 1024)
		_, err := connection.Read(buf)
		if err != nil {
			if err == io.EOF {
				return
			}
			panic(err)
		}
		fmt.Println(">", string(buf))

		resFormatted := parser.HandleRequest(string(buf), parsers)
		fmt.Println(resFormatted)
		_, err = connection.Write([]byte(resFormatted))
		if err != nil {
			panic(err)
		}
		echo(connection, uc)
	}

	echo(connection, uc)
}
