## [0.3.5] - Completed: 11/28/2025
#### 11/27/2025 + 11/28/2025
- Modifying color theme to be more dark gray and green themed, easier on the eyes when staring at it for hours
- Converted single line entries for Fleet and Systems into selectable cards
- Beginning to break down menus into substructures
#### 11/26/2025
- Removed unecessary struct in agent data collection routines. Updated functions accordingly.
- Tested with some new formatting on the lists and graphing and plotting data.
- Added TUI functions to view contracts in progress.
#### 11/25/2025
- Reworked the existing TUI menus, everything now focuses and updates correctly.
- Starting to build submenus for each section to actually interact with things.
- Up until today everything was purely displaying information. Next step is to make things interactive and updatable.
- Added a proper color theme to the global state based on [projekt0n/github-nvim-theme](https://github.com/projekt0n/github-nvim-theme).
#### 11/24/2025
- Moved app config structs into General package so they could be available everywhere.
    - The fact that Golang doesn't have true Global variables is really frustrating...
- Moved the menu builder funcs into their respective packages now that the structs are global (if imported)
    - This will let me keep like-code together. Once in a menu like Fleet, the code for the entire fleet menu tree will be together in the Fleet package
#### 11/23/2025
- Big refactor on the main function and how it handles passing focus around
- Created logic for displaying (with the intent of editing later) the config.yaml file
- This should set the stage for all further UI changes and make that substantially easier now that I can pass the entire app back and forth instead of just random flex primitives.
#### 11/22/2025
- This is a stub update, life happened and not much progress the last couple days.
- Initial stages of new agent registration are done.
- Needs to be automated and populate the config.yaml file with that token still but it should at least get us registered for now even if the token piece is still manual.
- Added logic to display leaderboards in the TUI. Still working out what data needs a menu in there as opposed to just running under the hood...
#### 11/18/2025
- Big changes to the UI builder & associated functions
- Can now display contracts, systems, waypoints, and ship fleet
- It's not the cleanest and needs lots of work, especially with resizing and scaling properly to the window size. 
- I also added lots of logging and 

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
