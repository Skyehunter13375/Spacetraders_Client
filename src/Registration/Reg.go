package Registration

import (
	"Spacetraders/src/General"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type RegPayload struct {
	Symbol  string `json:"symbol"`
	Faction string `json:"faction"`
}

func RegisterNewAgent(agentSymbol string) string {
	CFG, _ := General.GetConfig()
	var returns strings.Builder

	payload := RegPayload{
		Symbol:  agentSymbol,
		Faction: "VOID",
	}

	jsonBytes, err := json.Marshal(payload)
	if err != nil { General.LogErr("RegisterNewAgent: " + err.Error()) }

	req, _ := http.NewRequest(http.MethodPost, "https://api.spacetraders.io/v2/register", bytes.NewBuffer(jsonBytes))
	if err != nil { General.LogErr("RegisterNewAgent: " + err.Error()) }

	req.Header.Add("Authorization", "Bearer "+CFG.Tokens.Account)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil { General.LogErr("RegisterNewAgent: Error performing request:" + err.Error()); return returns.String() }
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil { General.LogErr("RegisterNewAgent: Data read failed:" + err.Error()); return returns.String() }

	fmt.Fprintf(&returns, "%s", body)
	return returns.String()

}
