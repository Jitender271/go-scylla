package handlers

import (
	"github.com/go-scylla/domain/cases"
	"github.com/go-scylla/domain/dtos"
	httpserver "github.com/go-scylla/infrastructure/http_server"
	"github.com/go-scylla/interfaces/http/factories"
)

type IReportingDataHandler interface{
	Create(httprequest httpserver.HttpRequest) httpserver.HttpResponse
	FindByPrimaryKey(httprequest httpserver.HttpRequest) httpserver.HttpResponse
	DeleteByPrimaryKey(httprequest httpserver.HttpRequest) httpserver.HttpResponse
	GetAll(httpserver httpserver.HttpRequest) httpserver.HttpResponse
}


type reportingHandler struct{
	httpResponseFactory factories.HttpResponseFactory
	createReporting  cases.ICreateReportingDataRepo
	findByPrimaryKey  cases.IGetReportingDataRepo
	deleteByPrimaryKey cases.IDeleteUseCase
	getAllReportingData cases.IGetAllReportingData
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

func (reportinghandler *reportingHandler) FindByPrimaryKey(request httpserver.HttpRequest) httpserver.HttpResponse{
	dto := dtos.ReportingDataPrimaryDTO{}

	err := dtos.ParseJson(request.Body, &dto, "ReportingDataDto")
	
	if err!=nil{
		reportinghandler.httpResponseFactory.BadRequest("body must be valid json", nil)
	}


	result, err := reportinghandler.findByPrimaryKey.Perform(request.Ctx, dto)
	if err!=nil{
		reportinghandler.httpResponseFactory.ErrorResponseMapper(err, nil)
	}
	return reportinghandler.httpResponseFactory.Ok(result, nil)
}


func (reportinghandler *reportingHandler) DeleteByPrimaryKey(request httpserver.HttpRequest) httpserver.HttpResponse{
	dto := dtos.ReportingDataPrimaryDTO{}

	err := dtos.ParseJson(request.Body, &dto, "ReportingPrimarykeyDataDto")
	
	if err!=nil{
		reportinghandler.httpResponseFactory.BadRequest("body must be valid json", nil)
	}

	if err := reportinghandler.deleteByPrimaryKey.Perform(request.Ctx, dto); err!=nil{
		reportinghandler.httpResponseFactory.ErrorResponseMapper(err, nil)
	}
	return reportinghandler.httpResponseFactory.Ok(nil, nil)
}

func (reportingHandler *reportingHandler) GetAll(request httpserver.HttpRequest) httpserver.HttpResponse{
	result, err := reportingHandler.getAllReportingData.Perform(request.Ctx)
	if err!=nil{
		reportingHandler.httpResponseFactory.ErrorResponseMapper(err, nil)
	}
	return reportingHandler.httpResponseFactory.Ok(result, nil)

}

func NewReportingDataHandler(createReportingData  cases.ICreateReportingDataRepo, getReportingData cases.IGetReportingDataRepo, deleteReportingData cases.IDeleteUseCase, getAllReportingData cases.IGetAllReportingData) *reportingHandler{
	httpResponseFactory := factories.NewHttpResponseFactory()
	return &reportingHandler{
		httpResponseFactory,
		createReportingData,
		getReportingData,
		deleteReportingData,
		getAllReportingData,
	}
}