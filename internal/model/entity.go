package model

import (
	"github.com/gocql/gocql"
)

type ReportingDetail struct {
    ID   gocql.UUID `db:"id"`
    Name string     `db:"name"`
    Data string     `db:"data"`
}