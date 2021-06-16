package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/xXNurioXx/simple-golang-mysql-experiment/config"
	. "github.com/xXNurioXx/simple-golang-mysql-experiment/structs"
)

func main() {
	// Open up a database connection using the settings.yml credentials.
	settings := config.ReadConfig()
	sourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", settings.Username, settings.Password, settings.Hostname, settings.Port, settings.Database)
	db, err := sql.Open("mysql", sourceName)

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished executing
	defer db.Close()

	// Execute the query
	results, err := db.Query("SELECT `server_id`, `server_domain`, `players`, `max_players` FROM `server_list` ORDER BY `last_ping` ASC LIMIT 10")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var server MinecraftServer
		// for each row, scan the result into the MinecraftServer composite object
		err = results.Scan(&server.Id, &server.Domain, &server.OnlinePlayers, &server.MaxPlayers)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic
		}
		// Print out the server's information
		log.Printf(fmt.Sprintf("%s %d", server.Domain, server.OnlinePlayers))
	}
}
