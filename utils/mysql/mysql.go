package utils

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	
	"github.com/SIMPLYBOYS/column_grouping/models"
)

type MySqlClientUtil struct {
}

func NewMySqlClientUtil() *MySqlClientUtil {
	return &MySqlClientUtil{}
}

var SqlConfig models.MysqlConfig

func InitialMysql() {
	SqlConfig = loadConfig()
	fmt.Print(SqlConfig)
}

func loadConfig() (config models.MysqlConfig) {
	// Open our jsonFile
	jsonFile, err := os.Open("config.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened config.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var mysqlConfig models.Configs

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &mysqlConfig)

	// fmt.Print(mysqlConfig.Configs[0])
	return mysqlConfig.Configs[0]
}

func GetOpenMysqlClient() (*sql.DB, error) {
	sourceName := SqlConfig.User + ":" + SqlConfig.Password + "@tcp(" + SqlConfig.Host + ":" + SqlConfig.Port + ")/" + models.DataBase + "?charset=utf8"
	db, err := sql.Open("mysql", sourceName)
	return db, err
}
