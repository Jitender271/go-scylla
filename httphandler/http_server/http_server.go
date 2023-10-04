package httpserver

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type IHttpServer interface {
	setup()
}

type HttpServer struct {
	address string
	router  *gin.Engine
	server  *http.Server
}

func (server *HttpServer) setup() {
	host := os.Getenv("host")
	port := os.Getenv("port")

	server.address = fmt.Sprintf("%s,%s", host, port)
	server.router = gin.New()

	server.server = &http.Server{
		Addr:    server.address,
		Handler: server.router,
	}
}

func (server *HttpServer) RegisterRoute(method, endpoint string, handler ...gin.HandlerFunc) error{
	switch method{
	case "POST":
		server.router.POST(endpoint, handler...)
	case "GET":
		server.router.GET(endpoint, handler...)

	default:
		return errors.New("method not allowed")
	}
	return nil

}

func (server *HttpServer) Run() error {

	err := server.server.ListenAndServe()
	if err != nil {
		logrus.Error(fmt.Sprintf("Error while trying to serve HTTPS: %s", err.Error()))
		return err
	}

	logrus.Info("Server running at https://", server.address)
	return nil

}
