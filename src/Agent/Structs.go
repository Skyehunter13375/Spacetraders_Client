package Agent

type AgentState struct {
	Data struct {
		AccountID   string `json:"accountId"`
		Symbol      string `json:"symbol"`
		HQ          string `json:"headquarters"`
		Credits     int    `json:"credits"`
		Faction     string `json:"startingFaction"`
		Ships       int    `json:"shipCount"`
		LastUpdated string
	} `json:"data"`
}
