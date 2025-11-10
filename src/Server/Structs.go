package Server

type GameState struct {
	Status       string     `json:"status"`
	Version      string     `json:"version"`
	LastReset    string     `json:"resetDate"`
	ServerResets ServResets `json:"serverResets"`
	Stats        ServStats  `json:"stats"`
	LastCheckIn  string
}

type ServResets struct {
	NextReset string `json:"next"`
	ResetFreq string `json:"frequency"`
}

type ServStats struct {
	Accounts  int `json:"accounts"`
	Agents    int `json:"agents"`
	Ships     int `json:"ships"`
	Systems   int `json:"systems"`
	Waypoints int `json:"waypoints"`
}
