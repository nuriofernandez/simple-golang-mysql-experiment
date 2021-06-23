# GoLang experimental project to test '[spf13/viper](https://github.com/spf13/viper)' and '[go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)' libraries

I used this project to test the '[spf13/viper](https://github.com/spf13/viper)' and '[go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)' libraries.

# Database connection

#### MySQL Connection handler: [Click here to see the source code](https://github.com/xXNurioXx/simple-golang-mysql-experiment/blob/master/src/database/connection/MySQL.go)
```go
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
```

#### Database settings struct definition: [Click here to see the source code](https://github.com/xXNurioXx/simple-golang-mysql-experiment/blob/master/src/config/Settings.go)
```go
type Settings struct {
	Database string `yaml:"database"`

	Username string `yaml:"username"`
	Password string `yaml:"password"`

	Hostname string `yaml:"hostname"`
	Port     int    `yaml:"port"`
}
```

#### Database settings management: [Click here to see the source code](https://github.com/xXNurioXx/simple-golang-mysql-experiment/blob/master/src/config/ViperConfigLoader.go)
```go
func ReadConfig() *Settings {
	// Create a new Viper configuration parser instance
	config := viper.New()

	// Set the file name of the configurations file
	config.SetConfigName("database-settings")
	config.SetConfigType("yml")

	// Set the path to look for the configurations file
	config.AddConfigPath(".")

	if err := config.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	settings := &Settings{}
	unmarshalErr := config.Unmarshal(settings)
	if unmarshalErr != nil {
		fmt.Printf("unable to decode into config struct, %v", unmarshalErr)
	}

	return settings
}
```

# Obtaining data from the database

#### Minecraft server struct definition: [Click here to see the source code](https://github.com/xXNurioXx/simple-golang-mysql-experiment/blob/master/src/structs/MinecraftServer.go)
```go
type MinecraftServer struct {
	Id            string `json:"id"`
	Domain        string `json:"domain"`
	Score         int    `json:"score"`
	Image         int    `json:"image"`
	OnlinePlayers int    `json:"onlinePlayers"`
	MaxPlayers    int    `json:"maxPlayers"`
}
```

#### Minecraft server sql fetcher: [Click here to see the source code](https://github.com/xXNurioXx/simple-golang-mysql-experiment/blob/master/src/database/fetchers/ServerFetcher.go)
```go
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
```

# Final usage of this implementation
```go
func main() {
	servers := fetchers.GetServers()
	for _, server := range servers {
		fmt.Println(fmt.Sprintf("Server '%s' has %d online players", server.Domain, server.OnlinePlayers))
	}
}
```