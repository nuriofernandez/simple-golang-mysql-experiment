package main

import (
	"fmt"
	"github.com/xXNurioXx/simple-golang-mysql-experiment/database/fetchers"
)

func main() {
	servers := fetchers.GetServers()
	for _, server := range servers {
		fmt.Println(fmt.Sprintf("Server '%s' has %d online players", server.Domain, server.OnlinePlayers))
	}
}
