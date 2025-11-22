package Registration

import "io"
import "encoding/json"
import "bytes"
import "net/http"
import "time"
import "Spacetraders/src/General"

type RegPayload struct {
	Symbol  string `json:"symbol"`
	Faction string `json:"faction"`
}

func RegisterNewAgent(agentSymbol string, faction string) (string, error) {
	CFG, _ := General.GetConfig()

	payload := RegPayload{
		Symbol:  agentSymbol,
		Faction: faction,
	}

	jsonBytes, err := json.Marshal(payload)
	if err != nil { General.LogErr("RegisterNewAgent: " + err.Error()) }

	req, _ := http.NewRequest(http.MethodPost, "https://api.spacetraders.io/v2/register", bytes.NewBuffer(jsonBytes))
	if err != nil { General.LogErr("RegisterNewAgent: " + err.Error()) }

	req.Header.Add("Authorization", "Bearer "+CFG.Tokens.Account)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil { General.LogErr("RegisterNewAgent: Error performing request:" + err.Error()); return "",err }
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil { General.LogErr("RegisterNewAgent: Data read failed:" + err.Error()); return "", err }

	// INFO: Begin unwrapping the json for ingest
	var wrapper map[string]json.RawMessage
	json.Unmarshal([]byte(body), &wrapper)

	var data map[string]json.RawMessage
	json.Unmarshal(wrapper["data"], &data)

	var token string
	json.Unmarshal(data["token"], &token)

	// TODO: Run the spacetraders_reset.sql file? Or do it manually here? 

	// TODO: Figure out an easier way to update config file with new agent token automagically

	return token, nil
}
