package connection

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xXNurioXx/simple-golang-mysql-experiment/config"
)

var connection *sql.DB

func createSQLConnection() *sql.DB {
	// Open up a database connection using the settings.yml credentials.
	settings := config.ReadConfig()
	sourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", settings.Username, settings.Password, settings.Hostname, settings.Port, settings.Database)

	connection, err := sql.Open("mysql", sourceName)

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	return connection
}

func GetConnection() *sql.DB {
	if connection == nil {
		// Database connection is undefined
		connection = createSQLConnection()
	}

	if err := connection.Ping(); err != nil {
		// Database connection is defined but closed
		connection = createSQLConnection()
	}

	// Return database connection
	return connection
}
