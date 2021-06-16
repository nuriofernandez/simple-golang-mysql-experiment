package fetchers

import (
	"database/sql"
	mysql "github.com/xXNurioXx/simple-golang-mysql-experiment/database/connection"
	. "github.com/xXNurioXx/simple-golang-mysql-experiment/structs"
)

const query = "SELECT `server_id`, `server_domain`, `players`, `max_players` FROM `server_list` ORDER BY `last_ping` ASC LIMIT 10"

func queryForResults() *sql.Rows {
	connection := mysql.GetConnection()

	// Execute the query
	results, err := connection.Query(query)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	return results
}

func GetServers() []MinecraftServer {
	// Obtain row results from sql database
	results := queryForResults()

	var servers []MinecraftServer
	for results.Next() {
		var server MinecraftServer
		// for each row, scan the result into the MinecraftServer composite object
		err := results.Scan(&server.Id, &server.Domain, &server.OnlinePlayers, &server.MaxPlayers)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic
		}

		// Redefine servers array with the new server
		servers = append(servers, server)
	}
	return servers
}
