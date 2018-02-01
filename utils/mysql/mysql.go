package utils

import "database/sql"
import "github.com/SIMPLYBOYS/column_grouping/models"

type MySqlClientUtil struct {
}

func NewMySqlClientUtil() *MySqlClientUtil {
	return &MySqlClientUtil{}
}

func GetOpenMysqlClient() (*sql.DB, error) {
	sourceName := models.Account + ":" + models.Password + "@tcp(" + models.HostName + ":" + models.DBPort + ")/" + models.DataBase + "?charset=utf8"
	db, err := sql.Open("mysql", sourceName)
	return db, err
}
