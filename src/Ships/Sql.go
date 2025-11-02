package Ships

import (
	"Spacetraders/src/General"
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func GetShipState(symbol string) ShipData {
	var sd ShipData
	db, _ := sql.Open("sqlite3", "SpaceTraders.db")
	defer db.Close()

	// Check the last update time, if more than 15 mins go grab new info
	var check_time string
	err2 := db.QueryRow(`SELECT last_updated FROM ship where symbol = 'NULL_SKY-1'`).Scan(&check_time)
	if err2 == sql.ErrNoRows {
		check_time = "2025-01-01 13:00:00"
	}

	t, _ := time.ParseInLocation("2006-01-02 15:04:05", check_time, time.Local)
	epoch := t.Unix()
	now := time.Now().Unix()
	if (now - epoch) > 900 {
		var sb strings.Builder
		fmt.Fprintf(&sb, "Updating ships status: Over 900 seconds since last refresh (Now:%d - Last:%d)", now, epoch)
		General.LogActivity(sb.String())
		UpdateShipState()
	}

	query := `
		SELECT 
			ship.*,
			navg.*,
			crew.*,
			fuel.*,
			frame.*,
			reactor.*,
			engine.*
		FROM ship AS ship
		INNER JOIN ship_nav     AS navg    ON navg.ship    = ship.symbol
		INNER JOIN ship_crew    AS crew    ON crew.ship    = ship.symbol
		INNER JOIN ship_fuel    AS fuel    ON fuel.ship    = ship.symbol
		INNER JOIN ship_frame   AS frame   ON frame.ship   = ship.symbol
		INNER JOIN ship_reactor AS reactor ON reactor.ship = ship.symbol
		INNER JOIN ship_engine  AS engine  ON engine.ship  = ship.symbol
		WHERE ship.symbol = ?
	`
	_ = db.QueryRow(query, symbol).Scan(
		&sd.Symbol,
		&sd.Registration.Name,
		&sd.Registration.Role,
		&sd.Registration.FactionSymbol,
		&sd.LastUpdated,

		&sd.Symbol,
		&sd.Nav.SystemSymbol,
		&sd.Nav.WaypointSymbol,
		&sd.Nav.Status,
		&sd.Nav.FlightMode,
		&sd.Nav.Route.Origin.Symbol,
		&sd.Nav.Route.Origin.Type,
		&sd.Nav.Route.Origin.X,
		&sd.Nav.Route.Origin.Y,
		&sd.Nav.Route.Destination.Symbol,
		&sd.Nav.Route.Destination.Type,
		&sd.Nav.Route.Destination.X,
		&sd.Nav.Route.Destination.Y,
		&sd.Nav.Route.Arrival,
		&sd.Nav.Route.DepartureTime,

		&sd.Symbol,
		&sd.Crew.Current,
		&sd.Crew.Required,
		&sd.Crew.Capacity,
		&sd.Crew.Rotation,
		&sd.Crew.Morale,
		&sd.Crew.Wages,

		&sd.Symbol,
		&sd.Fuel.Current,
		&sd.Fuel.Capacity,

		&sd.Symbol,
		&sd.Frame.Symbol,
		&sd.Frame.Name,
		&sd.Frame.Description,
		&sd.Frame.ModuleSlots,
		&sd.Frame.MountingPoints,
		&sd.Frame.FuelCapacity,
		&sd.Frame.Condition,
		&sd.Frame.Integrity,
		&sd.Frame.Quality,
		&sd.Frame.Requirements.Power,
		&sd.Frame.Requirements.Crew,

		&sd.Symbol,
		&sd.Reactor.Symbol,
		&sd.Reactor.Name,
		&sd.Reactor.Description,
		&sd.Reactor.Condition,
		&sd.Reactor.Integrity,
		&sd.Reactor.PowerOutput,
		&sd.Reactor.Quality,
		&sd.Reactor.Requirements.Crew,

		&sd.Symbol,
		&sd.Engine.Symbol,
		&sd.Engine.Name,
		&sd.Engine.Description,
		&sd.Engine.Condition,
		&sd.Engine.Integrity,
		&sd.Engine.Speed,
		&sd.Engine.Quality,
		&sd.Engine.Requirements.Power,
		&sd.Engine.Requirements.Crew,
	)

	//? Debug: Log the entire ship struct to JSON in file
	// jsonData, _ := json.Marshal(sd)
	// General.LogActivity(string(jsonData))

	return sd
}

func UpdateShipState() error {
	data := General.GetUrlJson("https://api.spacetraders.io/v2/my/ships", "agent")

	var wrapper map[string]json.RawMessage
	json.Unmarshal([]byte(data), &wrapper)

	var ships []ShipData
	json.Unmarshal(wrapper["data"], &ships)

	db, _ := sql.Open("sqlite3", "SpaceTraders.db")
	defer db.Close()

	for _, s := range ships {
		// ┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫ Upsert Ship ┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
		_, err := db.Exec(`
			INSERT INTO ship (
				symbol,
				name,
				role,
				faction,
				last_updated
			) VALUES (
				?,
				?,
				?,
				?,
				datetime('now', 'localtime')
			)
			ON CONFLICT (symbol) DO UPDATE SET
				symbol  = EXCLUDED.symbol,
				name    = EXCLUDED.name,
				role    = EXCLUDED.role,
				faction = EXCLUDED.faction,
				last_updated = datetime('now', 'localtime')
			`,
			s.Symbol,
			s.Registration.Name,
			s.Registration.Role,
			s.Registration.FactionSymbol,
		)
		if err != nil {
			return err
		}

		// ┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫ Upsert Ship Nav ┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
		_, err = db.Exec(`
			INSERT INTO ship_nav (
				ship,
				system,
				waypoint,
				status,
				flight_mode,
				origin,
				origin_type,
				origin_x,
				origin_y,
				destination,
				destination_type,
				destination_x,
				destination_y,
				arrival,
				departure
			) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)
			ON CONFLICT (ship) DO UPDATE SET
				ship             = EXCLUDED.ship,
				system           = EXCLUDED.system,
				waypoint         = EXCLUDED.waypoint,
				status           = EXCLUDED.status,
				flight_mode      = EXCLUDED.flight_mode,
				origin           = EXCLUDED.origin,
				origin_type      = EXCLUDED.origin_type,
				origin_x         = EXCLUDED.origin_x,
				origin_y         = EXCLUDED.origin_y,
				destination      = EXCLUDED.destination,
				destination_type = EXCLUDED.destination_type,
				destination_x    = EXCLUDED.destination_x,
				destination_y    = EXCLUDED.destination_y,
				arrival          = EXCLUDED.arrival,
				departure        = EXCLUDED.departure
			`,
			s.Symbol,
			s.Nav.SystemSymbol,
			s.Nav.WaypointSymbol,
			s.Nav.Status,
			s.Nav.FlightMode,
			s.Nav.Route.Origin.Symbol,
			s.Nav.Route.Origin.Type,
			s.Nav.Route.Origin.X,
			s.Nav.Route.Origin.Y,
			s.Nav.Route.Destination.Symbol,
			s.Nav.Route.Destination.Type,
			s.Nav.Route.Destination.X,
			s.Nav.Route.Destination.Y,
			s.Nav.Route.Arrival,
			s.Nav.Route.DepartureTime,
		)
		if err != nil {
			return err
		}

		// ┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫ Upsert Ship Crew ┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
		_, err = db.Exec(`
			INSERT INTO 
				ship_crew (ship, current, required, capacity, rotation, morale, wages) 
				VALUES (?,?,?,?,?,?,?)
			ON CONFLICT (ship) DO UPDATE SET
				ship      = EXCLUDED.ship,
				current   = EXCLUDED.current, 
				required  = EXCLUDED.required, 
				capacity  = EXCLUDED.capacity,
				rotation  = EXCLUDED.rotation, 
				morale    = EXCLUDED.morale, 
				wages     = EXCLUDED.wages
			`,
			s.Symbol,
			s.Crew.Current,
			s.Crew.Required,
			s.Crew.Capacity,
			s.Crew.Rotation,
			s.Crew.Morale,
			s.Crew.Wages,
		)
		if err != nil {
			return err
		}

		// ┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫ Upsert Ship Fuel ┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
		_, err = db.Exec(`
			INSERT INTO 
				ship_fuel (ship,current,capacity)
				VALUES (?,?,?)
			ON CONFLICT (ship) DO UPDATE SET
				ship     = EXCLUDED.ship,
				current  = EXCLUDED.current,
				capacity = EXCLUDED.capacity
		`,
			s.Symbol,
			s.Fuel.Current,
			s.Fuel.Capacity,
		)
		if err != nil {
			return err
		}

		// ┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫ Upsert Ship Frame ┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
		_, err = db.Exec(`
			INSERT INTO 
				ship_frame (ship,symbol,name,description,module_slots,mount_points,fuel_capacity,condition,integrity,quality,power_required,crew_required)
				VALUES (?,?,?,?,?,?,?,?,?,?,?,?)
			ON CONFLICT (ship) DO UPDATE SET
				ship           = EXCLUDED.ship,
				symbol         = EXCLUDED.symbol,
				name           = EXCLUDED.name,
				description    = EXCLUDED.description,
				module_slots   = EXCLUDED.module_slots,
				mount_points   = EXCLUDED.mount_points,
				fuel_capacity  = EXCLUDED.fuel_capacity,
				condition      = EXCLUDED.condition,
				integrity      = EXCLUDED.integrity,
				quality        = EXCLUDED.quality,
				power_required = EXCLUDED.power_required,
				crew_required  = EXCLUDED.crew_required
		`,
			s.Symbol,
			s.Frame.Symbol,
			s.Frame.Name,
			s.Frame.Description,
			s.Frame.ModuleSlots,
			s.Frame.MountingPoints,
			s.Frame.FuelCapacity,
			s.Frame.Condition,
			s.Frame.Integrity,
			s.Frame.Quality,
			s.Frame.Requirements.Power,
			s.Frame.Requirements.Crew,
		)
		if err != nil {
			return err
		}

		// ┣━━━━━━━━━━━━━━━━━━━━━━━━━━━┫ Upsert Ship Reactor ┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
		_, err = db.Exec(`
			INSERT INTO 
				ship_reactor (ship,symbol,name,description,condition,integrity,power_output,quality,crew_required)
				VALUES (?,?,?,?,?,?,?,?,?)
			ON CONFLICT (ship) DO UPDATE SET
				ship           = EXCLUDED.ship,
				symbol         = EXCLUDED.symbol,
				name           = EXCLUDED.name,
				description    = EXCLUDED.description,
				condition      = EXCLUDED.condition,
				integrity      = EXCLUDED.integrity,
				power_output   = EXCLUDED.power_output,
				quality        = EXCLUDED.quality,
				crew_required  = EXCLUDED.crew_required
		`,
			s.Symbol,
			s.Reactor.Symbol,
			s.Reactor.Name,
			s.Frame.Description,
			s.Frame.Condition,
			s.Frame.Integrity,
			s.Frame.Quality,
			s.Frame.Requirements.Power,
			s.Frame.Requirements.Crew,
		)
		if err != nil {
			return err
		}

		// ┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫ Upsert Ship Engine ┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
		_, err = db.Exec(`
			INSERT INTO 
				ship_engine (ship,symbol,name,description,condition,integrity,speed,quality,power_required,crew_required)
				VALUES (?,?,?,?,?,?,?,?,?,?)
			ON CONFLICT (ship) DO UPDATE SET
				ship           = EXCLUDED.ship,
				symbol         = EXCLUDED.symbol,
				name           = EXCLUDED.name,
				description    = EXCLUDED.description,
				condition      = EXCLUDED.condition,
				integrity      = EXCLUDED.integrity,
				speed          = EXCLUDED.speed,
				quality        = EXCLUDED.quality,
				power_required = EXCLUDED.power_required,
				crew_required  = EXCLUDED.crew_required
		`,
			s.Symbol,
			s.Engine.Symbol,
			s.Engine.Name,
			s.Engine.Description,
			s.Engine.Condition,
			s.Engine.Integrity,
			s.Engine.Speed,
			s.Engine.Quality,
			s.Engine.Requirements.Power,
			s.Engine.Requirements.Crew,
		)
		if err != nil {
			return err
		}
	}

	return nil
}
