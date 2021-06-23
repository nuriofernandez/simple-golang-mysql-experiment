package fetchers

import (
	"database/sql"
	mysql "github.com/xXNurioXx/simple-golang-mysql-experiment/database/connection"
	. "github.com/xXNurioXx/simple-golang-mysql-experiment/structs"
)

const query = "SELECT `server_id`, `server_domain`, `players`, `max_players`, `server_score`, `server_icon_id` FROM `server_list` ORDER BY `last_ping` ASC LIMIT 10"

func queryForResults() *sql.Rows {
	// Obtain the MySQL connection
	connection := mysql.GetConnection()

	// Execute the query and return the result rows
	results, err := connection.Query(query)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	return results
}

func parse(results *sql.Rows) MinecraftServer {
	var server MinecraftServer

	// Scan the result into the MinecraftServer composite struct
	err := results.Scan(&server.Id, &server.Domain, &server.OnlinePlayers, &server.MaxPlayers, &server.Score, &server.Image)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic
	}

	return server
}

func GetServers() []MinecraftServer {
	// Obtain row results from sql database
	results := queryForResults()

	// Collect all minecraft servers structs and return them all
	var servers []MinecraftServer
	for results.Next() {
		servers = append(servers, parse(results))
	}
	return servers
}
