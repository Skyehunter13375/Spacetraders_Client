package Ships

type ShipData struct {
	Symbol       string `json:"symbol"`
	LastUpdated  string
	Registration struct {
		Name          string `json:"name"`
		FactionSymbol string `json:"factionSymbol"`
		Role          string `json:"role"`
	} `json:"registration"`

	Nav struct {
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
	} `json:"nav"`

	Crew struct {
		Current  int    `json:"current"`
		Required int    `json:"required"`
		Capacity int    `json:"capacity"`
		Rotation string `json:"rotation"`
		Morale   int    `json:"morale"`
		Wages    int    `json:"wages"`
	} `json:"crew"`

	Frame struct {
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
	} `json:"frame"`

	Reactor struct {
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
	} `json:"reactor"`

	Engine struct {
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
	} `json:"engine"`

	Cooldown struct {
		ShipSymbol       string `json:"shipSymbol"`
		TotalSeconds     int    `json:"totalSeconds"`
		RemainingSeconds int    `json:"remainingSeconds"`
		Expiration       string `json:"expiration"`
	} `json:"cooldown"`

	Modules []struct {
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
	} `json:"modules"`

	Mounts []struct {
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
	} `json:"mounts"`

	Cargo struct {
		Capacity  int `json:"capacity"`
		Units     int `json:"units"`
		Inventory []struct {
			Symbol      string `json:"symbol"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Units       int    `json:"units"`
		} `json:"inventory"`
	} `json:"cargo"`

	Fuel struct {
		Current  int `json:"current"`
		Capacity int `json:"capacity"`
		Consumed struct {
			Amount    int    `json:"amount"`
			Timestamp string `json:"timestamp"`
		} `json:"consumed"`
	} `json:"fuel"`
}
