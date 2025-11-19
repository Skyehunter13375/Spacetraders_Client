package General

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func GetUrlJson(url string, tokenType string) string {
	CFG, _ := GetConfig()
	var returns strings.Builder
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Accept", "application/json")

	switch tokenType {
		case "agent":
			req.Header.Add("Authorization", "Bearer "+CFG.Tokens.Agent)
		case "account":
			req.Header.Add("Authorization", "Bearer "+CFG.Tokens.Account)
		default:
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil { LogErr("GetUrlJson: Connection " + err.Error()); return ""}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil { LogErr("GetUrlJson: Data Read Failed" + err.Error()); return ""}

	fmt.Fprintf(&returns, "%s", body)
	return returns.String()
}

func PostUrlJson(url string, tokenType string) (string, error) {
	CFG, _ := GetConfig()
	var returns strings.Builder

	req, _ := http.NewRequest(http.MethodPost, url, nil)
	req.Header.Set("Accept", "application/json")

	switch tokenType {
		case "agent":
			req.Header.Add("Authorization", "Bearer "+CFG.Tokens.Agent)
		case "account":
			req.Header.Add("Authorization", "Bearer "+CFG.Tokens.Account)
		default:
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)

	if err != nil { LogErr("GetUrlJson: Error performing request" + err.Error()); return "", err}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil { LogErr("GetUrlJson: Data Read Failed" + err.Error()); return "", err}

	fmt.Fprintf(&returns, "%s", body)
	return returns.String(), nil
}
