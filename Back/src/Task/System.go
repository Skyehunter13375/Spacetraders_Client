package Task

import "Spacetraders/src/Model"
import "encoding/json"

func UpdateSystem(symbol string) error {
	jsonStr := GetUrlJson("https://api.spacetraders.io/v2/systems/" + symbol, "")
	// LogActivity(string(jsonStr))

	var wrapper map[string]json.RawMessage
	err := json.Unmarshal([]byte(jsonStr), &wrapper)
	if err != nil { LogErr("UpdateSystem: " + err.Error()) }

	var s Model.System
	err = json.Unmarshal(wrapper["data"], &s)
	if err != nil { LogErr("UpdateSystem: " + err.Error()) }

	_, err = PG.Exec(`
		INSERT INTO systems (symbol, sector, constellation, name, type, x, y)
		VALUES (?,?,?,?,?,?,?)
		ON CONFLICT (symbol) DO UPDATE SET
			sector        = EXCLUDED.sector,
			constellation = EXCLUDED.constellation,
			name          = EXCLUDED.name,
			type          = EXCLUDED.type,
			x             = EXCLUDED.x,
			y             = EXCLUDED.y
		`,
		s.Symbol,
		s.Sector,
		s.Constellation,
		s.Name,
		s.Type,
		s.X,
		s.Y,
	)
	if err != nil { LogErr("UpdateSystem: Insert system failed: " + err.Error()) }

	// PERF:
	// Very heavy on API calls at the moment...
	// Also not reading asteroids which we will need later on...
	// I should store the data I have here from the current jsonStr, and only run UpdateWaypoint() later on if needed.
	for idx := range s.Waypoints {
		if s.Waypoints[idx].Type == "ASTEROID" {
			continue
		}
		err = UpdateWaypoint(s.Symbol, s.Waypoints[idx].Symbol)
		if err != nil { LogErr("UpdateSystem: Insert loop failure: " + err.Error()) }
	 }

	return nil
}

func UpdateWaypoint(system string, waypoint string) error {
	jsonStr := GetUrlJson("https://api.spacetraders.io/v2/systems/"+system+"/waypoints/"+waypoint, "")

	var wrapper map[string]json.RawMessage
	json.Unmarshal([]byte(jsonStr), &wrapper)

	var w Model.Waypoint
	err := json.Unmarshal(wrapper["data"], &w)
	if err != nil { LogErr("UpdateWaypoint: " + err.Error()) }

	traitStr := ""
	for idx, val := range w.Traits {
		if idx == 0 {
			traitStr = val.Symbol
		} else {
			traitStr = traitStr + "," + val.Symbol
		}
	}

	modStr := ""
	for idx, val := range w.Modifiers {
		if idx == 0 {
			modStr = val.Symbol
		} else {
			modStr = modStr + "," + val.Symbol
		}
	}

	_, err = PG.Exec(`
		INSERT INTO waypoints (system,symbol,type,x,y,orbits,construction,traits,modifiers)
		VALUES (?,?,?,?,?,?,?,?,?)
		ON CONFLICT (symbol) DO UPDATE SET
			system       = EXCLUDED.system,
			symbol       = EXCLUDED.symbol,
			type         = EXCLUDED.type,
			x            = EXCLUDED.x,
			y            = EXCLUDED.y,
			orbits       = EXCLUDED.orbits,
			construction = EXCLUDED.construction,
			traits       = EXCLUDED.traits,
			modifiers    = EXCLUDED.modifiers
		`,
		w.System,
		w.Symbol,
		w.Type,
		w.X,
		w.Y,
		w.Orbits,
		w.Construction,
		traitStr,
		modStr,
	)
	if err != nil { LogErr("UpdateWaypoint: Insert failed: " + err.Error()) }

	return nil
}

func UpdateShipyard(system string, symbol string) error {
	jsonStr := GetUrlJson("https://api.spacetraders.io/v2/systems/"+system+"/waypoints/"+symbol+"/shipyard", "")

	var wrapper map[string]json.RawMessage
	json.Unmarshal([]byte(jsonStr), &wrapper)

	var y Model.Shipyard
	err := json.Unmarshal(wrapper["data"], &y)
	if err != nil { LogErr("UpdateShipyard: " + err.Error()) }

	// TODO: Write SQL to upsert data

	return nil
}

func GetSystem(id string) Model.System {
	var Result Model.System

	query := `
		SELECT 
			symbol,
			sector,
			constellation,
			name,
			type,
			x,
			y
		FROM systems
		WHERE symbol = $1
	`

	err := PG.QueryRow(query, id).Scan(
		&Result.Symbol,
		&Result.Sector,
		&Result.Constellation,
		&Result.Name,
		&Result.Type,
		&Result.X,
		&Result.Y,
	)
	if err != nil { LogErr("GetSystem: " + err.Error()); return Result }

	return Result
}

func GetWaypoint(id string) Model.Waypoint {
	var result Model.Waypoint

	PG.QueryRow(`SELECT system,symbol,type,x,y,orbits,construction FROM waypoints WHERE symbol = $1`, id).Scan(
		&result.System,
		&result.Symbol,
		&result.Type,
		&result.X,
		&result.Y,
		&result.Orbits,
		&result.Construction,
	)

	return result
}
