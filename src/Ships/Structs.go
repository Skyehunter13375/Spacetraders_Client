package Ships

type Ship struct {
	Symbol       string        `json:"symbol"`
	Registration ShipReg       `json:"registration"`
	Nav          ShipNav       `json:"nav"`
	Crew         ShipCrew      `json:"crew"`
	Frame        ShipFrame     `json:"frame"`
	Reactor      ShipReactor   `json:"reactor"`
	Engine       ShipEngine    `json:"engine"`
	Cooldown     ShipCooldown  `json:"cooldown"`
	Modules      []ShipModules `json:"modules"`
	Mounts       []ShipMounts  `json:"mounts"`
	Cargo        ShipCargo     `json:"cargo"`
	Fuel         ShipFuel      `json:"fuel"`
	LastUpdated  string
}

type ShipReg struct {
	Name          string `json:"name"`
	FactionSymbol string `json:"factionSymbol"`
	Role          string `json:"role"`
}

type ShipNav struct {
	SystemSymbol   string `json:"systemSymbol"`
	WaypointSymbol string `json:"waypointSymbol"`
	Status         string `json:"status"`
	FlightMode     string `json:"flightMode"`
	Route          struct {
		Destination struct {
			Symbol       string `json:"symbol"`
			Type         string `json:"type"`
			SystemSymbol string `json:"systemSymbol"`
			X            int    `json:"x"`
			Y            int    `json:"y"`
		} `json:"destination"`
		Origin struct {
			Symbol       string `json:"symbol"`
			Type         string `json:"type"`
			SystemSymbol string `json:"systemSymbol"`
			X            int    `json:"x"`
			Y            int    `json:"y"`
		} `json:"origin"`
		DepartureTime string `json:"departureTime"`
		Arrival       string `json:"arrival"`
	} `json:"route"`
}

type ShipCrew struct {
	Current  int    `json:"current"`
	Required int    `json:"required"`
	Capacity int    `json:"capacity"`
	Rotation string `json:"rotation"`
	Morale   int    `json:"morale"`
	Wages    int    `json:"wages"`
}

type ShipFrame struct {
	Symbol         string `json:"symbol"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Condition      int    `json:"condition"`
	Integrity      int    `json:"integrity"`
	ModuleSlots    int    `json:"moduleSlots"`
	MountingPoints int    `json:"mountingPoints"`
	FuelCapacity   int    `json:"fuelCapacity"`
	Quality        int    `json:"quality"`
	Requirements   struct {
		Power int `json:"power"`
		Crew  int `json:"crew"`
		Slots int `json:"slots"`
	} `json:"requirements"`
}

type ShipReactor struct {
	Symbol       string `json:"symbol"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Condition    int    `json:"condition"`
	Integrity    int    `json:"integrity"`
	PowerOutput  int    `json:"powerOutput"`
	Quality      int    `json:"quality"`
	Requirements struct {
		Power int `json:"power"`
		Crew  int `json:"crew"`
		Slots int `json:"slots"`
	} `json:"requirements"`
}

type ShipEngine struct {
	Symbol       string `json:"symbol"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Condition    int    `json:"condition"`
	Integrity    int    `json:"integrity"`
	Speed        int    `json:"speed"`
	Quality      int    `json:"quality"`
	Requirements struct {
		Power int `json:"power"`
		Crew  int `json:"crew"`
		Slots int `json:"slots"`
	} `json:"requirements"`
}

type ShipCooldown struct {
	ShipSymbol       string `json:"shipSymbol"`
	TotalSeconds     int    `json:"totalSeconds"`
	RemainingSeconds int    `json:"remainingSeconds"`
	Expiration       string `json:"expiration"`
}

type ShipModules struct {
	Symbol       string `json:"symbol"`
	Capacity     int    `json:"capacity"`
	Range        int    `json:"range"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Requirements struct {
		Power int `json:"power"`
		Crew  int `json:"crew"`
		Slots int `json:"slots"`
	} `json:"requirements"`
}

type ShipMounts struct {
	Symbol       string   `json:"symbol"`
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Strength     int      `json:"strength"`
	Deposits     []string `json:"deposits"`
	Requirements struct {
		Power int `json:"power"`
		Crew  int `json:"crew"`
		Slots int `json:"slots"`
	} `json:"requirements"`
}

type ShipCargo struct {
	Capacity  int `json:"capacity"`
	Units     int `json:"units"`
	Inventory []struct {
		Symbol      string `json:"symbol"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Units       int    `json:"units"`
	} `json:"inventory"`
}

type ShipFuel struct {
	Current  int `json:"current"`
	Capacity int `json:"capacity"`
	Consumed struct {
		Amount    int    `json:"amount"`
		Timestamp string `json:"timestamp"`
	} `json:"consumed"`
}
