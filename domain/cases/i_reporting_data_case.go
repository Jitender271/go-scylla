package cases

import (
	"context"

	"github.com/go-scylla/domain/dtos"
)

type ICreateReportingDataRepo interface{
	Perform(ctx context.Context, dto dtos.ReportingDataDTO)(*dtos.ReportingDataDTO, error)

}