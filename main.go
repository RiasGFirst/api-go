// main.go
package main

import (
	"api-go/api"
	"api-go/databases"
	"flag"
	"fmt"
	"log"
	//"github.com/google/uuid"
)

func main() {
    fmt.Println("[API] Server started")
	databases.InitDB()

	listenAddr := flag.String("listen-addr", ":3000", "server listen address")
	flag.Parse()

	server := api.NewServer(*listenAddr)
	log.Fatal(server.Start())
}