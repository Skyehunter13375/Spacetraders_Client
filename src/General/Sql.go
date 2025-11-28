package General

import "database/sql"
import "fmt"
import _ "github.com/lib/pq"

var PG *sql.DB

func DB() error {
	CFG, _ := GetConfig()
	connStr := fmt.Sprintf("user=%s dbname=%s sslmode=%s", CFG.DB.User, CFG.DB.Name, CFG.DB.SSL)

	var err error
	PG, err = sql.Open(CFG.DB.Type, connStr)
	if err != nil { LogErr("DB: Connection failed: " + err.Error()); return err }

	return PG.Ping()
}
