package cases

import (
	"context"

	"github.com/go-scylla/domain/dtos"
)


type IGetAllReportingData interface{
	Perform(ctx context.Context)([]*dtos.ReportingDetailDataDTO, error)
}