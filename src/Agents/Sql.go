package Agents

import "Spacetraders/src/General"
import "database/sql"
import "encoding/json"
import "fmt"
import "time"

func GetAgentState(agent string) Agent {
	var a Agent

	var ts time.Time
	err := General.PG.QueryRow(`SELECT last_updated FROM agents where symbol = $1`, agent).Scan(&ts)
	if err == sql.ErrNoRows { // No timestamp found, force update
		ts = time.Unix(0, 0).UTC()
	}

	// Compare in UTC only
	if time.Since(ts) > 15*time.Minute {
		General.LogActivity(fmt.Sprintf("Updating agent status: (Now: %v - Last: %v)", time.Now().UTC(), ts))
		UpdateAgentState()
	}

	query := `
		SELECT 
			account_id,
			symbol,
			faction,
			hq,
			ships,
			credits,
			last_updated 
		FROM agents 
		WHERE symbol = $1`

	_ = General.PG.QueryRow(query, agent).Scan(
		&a.AccountID,
		&a.Symbol,
		&a.Faction,
		&a.HQ,
		&a.Ships,
		&a.Credits,
		&a.LastUpdated,
	)

	return a
}

func UpdateAgentState() error {
	jsonStr := General.GetUrlJson("https://api.spacetraders.io/v2/my/agent", "agent")
	var wrapper map[string]json.RawMessage
	err := json.Unmarshal([]byte(jsonStr), &wrapper)
	if err != nil { General.LogErr("UpdateAgentState: " + err.Error()) }

	var a Agent
	err = json.Unmarshal(wrapper["data"], &a)
	if err != nil { General.LogErr("UpdateAgentState: " + err.Error()) }


	_, err = General.PG.Exec(`
		INSERT INTO agents (
			account_id,
			symbol,
			credits,
			faction,
			hq,
			ships,
			last_updated
		) VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			NOW()
		)
		ON CONFLICT (symbol) DO UPDATE SET
			account_id   = EXCLUDED.account_id,
			symbol       = EXCLUDED.symbol,
			hq           = EXCLUDED.hq,
			credits      = EXCLUDED.credits,
			ships        = EXCLUDED.ships,
			faction      = EXCLUDED.faction,
			last_updated = EXCLUDED.last_updated
		`,
		a.AccountID,
		a.Symbol,
		a.Credits,
		a.Faction,
		a.HQ,
		a.Ships,
	)
	if err != nil { General.LogErr("UpdateAgentState agents: " + err.Error()); return err }

	return nil
}
