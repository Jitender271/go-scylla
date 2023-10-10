package cases

import (
	"context"

	"github.com/go-scylla/domain/dtos"
)

type IDeleteUseCase interface{
	Perform(cx context.Context, dto dtos.ReportingDataPrimaryDTO)( error)
}