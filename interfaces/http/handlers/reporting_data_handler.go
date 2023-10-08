package handlers

import (
	"github.com/go-scylla/domain/cases"
	"github.com/go-scylla/domain/dtos"
	httpserver "github.com/go-scylla/infrastructure/http_server"
	"github.com/go-scylla/interfaces/http/factories"
)

type IReportingDataHandler interface{
	Create(httprequest httpserver.HttpRequest) httpserver.HttpResponse
}


type reportingHandler struct{
	httpResponseFactory factories.HttpResponseFactory
	createReporting  cases.ICreateReportingDataRepo
}


func (reportinghandler *reportingHandler) Create(request httpserver.HttpRequest) httpserver.HttpResponse{
	dto := dtos.ReportingDataDTO{}

	err := dtos.ParseJson(request.Body, &dto, "ReportingDataDto")
	
	if err!=nil{
		reportinghandler.httpResponseFactory.BadRequest("body must be valid json", nil)
	}

	result, err := reportinghandler.createReporting.Perform(request.Ctx, dto)
	if err!=nil{
		reportinghandler.httpResponseFactory.ErrorResponseMapper(err, nil)
	}
	return reportinghandler.httpResponseFactory.Created(result, nil)

}

func NewReportingDataHandler(createReportingData  cases.ICreateReportingDataRepo) *reportingHandler{
	httpResponseFactory := factories.NewHttpResponseFactory()
	return &reportingHandler{
		httpResponseFactory,
		createReportingData,
	}
}