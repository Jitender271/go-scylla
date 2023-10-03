package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/go-scylla/db"
	"github.com/go-scylla/environments"
	"github.com/go-scylla/interfaces"
	"github.com/go-scylla/internal/model"
	"github.com/gocql/gocql"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("started main")
	environments.Init()

	/* Database connection */
	consistency := gocql.ParseConsistency(os.Getenv("SCYLLA_CONSISTENCY"))
	hosts := strings.Split(os.Getenv("SCYLLA_HOSTS"), ",")
	keyspace := "reporting"

	dbDataConnection := db.NewScyllaDBConnection(consistency, keyspace, hosts...)
	session, err := db.GetConnection(dbDataConnection)
	if err != nil {
		panic(err)
	}
	var ctx context.Context



	/* Query Builder */
	trackingModel := model.NewTrackingDataTable().Table
	builder := interfaces.NewQueryBuilder[model.ReportingDetail](trackingModel, session)

	reportingDetail := &model.ReportingDetail{
		ID: gocql.TimeUUID(),
		Name: "JK",
		Data: "data",
	}

	insertError := builder.Insert(ctx, reportingDetail)

	
	if insertError != nil {
		log.Error("Could not add tracking data. Error:  ", err.Error())
	}
	//Select(ctx, trackingDataPartitionKeyDTOToEntity(trackingData))
	
	results, err:= builder.SelectAll(ctx)

	if err != nil {
		
		log.Error("Could not find all tracking data. Error: ", err.Error())
		return
	}

	if results == nil {
		return
	}

	for _, value := range results {
	//	fmt.Print(key)
		fmt.Print(value)
	}



}
