package Model

type Faction struct {
	Symbol      string `json:"symbol"`
	Name        string `json:"name"`
	Description string `json:"description"`
	HQ          string `json:"headquarters"`
	Traits      []struct{
		Symbol  string `json:"symbol"`
		Name    string `json:"name"`
		Description string `json:"description"`
	} `json:"traits"`
	Recruiting  bool `json:"isRecruiting"`
}
