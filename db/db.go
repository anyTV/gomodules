package db

import (
	"database/sql"
	"errors"
	"fmt"

	logger "github.com/anyTV/gomodules/logging"
	_ "github.com/go-sql-driver/mysql"
)

var connections map[string]*sql.DB = make(map[string]*sql.DB)

func CreateDataSourceName(v DbConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?multiStatements=true",
		v.User, v.Pass, v.Host, v.Port, v.Db,
	)
}

type DbConfig struct {
	User string
	Pass string
	Host string
	Port string
	Db string
}

func CreateConnection(key string, d DbConfig) (*sql.DB, error) {
	dbCon, err := sql.Open("mysql", CreateDataSourceName(d))

	if err != nil {
		logger.Fatalf("Failed create connection(%s): %s", d.Db, err)
		return nil, errors.Join(fmt.Errorf("failed to create connection: %s"), err)
	}

	connections[key] = dbCon

	return dbCon, nil
}

func GetConnection(p string) (*sql.DB) {
	return connections[p]
}
