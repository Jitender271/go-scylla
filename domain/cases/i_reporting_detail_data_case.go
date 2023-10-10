package cases

import (
	"context"

	"github.com/go-scylla/domain/dtos"
)

type IGetReportingDataRepo interface{
	Perform(ctx context.Context, reportingData dtos.ReportingDataPrimaryDTO)(*dtos.ReportingDetailDataDTO, error)

}