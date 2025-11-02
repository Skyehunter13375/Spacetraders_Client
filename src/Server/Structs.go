package Server

type GameState struct {
	Status       string `json:"status"`
	Version      string `json:"version"`
	LastReset    string `json:"resetDate"`
	ServerResets struct {
		NextReset string `json:"next"`
		ResetFreq string `json:"frequency"`
	} `json:"serverResets"`
	Stats struct {
		Accounts  int `json:"accounts"`
		Agents    int `json:"agents"`
		Ships     int `json:"ships"`
		Systems   int `json:"systems"`
		Waypoints int `json:"waypoints"`
	} `json:"stats"`
	LastCheckIn string
}
