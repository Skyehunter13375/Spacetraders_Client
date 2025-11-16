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

	if err != nil {
		LogErr(fmt.Sprintf("Error performing request: %v\n", err))
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		LogErr(fmt.Sprintf("Data Read Failed: %v\n", err))
		return "", err
	}

	fmt.Fprintf(&returns, "%s", body)
	return returns.String(), nil
}
