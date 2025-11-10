package General

import (
	"database/sql"
	"fmt"
)

func GetToken(tokentype string) string {
	var token string
	db, err := sql.Open("postgres", "user=skyehunter dbname=spacetraders sslmode=disable")
	if err != nil {
		LogErr(fmt.Sprintf("DB open failed: %v", err))
	}
	defer db.Close()

	err = db.QueryRow(`SELECT token FROM tokens WHERE type = $1 LIMIT 1`, tokentype).Scan(&token)
	if err == sql.ErrNoRows {
		return ""
	}

	return string(token)
}
