package Waypoints

import (
	"Spacetraders/src/General"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
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
		General.LogErr(fmt.Sprintf("%v", err))
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

	for idx := range s.Waypoints {
		_, err = db.Exec(`
			INSERT INTO waypoints (system,symbol,type,x_coord,y_coord,orbits)
			VALUES ($1, $2, $3, $4, $5, $6)
			ON CONFLICT (symbol) DO UPDATE SET
				system  = EXCLUDED.system,
				symbol  = EXCLUDED.symbol,
				type    = EXCLUDED.type,
				x_coord = EXCLUDED.x_coord,
				y_coord = EXCLUDED.y_coord,
				orbits  = EXCLUDED.orbits
			`,
			s.Symbol,
			s.Waypoints[idx].Symbol,
			s.Waypoints[idx].Type,
			s.Waypoints[idx].Xcoord,
			s.Waypoints[idx].Ycoord,
			s.Waypoints[idx].Orbits,
		)
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
		General.LogErr(err.Error())
	}

	db, _ := sql.Open("postgres", "user=skyehunter dbname=spacetraders sslmode=disable")
	defer db.Close()

	_, err = db.Exec(`
		INSERT INTO ???? ()
		VALUES ()
		ON CONFLICT (symbol) DO UPDATE SET
		`,
		w.Symbol,
	)
	if err != nil {
		General.LogErr(fmt.Sprintf("%v", err))
	}

	return nil
}
