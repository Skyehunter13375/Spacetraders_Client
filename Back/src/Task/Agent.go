package Task

import "Spacetraders/src/Model"
import "encoding/json"

func GetAgentState(agent string) Model.Agent {
	var a Model.Agent
	query := `
		SELECT account_id, symbol, faction, hq, ships, credits, last_updated 
		FROM agents 
		WHERE symbol = ?
	`
	_ = PG.QueryRow(query, agent).Scan(
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

func UpdateAgentState(a *Model.Agent) error {
	var err error
	if a == nil {
		jsonStr := GetUrlJson("https://api.spacetraders.io/v2/my/agent", "agent")
		var wrapper map[string]json.RawMessage
		err := json.Unmarshal([]byte(jsonStr), &wrapper)
		if err != nil { LogErr("UpdateAgentState: " + err.Error()) }

		err = json.Unmarshal(wrapper["data"], &a)
		if err != nil { LogErr("UpdateAgentState: " + err.Error()) }
	}

	_, err = PG.Exec(`
		INSERT INTO agents (account_id, symbol, credits, faction, hq, ships, last_updated) 
		VALUES (?,?,?,?,?,?,datetime('now'))
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
	if err != nil { LogErr("UpdateAgentState agents: " + err.Error()); return err }

	return nil
}
