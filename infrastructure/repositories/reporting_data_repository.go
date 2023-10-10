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

func (repo *reportingDataRepository) GetReportingDetail(ctx context.Context, reportingData *dtos.ReportingDataPrimaryDTO) (*dtos.ReportingDetailDataDTO, error){
	result, err := repo.queryBuilder.Get(ctx, reportingDataPrimaryKeyDTOToEntity(reportingData))
	if err != nil{
		logrus.Error("could not add tracking data. error: ", err.Error())
		return nil, err
	}
	if result == nil {
		return nil, nil
	}
	return reportingDataEntityToDTO(result), nil
}

func (repo *reportingDataRepository) DeleteReportingDetail(ctx context.Context, reportingData *dtos.ReportingDataPrimaryDTO)error{
	if err := repo.queryBuilder.Delete(ctx, reportingDataPrimaryKeyDTOToEntity(reportingData)); err!=nil{
		logrus.Error("could not delete tracking data using primary key. error:  ", err.Error())
		return err
	}
	return nil
}

func (repo *reportingDataRepository) GetAllReportingDetail(ctx context.Context)([]*dtos.ReportingDetailDataDTO, error){
	result, err := repo.queryBuilder.SelectAll(ctx)
	if err != nil{
		logrus.Error("could not get reporting data. error:", err.Error())
		return nil, err
	}
	if result == nil{
		return nil, nil
	}
	reportingData := make([]*dtos.ReportingDetailDataDTO, 0, len(result))

	for _, value := range result{

		reportingData = append(reportingData, reportingDataEntityToDTO(&value))

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

func reportingDataPrimaryKeyDTOToEntity(dto *dtos.ReportingDataPrimaryDTO) (*model.ReportingDetail) {
	uuid, err := gocql.ParseUUID(dto.ID)
	if err != nil {
        logrus.Fatalf("Error converting string to gocql.UUID: %v", err)
    }
	reportingDataEntity := &model.ReportingDetail{
		ID: uuid,
	}

	return reportingDataEntity
}

func reportingDataEntityToDTO(model *model.ReportingDetail) (*dtos.ReportingDetailDataDTO) {
	return &dtos.ReportingDetailDataDTO{
		ID: model.ID.String(),
		Name: model.Name,
		Details: model.Data,
	}
	
}


func NewReportingDataRepository(querybuilder interfaces.IQueryBuilder[model.ReportingDetail]) *reportingDataRepository {
	return &reportingDataRepository{
		queryBuilder: querybuilder,
	
	}
}
