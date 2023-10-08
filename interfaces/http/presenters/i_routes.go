package presenters

import httpserver "github.com/go-scylla/infrastructure/http_server"

type IRoutes interface{
	Register(httpserver httpserver.IHTTPServer)
}