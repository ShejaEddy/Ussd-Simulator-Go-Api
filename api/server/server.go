package server

import (
	"github.com/gorilla/handlers"
	RouterFactory "github.com/projects/ussd/api/router"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Server struct {
	Port       int
	Addr       string
	httpServer *http.Server
}

func (s *Server) Start() {
	log.Println("Server has started on port", s.Port)
	log.Fatal(s.httpServer.ListenAndServe())
}

func NewServer(port int) *Server {
	var s Server

	s.Port = port
	s.Addr = ":" + strconv.Itoa(port)

	router := RouterFactory.NewRouter()

	handler := handlers.LoggingHandler(os.Stdout, handlers.CORS()(router.Router))

	s.httpServer = &http.Server{
		Addr:           s.Addr,
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return &s
}
