package General

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func GetToken(tokentype string) string {
	var token string
	db, _ := sql.Open("sqlite3", "src/DB/spacetraders.db")
	defer db.Close()
	err := db.QueryRow(`SELECT token FROM tokens WHERE type = ? LIMIT 1`, tokentype).Scan(&token)
	if err == sql.ErrNoRows {
		return ""
	}

	return string(token)
}
