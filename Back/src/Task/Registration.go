package Task

import "io"
import "encoding/json"
import "bytes"
import "net/http"
import "time"
import "Spacetraders/src/Model"

func RegisterNewAgent() error {
	CFG, _ := GetConfig()

	// TASK Web API Request
	payload := Model.RegPayload{
		Symbol:  CFG.API.AgentName,
		Faction: CFG.API.AgentFaction,
	}

	jsonBytes, err := json.Marshal(payload)
	if err != nil { LogErr("RegisterNewAgent: Failed to marshal the json payload: " + err.Error()) }

	req, _ := http.NewRequest(http.MethodPost, "https://api.spacetraders.io/v2/register", bytes.NewBuffer(jsonBytes))
	if err != nil { LogErr("RegisterNewAgent: Failed to create request: " + err.Error()) }

	req.Header.Add("Authorization", "Bearer "+CFG.API.AccntToken)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil { LogErr("RegisterNewAgent: Failed to connect to API: " + err.Error()); return err }
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil { LogErr("RegisterNewAgent: Failed to read GET body:" + err.Error()); return err }

	// TASK Unwrap the json for ingest
	var wrapper map[string]json.RawMessage
	err = json.Unmarshal([]byte(body), &wrapper)
	if err != nil { LogErr("RegisterNewAgent: Unmarshalling data: " + err.Error()) }

	var r Model.RegResult
	err = json.Unmarshal(wrapper["data"], &r)
	if err != nil { LogErr("RegisterNewAgent: Unmarshalling data: " + err.Error()) }

	// DEBG Some lovely printf debugging
	LogActivity("NewRegistration: " + string(body))
	LogActivity("NewAgentToken: " + r.Token)
	LogActivity("Starting System: " + r.Ships[0].Nav.SystemSymbol)
	
	// TASK Update the config file with the new agent token automagically
	CFG.API.AgentToken = r.Token
	err = SaveConfig(CFG)
	if err != nil { LogErr("RegisterNewAgent: Store new token: " + err.Error()) }

	// TASK Reset the DB, archiving the old one
	err = ResetDB()
	if err != nil { LogErr("RegisterNewAgent: Resetting DB: " + err.Error()) }

	// TASK Set client defaults table
	_, err = PG.Exec(`INSERT INTO client (is_paused,server_updated,agent_updated,fleet_updated,contracts_updated,systems_updated) VALUES (0, datetime('now'), datetime('now'), datetime('now'), datetime('now'), datetime('now'))` )
	if err != nil { LogErr("RegisterNewAgent: Defaulting client table: " + err.Error()) }

	// TASK Insert the server status data
	err = UpdateGameServerState()
	if err != nil { LogErr("RegisterNewAgent: Ingesting Server Data: " + err.Error()) }

	// TASK Insert agent data
	err = UpdateAgentState(&r.Agent)
	if err != nil { LogErr("RegisterNewAgent: Ingesting Agent Data: " + err.Error()) }

	// TASK Insert contract data
	err = UpdateContracts([]Model.Contract{r.Contract})
	if err != nil { LogErr("RegisterNewAgent: Ingesting Contract Data: " + err.Error()) }

	// TASK Insert faction data
	err = UpdateFaction(&r.Faction)
	if err != nil { LogErr("RegisterNewAgent: Ingesting Faction Data: " + err.Error()) }

	// TASK Insert fleet data
	err = UpdateShipState(r.Ships)
	if err != nil { LogErr("RegisterNewAgent: Ingesting Fleet Data: " + err.Error()) }

	// TASK Insert starting system data (This ONLY includes large celestial bodies, not asteroids)
	err = UpdateSystem(r.Ships[0].Nav.SystemSymbol)
	if err != nil { LogErr("RegisterNewAgent: Ingesting System Data: " + err.Error()) }

	return nil
}

