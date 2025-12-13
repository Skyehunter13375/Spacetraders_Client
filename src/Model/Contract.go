package Model

type Contract struct {
	ID               string        `json:"id"`
	Faction          string        `json:"factionSymbol"`
	Type             string        `json:"type"`
	Terms            ContractTerms `json:"terms"`
	Accepted         bool          `json:"accepted"`
	Fulfilled        bool          `json:"fulfilled"`
	Expiration       string        `json:"expiration"`
	DeadlineToAccept string        `json:"deadlineToAccept"`
	LastUpdated      string
}

type ContractTerms struct {
	Deadline string               `json:"deadline"`
	Payment  ContractPayment      `json:"payment"`
	Deliver  []ContractDeliveries `json:"deliver"`
}

type ContractPayment struct {
	OnAccepted  int `json:"onAccepted"`
	OnFulfilled int `json:"onFulfilled"`
}

type ContractDeliveries struct {
	Material       string `json:"tradeSymbol"`
	Destination    string `json:"destinationSymbol"`
	UnitsRequired  int    `json:"unitsRequired"`
	UnitsFulfilled int    `json:"unitsFulfilled"`
}
