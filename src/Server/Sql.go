package Server

import "Spacetraders/src/General"
import "database/sql"
import "encoding/json"
import "fmt"
import "time"

func UpdateGameServerState() error {
	var g GameState
	jsonStr := General.GetUrlJson("https://api.spacetraders.io/v2", "")
	err := json.Unmarshal([]byte(jsonStr), &g)
	if err != nil { General.LogErr("UpdateGameServerState: " + err.Error()); return err }

	_, err = General.PG.Exec("DELETE FROM server")
	if err != nil { General.LogErr("UpdateGameServerState: Delete failed" + err.Error()); return err }

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
	if err != nil { General.LogErr("UpdateGameServerState: Insert server failed" + err.Error()); return err }

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
	if err != nil { General.LogErr("UpdateGameServerState: Insert leadCreds failed" + err.Error()); return err }
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
	if err != nil { General.LogErr("UpdateGameServerState: Insert leadCharts failed" + err.Error()); return err }
	}

	return nil
}


func GetGameServerState() GameState {
	var g GameState
	var ts time.Time

	err := General.PG.QueryRow(`SELECT last_updated FROM server`).Scan(&ts)
	if err == sql.ErrNoRows { // No timestamp found, force update
		ts = time.Unix(0, 0).UTC()
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
