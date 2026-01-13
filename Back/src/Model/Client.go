package Model

type GameState struct {
	IsPaused   int    `json:"isPaused"`
	ServerTS   string `json:"serverTS"`
	AgentTS    string `json:"agentTS"`
	FleetTS    string `json:"fleetTS"`
	ContractTS string `json:"contractTS"`
	SystemsTS  string `json:"systemsTS"`
}
