package Task

import "database/sql"
import "fmt"
import "os"
import "time"
import _ "github.com/mattn/go-sqlite3"

var PG *sql.DB

// FEAT: Connect to SQLite database
func DbLite() error {
	var err error
	CFG, _ := GetConfig()
	PG, err = sql.Open("sqlite3", CFG.DB.DbPath)
	if err != nil {
		LogErr("DB: Connection failed: " + err.Error());
		return err
	}

	PG.Exec("PRAGMA foreign_keys = ON")
	PG.Exec("PRAGMA detailed_errors = ON")
	return PG.Ping()
}

// FEAT: Check if DB file exists, if not create it
func CheckDB() error {
    CFG, _ := GetConfig()
    _, err := os.Stat(CFG.DB.DbPath)
    if err == nil { return nil }

    if !os.IsNotExist(err) {
        LogErr("DB: Stat failed: " + err.Error())
		fmt.Println("Database not found: Creating " + CFG.DB.DbPath + " now...")
        return err
    }

	// TASK: Connect to the DB
	// If the DB doesn't exist this will create it
    db, err := sql.Open("sqlite3", CFG.DB.DbPath)
    if err != nil {
        LogErr("DB: Failed creating SQLite DB: " + err.Error())
		fmt.Println("Database could not be created...")
        return err
    }
    defer db.Close()

	// TASK: Get table and key info from setup file
	fmt.Println("Reading in schema setup from " + CFG.DB.DbBuild)
    schema, err := os.ReadFile(CFG.DB.DbBuild)
    if err != nil {
        LogErr("DB: Failed reading setup file: " + err.Error())
        return err
    }

	// TASK: Create tables & Foreign Keys based on setup file
    _, err = db.Exec(string(schema))
    if err != nil {
        LogErr("DB: Failed executing setup SQL: " + err.Error())
        return err
    }

    LogActivity("DB: Created new SQLite database and applied schema.")
    return nil
}


// FEAT: Reset the database after each server wipe
// Archive the previous one in case we want to look back
func ResetDB() error {
    CFG, _ := GetConfig()

	// TASK: Disconnect from existing DB file
	if PG != nil {
		if err := PG.Close(); err != nil { return err }
		PG = nil
	}

	// TASK: Archive the current DB
	if _, err := os.Stat(CFG.DB.DbPath); err == nil {
		archive := CFG.DB.DbPath + "_" + time.Now().Format("01_02_2006")
		if err := os.Rename(CFG.DB.DbPath, archive); err != nil {
			return err
		}
	}

	// TASK: Create the new DB by connecting to it.
	if err := DbLite(); err != nil {
		LogErr(err.Error())
		return err
	}

	// TASK: Get table and key info from setup file
	fmt.Println("Reading in schema setup from " + CFG.DB.DbBuild)
    schema, err := os.ReadFile(CFG.DB.DbBuild)
    if err != nil {
        LogErr("DB: Failed reading setup file: " + err.Error())
        return err
    }

	// TASK: Create tables & Foreign Keys based on setup file
    _, err = PG.Exec(string(schema))
    if err != nil {
        LogErr("DB: Failed executing setup SQL: " + err.Error())
        return err
    }

    LogActivity("DB: Created new SQLite database and applied schema.")
    return nil
}
