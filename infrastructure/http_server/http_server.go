package httpserver

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type IHTTPServer interface {
	Setup()
	RegisterRoute(method string, endpoint string, handler ...gin.HandlerFunc) error
	Run() error
}

type HTTPServer struct {
	router  *gin.Engine
	address string
	server  *http.Server
}

func (server *HTTPServer) Setup() {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	server.address = fmt.Sprintf("%s:%s", host, port)

	server.router = gin.New()

	server.server = &http.Server{
		Addr:    server.address,
		Handler: server.router,
	}
}

func (server *HTTPServer) RegisterRoute(method string, endpoint string, handler ...gin.HandlerFunc) error {
	switch method {
	case "POST":
		server.router.POST(endpoint, handler...)
	case "GET":
		server.router.GET(endpoint, handler...)
	case "PUT": // Put updates the entire resource
		server.router.PUT(endpoint, handler...)
	case "PATCH": // Patch updates partially
		server.router.PATCH(endpoint, handler...)
	case "DELETE":
		server.router.DELETE(endpoint, handler...)
	default:
		return errors.New("ethod now allowed")
	}

	return nil
}

func (server *HTTPServer) Run() error {
	//certPath := os.Getenv("TLS_CERT_PATH")
	//keyPath := os.Getenv("TLS_KEY_PATH")
	

	err := server.server.ListenAndServe()
	if err != nil {
		logrus.Error(fmt.Sprintf("Error while trying to serve HTTPS: %s", err.Error()))
		return err
	}

	logrus.Error(fmt.Sprintf("Server running at https://%s", server.address))
	return nil
}

func NewHTTPServer() *HTTPServer {
	return &HTTPServer{}
}