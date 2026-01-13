package Task

import "Spacetraders/src/Model"

func GetClientState() Model.GameState {
	var CS = Model.GameState{}
	_ = PG.QueryRow(`SELECT * FROM client`).Scan(
		&CS.IsPaused,
		&CS.ServerTS,
		&CS.AgentTS,
		&CS.FleetTS,
		&CS.ContractTS,
		&CS.SystemsTS,
	)
	return CS
}

func PauseClient(newState bool) error {
	var err error
	var STMT string
	if newState == true {
		STMT = `UPDATE client SET is_paused = 1`
	} else {
		STMT = `UPDATE client SET is_paused = 0`
	}
	_, err = PG.Exec(STMT)
	if err != nil { 
		LogErr("PauseClient: " + err.Error()) 
		return err
	}
	return nil
}
