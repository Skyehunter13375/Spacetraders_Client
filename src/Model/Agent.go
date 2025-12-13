package Model

type Agent struct {
	AccountID   string `json:"accountId"`
	Symbol      string `json:"symbol"`
	HQ          string `json:"headquarters"`
	Credits     int    `json:"credits"`
	Faction     string `json:"startingFaction"`
	Ships       int    `json:"shipCount"`
	LastUpdated string
}

