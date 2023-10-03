package httpserver

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)


type IHttpServer interface{
	setup()


}

type HttpServer struct{
	address string
	router *gin.Engine
	server *http.Server
}


func (server *HttpServer) setup(){
	host:=os.Getenv("host")
	port := os.Getenv("port")

	server.address = fmt.Sprintf("%s,%s", host, port)
	server.router = gin.New()

	server.server = &http.Server{
		Addr: server.address,
		Handler: server.router,
	}
}


func (server *HttpServer) RegisterRoute(){


}



func (server *HttpServer) Run(){
	
}