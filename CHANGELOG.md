## [0.3.0] - Completed: 11/16/2025
- Another big refactor while I am learning which packages I like
- Filling in gaps in data collection routines
#### 11/16/2025
- Complete overhaul of postgres connections. Simplified and trimmed lots of fat.
- Added a config file to standardize DB connection info and account/agent tokens.
- Fixed all the timestamp checking (UTC vs Local time) and rebuilt the database.
- Abandoned Bubbletea in favor of Tview. Much simpler, easier to understand and less abstract.
- Shifting gears from data collection to TUI building then to start actually playing the game.
- Also added a quick feature milestone tracker to the README. Swiped that idea from someone else's project.
#### 11/13/2025
- New agent registration function so I don't have to keep going to the web page on reset. Requires testing still.
- Basic leaderboard data collection started.
#### 11/11/2025
- Created structs and function template for shipyard data collection but can't test until next reset or I move systems.

## [0.2.5] - Completed: 11/10/2025
- Finishing out initial stages of data collection
#### 11/10/2025
- Fixed my ingest routines to properly use json.Unmarshal() with wrappers to get rid of the "data" fields
- Rebuilt my structs for all data types for ease of using interfaces and SQL tuning later on
- Gave up on SQLite and switched back to PostgreSQL. I need the ability to store arrays in a single column for cross reference later.
- Rebuilt the PSQL DB and updated all functions to write data there again.
#### 11/09/2025
- Added system and waypoint data collection routines
- Built DB tables to contain that data

## [0.2.0] - Completed: 11/04/2025
- Early big refactor 
#### 11/04/2025
- Added contract tables to DB and built associated structs
- Built Get,Update,Display routine framework for the contract data.
#### 11/03/2025
- Moved development environment to a local machine from a remote server.
- Switched from Tview to BubbleTea
- Abandoned the server performance data for now, not really relevant and distracting.
#### 11/01/2025
- Switched from PostgreSQL to SQLite

## [0.1.0] - Completed: 10/20/2025
- Project start
- Registering an agent
- Laying the basic foundation for how the TUI is going to work
- Gathering the information required from both the host server and the SpaceTraders web API
- Found and installed the base packages I'll be using for this project
- Created and set up the PGSQL DB and all permissions
- Started pre-defining some of the tables I knew I'd need.
