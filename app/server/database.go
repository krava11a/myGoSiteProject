package server

import (
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func InitDB() (err error)  {
	var dataSourceName = "sqlserver://sa:P@ssw0rd@localhost:1433?database=test"
	Db ,err = sqlx.Connect("sqlserver",dataSourceName)
	if err!=nil{
		return
	}
	err = Db.Ping()
	return
}