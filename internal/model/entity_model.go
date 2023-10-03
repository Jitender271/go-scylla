package model

import (
	
	gocqlxtable "github.com/go-scylla/interfaces"
	"github.com/scylladb/gocqlx/v2/table"
)

type TrackingDataTable struct {
	Table gocqlxtable.ITable
}

func NewTrackingDataTable() *TrackingDataTable {
	trackingDataMetadata := gocqlxtable.Metadata{
		M: &table.Metadata{
			Name: "reporting_details",
			Columns: []string{
				"id", "data", "name",
	
			},
			PartKey: []string{"id"},
			SortKey: []string{},
		},
	}

	trackingDataTable := gocqlxtable.New(*trackingDataMetadata.M)

	return &TrackingDataTable{
		Table: trackingDataTable,
	}
}
