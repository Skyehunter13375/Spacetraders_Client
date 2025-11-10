package Server

import (
	"Spacetraders/src/General"
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

func GetGameServerState() GameState {
	var g GameState
	var ts string
	db, err := sql.Open("postgres", "user=skyehunter dbname=spacetraders sslmode=disable")
	if err != nil {
		General.LogErr(fmt.Sprintf("DB open failed: %v", err))
	}
	defer db.Close()

	// Check the last update time, if more than 15 mins go grab new info
	err2 := db.QueryRow(`SELECT last_updated FROM server`).Scan(&ts)
	if err2 == sql.ErrNoRows {
		ts = "2025-01-01 13:00:00"
	}

	t, _ := time.ParseInLocation("2006-01-02 15:04:05", ts, time.Local)
	epoch := t.Unix()
	now := time.Now().Unix()
	if (now - epoch) > 900 {
		var sb strings.Builder
		fmt.Fprintf(&sb, "Updating game server status: Over 900 seconds since last refresh (Now:%d - Last:%d)", now, epoch)
		General.LogActivity(sb.String())
		UpdateGameServerState()
	}

	// Once game state is updated we go collect data
	_ = db.QueryRow(
		`SELECT 
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
		return err
	}

	db, err := sql.Open("postgres", "user=skyehunter dbname=spacetraders sslmode=disable")
	if err != nil {
		General.LogErr(fmt.Sprintf("DB open failed: %v", err))
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM server")
	if err != nil {
		return err
	}

	_, err = db.Exec(`
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
		General.LogErr(fmt.Sprintf("%v", err))
		return err
	}

	return nil
}
