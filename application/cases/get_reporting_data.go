package cases

import (
	"context"

	"github.com/go-scylla/application/interfaces"
	"github.com/go-scylla/domain/dtos"
)


type getReportingData struct{
	reportingData interfaces.IReportingData
}


func(repo *getReportingData) Perform(ctx context.Context, reportingData dtos.ReportingDataPrimaryDTO)(*dtos.ReportingDetailDataDTO, error){
	return repo.reportingData.GetReportingDetail(ctx, &reportingData)
}

func GetReportingData(repo interfaces.IReportingData) *getReportingData{
	return &getReportingData{
		repo,
	}
}