package main

import (
	"log"
	"os"

	"github.com/gauravsarma1992/todo-tf-provider/server"
)

func main() {
	var (
		apiServer *server.ApiServer
		err       error
	)
	log.Println("Starting API server")
	if apiServer, err = server.Default(); err != nil {
		log.Println(err)
		os.Exit(-1)
	}
	log.Println(apiServer.Run())
}
