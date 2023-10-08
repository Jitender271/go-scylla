package adapters

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	httpserver "github.com/go-scylla/infrastructure/http_server"
	"github.com/sirupsen/logrus"
)

func HandlerAdapter(handler func(httpserver.HttpRequest) httpserver.HttpResponse) gin.HandlerFunc{
	return func(ctx *gin.Context){
		body, err := io.ReadAll(ctx.Request.Body)
		if err != nil{
			logrus.Error("[handleradapter] error while trying to read request body: ", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}
		ctx.Request.Body = io.NopCloser(bytes.NewBuffer(body))

		params := make(map[string]string)
		for _, param := range ctx.Params{
			params[param.Key] = param.Value
		}

		request := httpserver.HttpRequest{
			Body: body,
			Headers: ctx.Request.Header,
			Params: params,
			Query: ctx.Request.URL.Query(),
			Ctx: ctx.Request.Context(),
		}

		result := handler(request)
		ctx.JSON(result.StatusCode, result.Body)

	}

}