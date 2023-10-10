package cases

import (
	"context"

	"github.com/go-scylla/application/interfaces"
	"github.com/go-scylla/domain/dtos"
)

type getAllReportingData struct{
	reportingData interfaces.IReportingData

}

func(repo *getAllReportingData) Perform(ctx context.Context)([]*dtos.ReportingDetailDataDTO, error){
	return repo.reportingData.GetAllReportingDetail(ctx)
}

func GetAllReportingDetailData(reportingDetail interfaces.IReportingData)*getAllReportingData{
	return &getAllReportingData{
		reportingDetail,
	}
}