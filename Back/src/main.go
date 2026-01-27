package main

import "time"
import "Spacetraders/src/Model"
import "Spacetraders/src/Task"

// ┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫ General Funcs ┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
// Functional sleep to prevent reaching request limit on the server
func Tick() {
	// TODO Make this sleep async so that we can still run manual tasks while the automation is paused
	time.Sleep(300 * time.Second)
}

// ┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫ Main ┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
func main() {
	var CurrentState = Model.GameState{}

	// TASK  Test the DB to make sure it's set up correctly
	if err := Task.CheckDB(); err != nil {
		Task.LogErr(err.Error())
		panic(err)
	}

	// TASK Connect to DB and store connection interface globally
	if err := Task.DbLite(); err != nil {
		Task.LogErr(err.Error())
		panic(err)
	}


	for {
		CFG,_ := Task.GetConfig()

		// TASK Register a new agent if the server just reset (Will crash the program if fails)
		// TODO Create new logic for determining if we need a new agent or not.
		if CFG.API.AgentToken == "" {
			Task.LogActivity("No agent token found, registering a new agent")
			err := Task.RegisterNewAgent()
			if err != nil { Task.LogErr("Main: DoPostResetStuff: " + err.Error()); panic(err) }
		}

		CurrentState = Task.GetClientState()

		// TASK If we paused the game, do not execute any scheduled commands.
		if CurrentState.IsPaused == 1 { 
			Task.LogActivity("Game state is paused: Performing only manual tasks")

			// Execute tasks requested from the message table

			// Sleep briefly
			Tick()
			
			// Recheck the client state in case things need to be paused
			CurrentState = Task.GetClientState()
		} else {
			// Execute all scheduled tasks
			Task.LogActivity("Game state NOT paused: Performing all scheduled tasks")

			// Execute tasks requested from the message table
			
			// Sleep briefly
			Tick()
			
			// Recheck the client state in case things need to be paused
			CurrentState = Task.GetClientState()
		}
	}
}
