package Task

import "io"
import "encoding/json"
import "bytes"
import "net/http"
import "time"
import "Spacetraders/src/Model"

func RegisterNewAgent() error {
	CFG, _ := GetConfig()

	// TASK: Web API Request
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

	// TASK: Unwrap the json for ingest
	var wrapper map[string]json.RawMessage
	err = json.Unmarshal([]byte(body), &wrapper)
	if err != nil { LogErr("RegisterNewAgent: Unmarshalling data: " + err.Error()) }

	var r Model.RegResult
	err = json.Unmarshal(wrapper["data"], &r)
	if err != nil { LogErr("RegisterNewAgent: Unmarshalling data: " + err.Error()) }

	// DEBG:
	LogActivity("NewRegistration: " + string(body))
	LogActivity("NewAgentToken: " + r.Token)
	
	// TASK: Figure out an easier way to update config file with new agent token automagically
	CFG.API.AgentToken = r.Token
	err = SaveConfig(CFG)
	if err != nil { LogErr("RegisterNewAgent: Store new token: " + err.Error()) }

	// TASK: Reset the DB, archiving the old one if it exists
	err = ResetDB()
	if err != nil { LogErr("RegisterNewAgent: Resetting DB: " + err.Error()) }

	// TASK: Insert agent data
	err = UpdateAgentState(&r.Agent)
	if err != nil { LogErr("RegisterNewAgent: Ingesting Agent Data: " + err.Error()) }

	// TASK: Insert contract data
	err = UpdateContracts([]Model.Contract{r.Contract})
	if err != nil { LogErr("RegisterNewAgent: Ingesting Contract Data: " + err.Error()) }

	// TASK: Insert faction data
	err = UpdateFaction(&r.Faction)
	if err != nil { LogErr("RegisterNewAgent: Ingesting Faction Data: " + err.Error()) }

	// TASK: Capture the data for the system this agent started in
	err = UpdateSystem(r.Ships[0].Nav.SystemSymbol)
	if err != nil { LogErr("RegisterNewAgent: Ingesting Faction Data: " + err.Error()) }

	return nil
}

