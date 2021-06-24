package main

import (
	"fmt"
	"github.com/xXNurioXx/simple-golang-mysql-experiment/database/fetchers"
)

func main() {
	servers := fetchers.GetServers()
	for _, server := range servers {
		fmt.Printf("Server '%s' has %d online players\n", server.Domain, server.OnlinePlayers)
	}
}
