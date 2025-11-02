package General

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func GetUrlJson(url string, token string) string {
	var returns strings.Builder
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Accept", "application/json")

	switch token {
	case "agent":
		tokenParam := "Bearer " + GetToken("agent")
		req.Header.Add("Authorization", tokenParam)

	case "account":
		tokenParam := "Bearer " + GetToken("account")
		req.Header.Add("Authorization", tokenParam)

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
