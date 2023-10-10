package dtos

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/go-scylla/internal/model"
	"github.com/gocql/gocql"
)

type ReportingDataDTO struct {
	Name    string `json:"name" validate:"required"`
	Details string `json:"details" validate:"required"`
}

type ReportingDetailDataDTO struct {
	ID		string	`json:"id" validate:"required"`
	Name    string `json:"name" validate:"required"`
	Details string `json:"details" validate:"required"`
}

type ReportingDataPrimaryDTO struct {
	ID		string	`json:"id" validate:"required"`
}

func ParseJson(data []byte, dto interface{}, dtoName string) error {
	err := json.Unmarshal(data, dto)
	if err != nil {
		return fmt.Errorf("[ParseJson] Error: %s", err.Error())
	}

	err = isValid(dto, dtoName)
	if err != nil {
		return fmt.Errorf("[isValid] Error: %s", err.Error())
	}

	fmt.Print(err)

	return nil
}

func isValid(dto interface{}, dtoName string) error {
	v := validator.New()

	err := v.Struct(dto)
	if err != nil {
		return fmt.Errorf("error during %s validation: %s", dtoName, err.Error())
	}

	return nil
}

func ReportingDataDTOToEntity(dto *ReportingDataDTO) *model.ReportingDetail{
	return &model.ReportingDetail{
		ID: gocql.TimeUUID(),
		Name: dto.Name,
		Data: dto.Details,
	}

}