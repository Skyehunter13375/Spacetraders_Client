package General

import "database/sql"
import _ "github.com/mattn/go-sqlite3"
import "os"

var PG *sql.DB

// FEAT: Check if DB file exists, if not create it
func CheckDB() error {
    CFG, _ := GetConfig()
    _, err := os.Stat(CFG.DB.DbPath)
    if err == nil { return nil }

    if !os.IsNotExist(err) {
        // Some other filesystem error
        LogErr("DB: Stat failed: " + err.Error())
        return err
    }

    // DB does not exist — create it by opening SQLite
    db, err := sql.Open("sqlite3", CFG.DB.DbPath)
    if err != nil {
        LogErr("DB: Failed creating SQLite DB: " + err.Error())
        return err
    }
    defer db.Close()

    // Read schema/setup file
    schema, err := os.ReadFile(CFG.DB.DbBuild)
    if err != nil {
        LogErr("DB: Failed reading setup file: " + err.Error())
        return err
    }

    // Execute setup SQL
    _, err = db.Exec(string(schema))
    if err != nil {
        LogErr("DB: Failed executing setup SQL: " + err.Error())
        return err
    }

    LogActivity("DB: Created new SQLite database and applied schema.")
    return nil
}

// FEAT: Connect to SQLite database
func DbLite() error {
	var err error
	CFG, _ := GetConfig()
	PG, err = sql.Open("sqlite3", CFG.DB.DbPath)
	if err != nil {
		LogErr("DB: Connection failed: " + err.Error());
		return err
	}
	return PG.Ping()
}
