package cases

import (
	"context"

	"github.com/go-scylla/application/interfaces"
	"github.com/go-scylla/domain/dtos"
)


type createReportingData struct{
	reportingData interfaces.IReportingData
}


func(repo *createReportingData) Perform(ctx context.Context, data dtos.ReportingDataDTO)(*dtos.ReportingDataDTO, error){
	return repo.reportingData.AddReportingDetail(ctx, &data)
}

func NewReportingData(repo interfaces.IReportingData) *createReportingData{
	return &createReportingData{
		repo,
	}
}