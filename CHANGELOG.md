### [0.0.06] - 10/25/2025
###### Added
- Built up the TUI for the ship display, now can display up to 4 ships at once and their basic info
- Including percentage bars for crew capacity, fuel capacity, cargo capacity etc. 
###### Changed
- Switched from PSQL to SQLite for the sake of learning something new
    - Also for the added benefit of being easily portable in the future
- Rebuilt all of the existing tables from PSQL into SQLite and began populating them
- Modified all of the existing routines to utilize SQLite as well
###### Fixed
- N/A

--- 

### [0.0.05] - 10/25/2025
###### Added
- Built new SQL tables for ships, components, modules etc.
- Created a single function that collects ship data from the web API and ingests that into PSQL
###### Changed
- Further breakdown of the tview.Primitives so I have more granular control over what is displayed
- Massively simplified the directory structure of the project.
###### Fixed
- N/A

--- 

### [0.0.04] - 10/25/2025
###### Added
- Registered a new agent and repopulated all of the data automagically
- Functions complete to display the basics of the game and agent state in the TUI now
###### Changed
- N/A
###### Fixed
- N/A

--- 

### [0.0.03] - 10/24/2025
###### Added
- Data collection from the host server so I can monitor system performance as well as game state all in once place.
- Enabled sysstat logging so I can keep historical record of this as well.
###### Changed
- Updated all required functions to return tview.Primitives instead of just raw text
###### Fixed
- N/A

--- 

### [0.0.02] - 10/21/2025
###### Added
- Integration with PGSQL
- Began ingesting my agent and game server data
###### Changed
- Project was monolithic to start, broke it down into separate directories for easy maintainability.
###### Fixed
- N/A

--- 

### [0.0.01] - 10/20/2025
###### Added
- First day of a new project!
- Built a basic main.go for testing and learning the language
- Found and installed the base packages I'll be using for this project
    - tview
    - pgx
    - gopsutil
- Created and set up the PGSQL DB and all permissions
- Started pre-defining some of the tables I knew I'd need.
###### Changed
- N/A
###### Fixed
- N/A

--- 
