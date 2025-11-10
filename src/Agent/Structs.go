package Agent

type Agent struct {
	Data AgentData `json:"data"`
}

type AgentData struct {
	AccountID   string `json:"accountId"`
	Symbol      string `json:"symbol"`
	HQ          string `json:"headquarters"`
	Credits     int    `json:"credits"`
	Faction     string `json:"startingFaction"`
	Ships       int    `json:"shipCount"`
	LastUpdated string
}
