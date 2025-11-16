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
	var returns strings.Builder

	payload := RegPayload{
		Symbol:  agentSymbol,
		Faction: "VOID",
	}

	jsonBytes, err := json.Marshal(payload)
	if err != nil {
		General.LogErr(err.Error())
	}

	req, _ := http.NewRequest(http.MethodPost, "https://api.spacetraders.io/v2/register", bytes.NewBuffer(jsonBytes))
	if err != nil {
		General.LogErr(err.Error())
	}

	tokenParam := "Bearer " + General.GetToken("account")
	req.Header.Add("Authorization", tokenParam)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Fprintf(&returns, "Error performing request: %v\n", err)
		return returns.String()
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(&returns, "Data Read Failed: %s", err)
		return returns.String()
	}

	fmt.Fprintf(&returns, "%s", body)
	return returns.String()

}
