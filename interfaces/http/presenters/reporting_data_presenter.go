package presenters

import (
	"github.com/go-scylla/infrastructure/adapters"
	httpserver "github.com/go-scylla/infrastructure/http_server"
	"github.com/go-scylla/interfaces/http/handlers"
)

type reportingDataPresenter struct{

	handler handlers.IReportingDataHandler

}

func (presenter *reportingDataPresenter) Register(httpserver httpserver.IHTTPServer){
	httpserver.RegisterRoute("POST" , "api/v1/reporting", adapters.HandlerAdapter(presenter.handler.Create))
	httpserver.RegisterRoute("POST" , "api/v1/report", adapters.HandlerAdapter(presenter.handler.FindByPrimaryKey))
	httpserver.RegisterRoute("DELETE" , "api/v1/delete/report", adapters.HandlerAdapter(presenter.handler.DeleteByPrimaryKey))

}

func NewDataPresenter(handler handlers.IReportingDataHandler) *reportingDataPresenter{
	return &reportingDataPresenter{
		handler,
	}
	
}