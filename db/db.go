package db

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/XSAM/otelsql"
	logger "github.com/anyTV/gomodules/v2/logging"
	_ "github.com/go-sql-driver/mysql"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.40.0"
)

var connections map[string]*sql.DB = make(map[string]*sql.DB)
var register map[string]metric.Registration = make(map[string]metric.Registration)

func CreateDataSourceName(v DbConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?multiStatements=true&parseTime=true",
		v.User, v.Pass, v.Host, v.Port, v.Db,
	)
}

type DbConfig struct {
	User string
	Pass string
	Host string
	Port string
	Db   string
}

func CreateConnection(d DbConfig) (*sql.DB, error) {
	return sql.Open("mysql", CreateDataSourceName(d))
}

func CreateConnectionWithOTEL(d DbConfig) (*sql.DB, error) {
	return otelsql.Open(
		"mysql",
		CreateDataSourceName(d),
		otelsql.WithAttributes(semconv.DBSystemNameMySQL),
		otelsql.WithDisableSkipErrMeasurement(true),
		otelsql.WithSpanOptions(otelsql.SpanOptions{
			DisableErrSkip: true,
			OmitConnResetSession: true,
			OmitConnPrepare: true,
		}),
	)
}

func AddConnectionWithOTEL(key string, d DbConfig) (*sql.DB, error) {
	con, err := CreateConnectionWithOTEL(d)

	if err != nil {
		logger.Fatalf("Failed create connection(%s): %s", d.Db, err)
		return nil, errors.Join(err, fmt.Errorf("failed to create db connection"))
	}

	reg, err := otelsql.RegisterDBStatsMetrics(
		con, otelsql.WithAttributes(semconv.DBSystemNameMySQL),
	)

	if err != nil {
		logger.Errorf("Failed to register stat metrics (%s): %s", d.Db, err)
		return nil, errors.Join(err, errors.New("failed to register stat metrics"))
	}

	connections[key] = con
	register[key] = reg

	return con, nil
}

func CloseAll() {
	for k := range connections {
		connections[k].Close()
		register[k].Unregister()
	}
}

func AddConnection(key string, d DbConfig) (*sql.DB, error) {
	con, err := CreateConnection(d)

	if err != nil {
		logger.Fatalf("Failed create connection(%s): %s", d.Db, err)
		return nil, errors.Join(err, errors.New("failed to create connection"))
	}

	connections[key] = con

	return con, nil
}

func GetConnection(p string) (*sql.DB, bool) {
	con, ok := connections[p]

	return con, ok
}

func GetConnections() map[string]*sql.DB {
	return connections
}
