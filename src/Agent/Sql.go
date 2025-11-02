package Agent

import (
	"Spacetraders/src/General"
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func GetAgentState(agent string) AgentState {
	var a AgentState
	db, _ := sql.Open("sqlite3", "src/DB/spacetraders.db")
	defer db.Close()

	// Check the last update time, if more than 15 mins go grab new info
	var check_time string
	err2 := db.QueryRow(`SELECT last_updated FROM agents where symbol = ?`, agent).Scan(&check_time)
	if err2 == sql.ErrNoRows {
		check_time = "2025-01-01 13:00:00"
	}

	t, _ := time.ParseInLocation("2006-01-02 15:04:05", check_time, time.Local)
	epoch := t.Unix()
	now := time.Now().Unix()
	if (now - epoch) > 900 {
		var sb strings.Builder
		fmt.Fprintf(&sb, "Updating agent %s status: Over 900 seconds since last refresh (Now:%d - Last:%d)", agent, now, epoch)
		General.LogActivity(sb.String())
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
		WHERE symbol = ?`

	_ = db.QueryRow(query, agent).Scan(
		&a.Data.AccountID,
		&a.Data.Symbol,
		&a.Data.Faction,
		&a.Data.HQ,
		&a.Data.Ships,
		&a.Data.Credits,
		&a.Data.LastUpdated,
	)

	return a
}

func UpdateAgentState() error {
	var a AgentState

	jsonStr := General.GetUrlJson("https://api.spacetraders.io/v2/my/agent", "agent")
	err := json.Unmarshal([]byte(jsonStr), &a)
	if err != nil {
		return err
	}

	db, err := sql.Open("sqlite3", "src/DB/spacetraders.db")
	defer db.Close()

	_, err = db.Exec(`
		INSERT INTO agents (
			account_id,
			symbol,
			credits,
			faction,
			hq,
			ships,
			last_updated
		) VALUES (
			?,
			?,
			?,
			?,
			?,
			?,
			datetime('now', 'localtime')
		)
		ON CONFLICT (symbol) DO UPDATE SET
			account_id   = EXCLUDED.account_id,
			symbol       = EXCLUDED.symbol,
			hq           = EXCLUDED.hq,
			credits      = EXCLUDED.credits,
			ships        = EXCLUDED.ships,
			faction      = EXCLUDED.faction,
			last_updated = datetime('now', 'localtime')
		`,
		a.Data.AccountID,
		a.Data.Symbol,
		a.Data.Credits,
		a.Data.Faction,
		a.Data.HQ,
		a.Data.Ships,
	)
	if err != nil {
		return err
	}

	return nil
}
