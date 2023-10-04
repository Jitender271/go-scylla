package interfaces

import (
	"context"

	"github.com/go-scylla/domain/dtos"
)


type IReportingData interface{
	AddReportingDetail(ctx context.Context, reportingData *dtos.ReportingDataDTO) (*dtos.ReportingDataDTO, error)

}