package main

import "time"
import "Spacetraders/src/Model"
import "Spacetraders/src/Task"

// ┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫ General Funcs ┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
// Functional sleep to prevent reaching request limit on the server
func Tick() {
	time.Sleep(10 * time.Second)
}

// ┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫ Main ┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
func main() {
	var CurrentState = Model.GameState{}
	
	// FEAT: Test the DB to make sure it's set up correctly
	if err := Task.CheckDB(); err != nil {
		Task.LogErr(err.Error())
		panic(err)
	}

	// FEAT: Connect to DB and store connection interface globally
	if err := Task.DbLite(); err != nil {
		Task.LogErr(err.Error())
		panic(err)
	}

	// FEAT: Register new agent if needed - Post Reset
	CFG,_ := Task.GetConfig()
	if CFG.API.AgentToken == "" {
		err := Task.RegisterNewAgent()
		if err != nil { Task.LogErr("Main: RegisterNewAgent: " + err.Error()); panic(err) }
	}
	
	// TASK: Create function here to keep the process running and monitoring all ships in the fleet
	for {
		CurrentState = Task.GetClientState()

		// If we paused the game, do not execute any scheduled commands.
		if CurrentState.IsPaused == 1 { 
			Task.LogActivity("Game state is paused: Performing only manual tasks")

			// Execute tasks requested from the message table

			// Sleep briefly
			Tick()
			
			// Recheck the client state in case things need to be paused
			CurrentState = Task.GetClientState()
		} else {
			// Execute all scheduled tasks

			// Execute tasks requested from the message table
			
			// Sleep briefly
			Tick()
			
			// Recheck the client state in case things need to be paused
			CurrentState = Task.GetClientState()
		}
	}
}
