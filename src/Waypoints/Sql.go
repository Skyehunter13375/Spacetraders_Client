package Waypoints

import (
	"Spacetraders/src/General"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
)

func UpdateSystem(symbol string) error {
	jsonStr := General.GetUrlJson("https://api.spacetraders.io/v2/systems/"+symbol, "")
	// General.LogActivity(string(jsonStr))
	var s System
	err := json.Unmarshal([]byte(jsonStr), &s)
	if err != nil {
		log.Fatal(err)
	}

	db, _ := sql.Open("sqlite3", "SpaceTraders.db")
	defer db.Close()

	_, err = db.Exec(`
		INSERT INTO systems (symbol,sector,constellation,name,type,x_coord,y_coord)
		VALUES (?,?,?,?,?,?,?)
		ON CONFLICT (symbol) DO UPDATE SET
			symbol        = EXCLUDED.symbol,
			sector        = EXCLUDED.sector,
			constellation = EXCLUDED.constellation,
			name          = EXCLUDED.name,
			type          = EXCLUDED.type,
			x_coord       = EXCLUDED.x_coord,
			y_coord       = EXCLUDED.y_coord
		`,
		s.Data.Symbol,
		s.Data.Sector,
		s.Data.Constellation,
		s.Data.Name,
		s.Data.Type,
		s.Data.Xcoord,
		s.Data.Ycoord,
	)
	if err != nil {
		General.LogErr(fmt.Sprintf("%v", err))
	}

	for idx := range s.Data.Waypoints {
		// var SQL strings.Builder
		// fmt.Fprintf(&SQL, `INSERT INTO waypoints (system,symbol,x_coord,y_coord,orbits)
		// 	VALUES (%s,%s,%d,%d,%s)
		// 	ON CONFLICT (symbol) DO UPDATE SET
		// 		symbol  = EXCLUDED.system,
		// 		sector  = EXCLUDED.symbol,
		// 		x_coord = EXCLUDED.x_coord,
		// 		y_coord = EXCLUDED.y_coord,
		// 		orbits  = EXCLUDED.orbits`,
		// 	s.Data.Symbol,
		// 	s.Data.Waypoints[idx].Symbol,
		// 	s.Data.Waypoints[idx].Xcoord,
		// 	s.Data.Waypoints[idx].Ycoord,
		// 	s.Data.Waypoints[idx].Orbits,
		// )
		// General.LogActivity(SQL.String())

		_, err = db.Exec(`
			INSERT INTO waypoints (system,symbol,type,x_coord,y_coord,orbits)
			VALUES (?,?,?,?,?,?)
			ON CONFLICT (symbol) DO UPDATE SET
				system  = EXCLUDED.system,
				symbol  = EXCLUDED.symbol,
				type    = EXCLUDED.type,
				x_coord = EXCLUDED.x_coord,
				y_coord = EXCLUDED.y_coord,
				orbits  = EXCLUDED.orbits
			`,
			s.Data.Symbol,
			s.Data.Waypoints[idx].Symbol,
			s.Data.Waypoints[idx].Type,
			s.Data.Waypoints[idx].Xcoord,
			s.Data.Waypoints[idx].Ycoord,
			s.Data.Waypoints[idx].Orbits,
		)
		if err != nil {
			General.LogErr(fmt.Sprintf("%v", err))
		}
	}

	return nil
}
