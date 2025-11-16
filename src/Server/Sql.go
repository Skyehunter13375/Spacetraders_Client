package Server

import (
	"Spacetraders/src/General"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func GetGameServerState() GameState {
	var g GameState
	var ts time.Time

	err := General.PG.QueryRow(`SELECT last_updated FROM server`).Scan(&ts)
	if err == sql.ErrNoRows { // No timestamp found, force update
		ts = time.Unix(0, 0).UTC()
	} else if err != nil { // DB error, force update
		General.LogErr(fmt.Sprintf("DB error: %v", err))
		ts = time.Now().UTC().Add(-24 * time.Hour)
	}

	// Compare in UTC only
	if time.Since(ts) > 15*time.Minute {
		General.LogActivity(fmt.Sprintf("Grabbing updated game server status: (Now: %v - Last: %v)", time.Now().UTC(), ts))
		UpdateGameServerState()
	}

	// Pull updated values
	_ = General.PG.QueryRow(`
        SELECT server_up,game_version,agents,ships,systems,waypoints,accounts,
               reset_date,next_reset,reset_freq,last_updated
        FROM server`).Scan(
		&g.Status,
		&g.Version,
		&g.Stats.Agents,
		&g.Stats.Ships,
		&g.Stats.Systems,
		&g.Stats.Waypoints,
		&g.Stats.Accounts,
		&g.LastReset,
		&g.ServerResets.NextReset,
		&g.ServerResets.ResetFreq,
		&g.LastCheckIn,
	)

	return g
}

func UpdateGameServerState() error {
	var g GameState
	jsonStr := General.GetUrlJson("https://api.spacetraders.io/v2", "")
	err := json.Unmarshal([]byte(jsonStr), &g)
	if err != nil {
		General.LogErr(err.Error())
		return err
	}

	_, err = General.PG.Exec("DELETE FROM server")
	if err != nil {
		General.LogErr(err.Error())
		return err
	}

	_, err = General.PG.Exec(`
		INSERT INTO server (
			server_up,
			game_version,
			agents,
			ships,
			systems,
			waypoints,
			accounts,
			reset_date,
			next_reset,
			reset_freq,
			last_updated
		) VALUES (
		 	true, 
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			$8,
			$9,
			NOW()
		)`,
		g.Version,
		g.Stats.Agents,
		g.Stats.Ships,
		g.Stats.Systems,
		g.Stats.Waypoints,
		g.Stats.Accounts,
		g.LastReset,
		g.ServerResets.NextReset,
		g.ServerResets.ResetFreq,
	)
	if err != nil {
		General.LogErr(err.Error())
		return err
	}

	for i := range g.Leaderboards.MostCredits {
		_, err = General.PG.Exec(`
			INSERT INTO leaderboard_creds (agent,credits)
			VALUES ($1,$2)
			ON CONFLICT (agent) DO UPDATE SET
				agent   = EXCLUDED.agent,
				credits = EXCLUDED.credits
			`,
			g.Leaderboards.MostCredits[i].Agent,
			g.Leaderboards.MostCredits[i].Creds,
		)
		if err != nil {
			General.LogErr(err.Error())
		}
	}

	for i := range g.Leaderboards.MostCharted {
		_, err = General.PG.Exec(`
			INSERT INTO leaderboard_charts (agent,charts)
			VALUES ($1,$2)
			ON CONFLICT (agent) DO UPDATE SET
				agent  = EXCLUDED.agent,
				charts = EXCLUDED.charts
			`,
			g.Leaderboards.MostCharted[i].Agent,
			g.Leaderboards.MostCharted[i].Charts,
		)
		if err != nil {
			General.LogErr(err.Error())
		}
	}

	return nil
}
