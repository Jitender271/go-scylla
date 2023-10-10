package cmd

import (

	"strings"
	"os"

	"github.com/go-scylla/application/cases"
	_ "github.com/go-scylla/environments"
	"github.com/go-scylla/infrastructure/db/intializedb"
	"github.com/go-scylla/infrastructure/db/scylladb"
	httpserver "github.com/go-scylla/infrastructure/http_server"
	"github.com/go-scylla/infrastructure/repositories"
	"github.com/go-scylla/interfaces/http/handlers"
	"github.com/go-scylla/interfaces/http/presenters"
	"github.com/go-scylla/internal/model"
	"github.com/gocql/gocql"
)

type Container struct {
	server    httpserver.IHTTPServer
	routes    presenters.IRoutes
	dbSession scylladb.ISessionx
}

func NewContainer() *Container {
	/* Database connection */
	consistency := gocql.ParseConsistency(os.Getenv("SCYLLA_CONSISTENCY"))
	hosts := strings.Split(os.Getenv("SCYLLA_HOSTS"), ",")
	keyspace := "reporting"

	dbDataConnection := intializedb.NewScyllaDBConnection(consistency, keyspace, hosts...)
	session, err := intializedb.GetConnection(dbDataConnection)
	if err != nil {
		panic(err)
	}
	//var ctx context.Context

	/* Query Builder */
	trackingModel := model.NewTrackingDataTable().Table

	builder := scylladb.NewQueryBuilder[model.ReportingDetail](trackingModel, session)
	reportingDataRepo := repositories.NewReportingDataRepository(builder)

	/* Use cases */
	createTrackingDataUsecase := cases.NewReportingData(reportingDataRepo)
	getReportingDataUseCase := cases.GetReportingData(reportingDataRepo)
	deleteReportingDataUseCase := cases.DeleteReportingData(reportingDataRepo)

	/* HTTP server */
	httpServer := httpserver.NewHTTPServer()

	/* Handlers */
	routingHandler := handlers.NewReportingDataHandler(createTrackingDataUsecase, getReportingDataUseCase, deleteReportingDataUseCase)

	/* Routes (Presenters) */
	trackingPresenter := presenters.NewDataPresenter(routingHandler)

	return &Container{
		server:    httpServer,
		routes:    trackingPresenter,
		dbSession: session,
	}

}
