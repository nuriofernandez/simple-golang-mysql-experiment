package structs

type MinecraftServer struct {
	Id            string `json:"id"`
	Domain        string `json:"domain"`
	Score         int    `json:"score"`
	Image         int    `json:"image"`
	OnlinePlayers int    `json:"onlinePlayers"`
	MaxPlayers    int    `json:"maxPlayers"`
}
