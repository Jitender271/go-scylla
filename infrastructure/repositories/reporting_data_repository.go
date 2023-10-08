package repositories

import (
	"context"

	"github.com/go-scylla/application/interfaces"
	"github.com/go-scylla/domain/dtos"
	"github.com/go-scylla/internal/model"
	"github.com/gocql/gocql"
	"github.com/sirupsen/logrus"
)

type reportingDataRepository struct{
	queryBuilder interfaces.IQueryBuilder[model.ReportingDetail]
}


func (repo *reportingDataRepository) AddReportingDetail(ctx context.Context, reportingData *dtos.ReportingDataDTO) (*dtos.ReportingDataDTO, error){
	err := repo.queryBuilder.Insert(ctx, reportingDataDTOToEntity(reportingData))
	if err != nil{
		logrus.Error("could not add tracking data. error: ", err.Error())
		return nil, err

	}
	return reportingData, nil
}


func reportingDataDTOToEntity(dto *dtos.ReportingDataDTO) *model.ReportingDetail{

	reportingDataEntity := &model.ReportingDetail{
		ID:   gocql.TimeUUID(),
		Name: dto.Name,
		Data: dto.Details,
	}
	return reportingDataEntity
}

func NewReportingDataRepository(querybuilder interfaces.IQueryBuilder[model.ReportingDetail]) *reportingDataRepository {
	return &reportingDataRepository{
		queryBuilder: querybuilder,
	
	}
}
