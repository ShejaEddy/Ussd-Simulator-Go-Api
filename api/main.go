package main

import (
	serverFactory "github.com/projects/ussd/api/server"
	"os"
	"strconv"
)

var port int

func init() {
	rawport := os.Getenv("PORT")

	if len(rawport) > 0 {
		var err error

		port, err = strconv.Atoi(rawport)

		if err != nil {
			panic(err)
		}
	} else {
		port = 2000
	}
}

func main() {
	server := serverFactory.NewServer(port)
	server.Start()
}
