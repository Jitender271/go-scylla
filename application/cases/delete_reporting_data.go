package cases

import (
	"context"

	"github.com/go-scylla/application/interfaces"
	"github.com/go-scylla/domain/dtos"
)

type deleteReportingData struct{
	reportingData interfaces.IReportingData
}


func(repo *deleteReportingData) Perform(ctx context.Context, reportingData dtos.ReportingDataPrimaryDTO)(error){
	return repo.reportingData.DeleteReportingDetail(ctx, &reportingData)
}

func DeleteReportingData(repo interfaces.IReportingData) *deleteReportingData{
	return &deleteReportingData{
		repo,
	}
}