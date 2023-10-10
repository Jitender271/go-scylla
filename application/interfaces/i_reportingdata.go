package interfaces

import (
	"context"

	"github.com/go-scylla/domain/dtos"
)


type IReportingData interface{
	AddReportingDetail(ctx context.Context, reportingData *dtos.ReportingDataDTO) (*dtos.ReportingDataDTO, error)
	GetReportingDetail(ctx context.Context, reportingData *dtos.ReportingDataPrimaryDTO)(*dtos.ReportingDetailDataDTO, error)
	DeleteReportingDetail(ctx context.Context, reportingData *dtos.ReportingDataPrimaryDTO)error
	GetAllReportingDetail(ctx context.Context)([]*dtos.ReportingDetailDataDTO, error)
}