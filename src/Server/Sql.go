package Server

import "Spacetraders/src/General"
import "encoding/json"
import "time"

// FEAT: Capture and store the state of the game
func UpdateGameServerState() error {
	var g GameState
	jsonStr := General.GetUrlJson("https://api.spacetraders.io/v2", "")

	err := json.Unmarshal([]byte(jsonStr), &g)
	if err != nil { General.LogErr("UpdateGameServerState: " + err.Error()); return err }

	_, err = General.PG.Exec("DELETE FROM server")
	if err != nil { General.LogErr("UpdateGameServerState: Delete failed" + err.Error()); return err }

	// TASK: Capture and store server state
	STMT1 := `
		INSERT INTO server (server_up,game_version,agents,ships,systems,waypoints,accounts,reset_date,next_reset,reset_freq,last_updated) 
		VALUES (1,?,?,?,?,?,?,?,?,?,strftime('%Y-%m-%dT%H:%M:%fZ','now'))
	`
	_, err = General.PG.Exec(STMT1, g.Version, g.Stats.Agents, g.Stats.Ships, g.Stats.Systems, g.Stats.Waypoints, g.Stats.Accounts, g.LastReset, g.ServerResets.NextReset, g.ServerResets.ResetFreq)
	if err != nil { General.LogErr("UpdateGameServerState: Insert server failed: " + STMT1 + err.Error()); return err }

	// TASK: Capture and store current leaderboard (credits)
	STMT2 := `
		INSERT INTO leaderboard_creds (agent,credits)
		VALUES (?,?)
		ON CONFLICT (agent) DO UPDATE SET
			agent   = EXCLUDED.agent,
			credits = EXCLUDED.credits
	`
	for i := range g.Leaderboards.MostCredits {
		_, err = General.PG.Exec(STMT2, g.Leaderboards.MostCredits[i].Agent, g.Leaderboards.MostCredits[i].Creds )
		if err != nil { General.LogErr("UpdateGameServerState: Insert leadCreds failed:" + STMT2 + err.Error()); return err }
	}

	// TASK: Capture and store current leaderboard (charts)
	STMT3 := `
		INSERT INTO leaderboard_charts (agent,charts)
		VALUES (?,?)
		ON CONFLICT (agent) DO UPDATE SET
			agent  = EXCLUDED.agent,
			charts = EXCLUDED.charts
	`
	for i := range g.Leaderboards.MostCharted {
		_, err = General.PG.Exec(STMT3, g.Leaderboards.MostCharted[i].Agent, g.Leaderboards.MostCharted[i].Charts)
		if err != nil { General.LogErr("UpdateGameServerState: Insert leadCharts failed:" + STMT3 + err.Error()); return err }
	}

	return nil
}

// FEAT: Get game state from stored SQLite data
func GetGameServerState() GameState {
	var g GameState
	
	// TASK: Check last updated timestamp, if > 15 mins go pull new data
	tsStr := "1970-01-01T00:00:00Z"
	General.PG.QueryRow(`SELECT last_updated FROM server`).Scan(&tsStr)
	ts, _ := time.Parse(time.RFC3339, tsStr)
	if time.Since(ts) > 15*time.Minute { UpdateGameServerState() }

	// TASK: Pull updated values
	_ = General.PG.QueryRow(`SELECT server_up,game_version,agents,ships,systems,waypoints,accounts,reset_date,next_reset,reset_freq,last_updated FROM server`).Scan(
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

// FEAT: Get leaderboard state from SQLite data
func GetLeaderboards() Leaderboards {
	var result Leaderboards

	data,err := General.PG.Query(`SELECT agent,credits FROM leaderboard_creds ORDER BY credits DESC LIMIT 10`)
	if err != nil { General.LogErr("GetLeaderboards: " + err.Error()) } 

	cnt := 0
	for data.Next() {
		var temp LeaderCredits
		data.Scan(&temp.Agent, &temp.Creds)
		result.MostCredits = append(result.MostCredits, temp)
		cnt++
	}
	
	data2,err := General.PG.Query(`SELECT agent,charts FROM leaderboard_charts ORDER BY charts DESC LIMIT 10`)
	if err != nil { General.LogErr("GetLeaderboards: " + err.Error()) } 

	cnt = 0
	for data2.Next() {
		var temp2 LeaderCharts
		data2.Scan(&temp2.Agent, &temp2.Charts)
		result.MostCharted = append(result.MostCharted, temp2)
		cnt++
	}

	return result
}
