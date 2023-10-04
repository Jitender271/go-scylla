package repositories

import (
	"context"

	"github.com/go-scylla/application/interfaces"
	"github.com/go-scylla/domain/dtos"
	"github.com/go-scylla/internal/model"
	"github.com/sirupsen/logrus"
)

type reportingDataRepository struct{
	queryBuilder interfaces.IQueryBuilder[model.ReportingDetail]
}


func (repo *reportingDataRepository) AddReportingDetail(ctx context.Context, reportingData *dtos.ReportingDataDTO) (*dtos.ReportingDataDTO, error){
	err := repo.queryBuilder.Insert(ctx, dtos.ReportingDataDTOToEntity(reportingData))
	if err != nil{
		logrus.Error("could not add tracking data. error: ", err.Error())
		return nil, err

	}
	return reportingData, nil
}