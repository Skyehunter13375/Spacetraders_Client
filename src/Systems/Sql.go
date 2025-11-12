package Waypoints

import (
	"Spacetraders/src/General"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/lib/pq" // PostgreSQL driver
)

func UpdateSystem(symbol string) error {
	jsonStr := General.GetUrlJson("https://api.spacetraders.io/v2/systems/"+symbol, "")
	// General.LogActivity(string(jsonStr))

	var wrapper map[string]json.RawMessage
	err := json.Unmarshal([]byte(jsonStr), &wrapper)
	if err != nil {
		log.Fatal(err)
	}

	var s System
	err = json.Unmarshal(wrapper["data"], &s)
	if err != nil {
		General.LogErr("UpdateSystem: "+err.Error())
	}

	db, err := sql.Open("postgres", "user=skyehunter dbname=spacetraders sslmode=disable")
	if err != nil {
		General.LogErr(fmt.Sprintf("DB open failed: %v", err))
		return err
	}
	defer db.Close()

	_, err = db.Exec(`
		INSERT INTO systems (symbol, sector, constellation, name, type, x_coord, y_coord)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT (symbol) DO UPDATE SET
			sector        = EXCLUDED.sector,
			constellation = EXCLUDED.constellation,
			name          = EXCLUDED.name,
			type          = EXCLUDED.type,
			x_coord       = EXCLUDED.x_coord,
			y_coord       = EXCLUDED.y_coord
		`,
		s.Symbol,
		s.Sector,
		s.Constellation,
		s.Name,
		s.Type,
		s.Xcoord,
		s.Ycoord,
	)
	if err != nil {
		General.LogErr(fmt.Sprintf("%v", err))
	}

	// PERF:
	// Very heavy on API calls at the moment...
	// Also not reading asteroids which we will need later on...
	// I should store the data I have here from the current jsonStr, and only run UpdateWaypoint() later on if needed.
	for idx := range s.Waypoints {
		if s.Waypoints[idx].Type == "ASTEROID" {
			continue
		}
		err = UpdateWaypoint(s.Symbol, s.Waypoints[idx].Symbol)
		if err != nil {
			General.LogErr(fmt.Sprintf("%v", err))
		}
	}

	return nil
}

func UpdateWaypoint(system string, waypoint string) error {
	jsonStr := General.GetUrlJson("https://api.spacetraders.io/v2/systems/"+system+"/waypoints/"+waypoint, "")

	var wrapper map[string]json.RawMessage
	json.Unmarshal([]byte(jsonStr), &wrapper)

	var w Waypoint
	err := json.Unmarshal(wrapper["data"], &w)
	if err != nil {
		General.LogErr("UpdateWaypoint: " + err.Error())
	}

	traitArr := make([]string, len(w.Traits))
	for idx, val := range w.Traits {
		traitArr[idx] = val.Symbol
	}

	modArr := make([]string, len(w.Modifiers))
	for idx, val := range w.Modifiers {
		modArr[idx] = val.Symbol
	}

	db, _ := sql.Open("postgres", "user=skyehunter dbname=spacetraders sslmode=disable")
	defer db.Close()

	_, err = db.Exec(`
		INSERT INTO waypoints (system,symbol,type,x_coord,y_coord,orbits,construction,traits,modifiers)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
		ON CONFLICT (symbol) DO UPDATE SET
			system       = EXCLUDED.system, 
			symbol       = EXCLUDED.symbol, 
			type         = EXCLUDED.type, 
			x_coord      = EXCLUDED.x_coord, 
			y_coord      = EXCLUDED.y_coord, 
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
		pq.Array(traitArr),
		pq.Array(modArr),
	)
	if err != nil {
		General.LogErr(fmt.Sprintf("%v", err))
	}

	return nil
}

func UpdateShipyard(system string, symbol string) error {
	jsonStr := General.GetUrlJson("https://api.spacetraders.io/v2/systems/"+system+"/waypoints/"+symbol+"/shipyard", "")

	var wrapper map[string]json.RawMessage
	json.Unmarshal([]byte(jsonStr), &wrapper)

	var y Shipyard
	err := json.Unmarshal(wrapper["data"], &y)
	if err != nil {
		General.LogErr("UpdateShipyard: " + err.Error())
	}

	return nil
}
