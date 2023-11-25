package database

import (
	"fmt"

	"github.com/GOXayyasang/golang/config"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func InitDB() (*sqlx.DB, error) {
	dbConf := config.GetDatabaseConfig()
	connString := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s", dbConf.User, dbConf.Password, dbConf.Server, dbConf.Database)
	conn, err := sqlx.Open("sqlserver", connString)
	if err != nil {
		return nil, err
	}
	DB = conn
	// err = conn.Ping()
	// if err != nil {
	// 	return nil, err
	// }
	return conn, nil
}
