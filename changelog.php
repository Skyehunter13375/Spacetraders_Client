<!DOCTYPE html>
<html lang="en">
    <head>
        <title>NULL SKY</title>
        <meta name="viewport">
        <meta charset="UTF-8">
        <meta content="width=device-width, initial-scale=1.0">
        <link rel="stylesheet" href="Style/style.css">
        <?php
            $thisFileName = basename($_SERVER['SCRIPT_FILENAME']);
            require_once("subroutines/Main.php");
            require_once("subroutines/Sql.php");
        ?>
    </head>

    <body>
        <?php print_header($thisFileName); ?>
        <div class="container">
            <div class='panel' style='margin-bottom: 1rem;'>
            <h2 class='text-primary'>01/19/2026</h2>
            <ul class='text-white'>
                <li>Finished and tested the new agent registration code.</li>
                <li>It works as intended now, resetting everything and grabbing initial data from the registration response.</li>
                <li>Fixed a typo in my sqlitedb builder.</li>
                <li>Starting to work on the new main script so that it just runs autonomously as a systemd service.</li>
                <li>Eventually it will always be running and just run jobs placed in the SQLitedb automagically as manual tasks between scheduled events/listeners.</li>
                <li>I also realized I am mapping the system incorrectly, I currently have everything orbiting 0,0 by mistake. Moons and space stations might orbit planets so I'll fix that later tonight and add color schemes to the web UI.</li>
            </ul>
            </div>

            <div class='panel' style='margin-bottom: 1rem;'>
            <h2 class='text-primary'>01/13/2026</h2>
            <ul class='text-white'>
                <p>Took a short break for Xmas/New Year and to focus on my Linux+ certification training</p>
                <li>Stripped out the TUI entirely, I do still have that code but it's been removed from the project dirs.</li>
                <li>Converting everything over use proper state control that way PHP and Golang can intermingle and automation can be paused for manual intervention via the web UI.</li>
                <li>Go will run the entire project under the hood, PHP web pages will serve as the user interface, allowing the user to pause the game and manually interact with the pieces.</li>
                <li>Each component (fleet, agent, contracts etc) will have it's own scheduled interval to be updated on automagically unless the client is paused.</li>
                <li>Reworked the system display page to use Plotly.js instead of Chart.js. Looks much better this way in my opinion.</li>
            </ul>
            </div>

            <div class='panel' style='margin-bottom: 1rem;'>
            <h2 class='text-primary'>12/31/2025</h2>
            <ul class='text-white'>
                <li>Cleaned up a ton of the redundant CSS that was left over from switching back from Tailwind to pure CSS.</li>
                <li>Dabbled with importing a Github activity calendar but ultimately decided not to include it for now.</li>
                <li>Removed the javascript I used to read markdown for the about me and changelog and just converted it to pure HTML.</li>
                <li>Created the bones of a method to display how a system actually looks by plotting points in a JS chart and drawing fake orbit circles.</li>
                <li>Cobbled together a baseline for the poem that defines the theme of this project.</li>
                <li>Today is mostly stylistic changes before I get back into Golang and rip out all the TUI components leaving just the CLI.</li>
                <li>Set up cloudflare app path so I can finally share the link now that all pages are at least semi-functional again!!!</li>
            </ul>
            </div>

            <div class='panel' style='margin-bottom: 1rem;'>
            <h2 class='text-primary'>12/28/2025</h2>
            <ul class='text-white'>
                <li>Reworked the front end for the home page and fleet pages, still need to work on contracts and system viewer.</li>
                <li>This is the setup setp for converting all the TUI stuff over to PHP and just using GO for the back end with SQLite.</li>
                <li>I also want to get all of the PHP SQL stuff removed, so it just calls various GO programs to get the data out of SQLite and update it if needed all in one place.</li>
                <li>I am also working on abandoning TailwindCSS for raw CSS of my own design. Decided I don't like the clutter that tailwind creates and I can simplify it dramatically while keeping the same idea.</li>
                <li>Working on a plan to create an actual visual system model viewer, not sure what tech I want to use to build that just yet.</li>
                <li>Also need to write up an ABOUT ME page soon when I make the page public again.</li>
            </ul>
            </div>

            <div class='panel' style='margin-bottom: 1rem;'>
            <h2 class='text-primary'>12/26/2025</h2>
            <ul class='text-white'>
                <li>Merged my PHP and Golang projects together into one package.</li>
                <li>From now on I will serve a PHP based web UI to view and interact with the game.</li>
                <li>Golang will serve as the backbone still to update and maintain the SQLite database.</li>
                <li>Today's commits will be strictly for getting the baseline of that merger up and running.</li>
                <li>Will attempt to maintain all previous git history, not ready for a rebase just yet.</li>
            </ul>
            </div>

            <div class='panel' style='margin-bottom: 1rem;'>
            <h2 class='text-primary'>12/13/2025</h2>
            <ul class='text-white'>
                <li>Completed registration automation feature.</li>
                <li>Now the config program will automagically register a new agent after a reset (as long as the agent token is removed from the config. This will be changed later on.</li>
                <li>Initial capture of Agent, System, Fleet, and Contract data is done during registration of the new agent to minimize calls to the API.</li>
                <li>I also updated some of the logging routines and error handling to be a bit more verbose going forward.</li>
            </ul>
            </div>

            <div class='panel' style='margin-bottom: 1rem;'>
            <h2 class='text-primary'>12/12/2025</h2>
            <ul class='text-white'>
                <li>Big simplification of package structures so there's less cross importing avoiding conflicts.</li>
                <li>Working on initial registration script so that it captures preliminary data to kickstart the process after a reset.</li>
                <li>Created models for the entire registration process so it can be automated now.</li>
                <li>Set up new tables for factions</li>
            </ul>
            </div>

            <div class='panel' style='margin-bottom: 1rem;'>
            <h2 class='text-primary'>12/08/2025</h2>
            <ul class='text-white'>
                <li>Mostly finished conversion back to SQLite.</li>
                <li>Automated most of the first time setup stuff.</li>
                <li>Now you just create the config.yaml file and it should do everything automagically.</li>
                <li>Systems capture is not done yet. I need to actually sit and write my registration function to fill all of that out the first time the agent is registered. Then it becomes a non-issue.</li>
                <li>Getting close to done with a primitive version of the UI, now I can actually start playing the game and automating ship action.</li>
            </ul>
            </div>

            <div class='panel' style='margin-bottom: 1rem;'>
            <h2 class='text-primary'>12/07/2025</h2>
            <ul class='text-white'>
                <li>Converting back to SQLite...again.</li>
                <li>Condensing and simplifying the file structure, getting too large.</li>
                <li>Reworking comment structure as well as I've fully switched from VsCode to NeoVim now.</li>
            </ul>
            </div>

            <div class='panel' style='margin-bottom: 1rem;'>
            <h2 class='text-primary'>11/28/2025</h2>
            <ul class='text-white'>
                <li>Modifying color theme to be more dark gray and green themed, easier on the eyes when staring at it for hours.</li>
                <li>Converted single line entries for Fleet and Systems into selectable cards.</li>
                <li>Beginning to break down menus into substructures.</li>
            </ul>
            </div>

            <div class='panel' style='margin-bottom: 1rem;'>
            <h2 class='text-primary'>11/26/2025</h2>
            <ul class='text-white'>
                <li>Removed unecessary struct in agent data collection routines. Updated functions accordingly.</li>
                <li>Tested with some new formatting on the lists and graphing and plotting data.</li>
                <li>Added TUI functions to view contracts in progress.</li>
            </ul>
            </div>

            <div class='panel' style='margin-bottom: 1rem;'>
            <h2 class='text-primary'>11/25/2025</h2>
            <ul class='text-white'>
                <li>Reworked the existing TUI menus, everything now focuses and updates correctly.</li>
                <li>Starting to build submenus for each section to actually interact with things.</li>
                <li>Up until today everything was purely displaying information. Next step is to make things interactive and updatable.</li>
                <li>Added a proper color theme to the global state based on <a href="https://github.com/projekt0n/github-nvim-theme">projekt0n/github-nvim-theme</a>.</li>
            </ul>
            </div>

            <div class='panel' style='margin-bottom: 1rem;'>
            <h2 class='text-primary'>11/24/2025</h2>
            <ul class='text-white'>
                <li>Moved app config structs into General package so they could be available everywhere.
                <ul class='text-white'>
                    <li>The fact that Golang doesn't have true Global variables is really frustrating...</li>
                </ul>
                </li>
                <li>Moved the menu builder funcs into their respective packages now that the structs are global (if imported).
                <ul class='text-white'>
                    <li>This will let me keep like-code together. Once in a menu like Fleet, the code for the entire fleet menu tree will be together in the Fleet package.</li>
                </ul>
                </li>
            </ul>
            </div>

            <div class='panel' style='margin-bottom: 1rem;'>
            <h2 class='text-primary'>11/23/2025</h2>
            <ul class='text-white'>
                <li>Big refactor on the main function and how it handles passing focus around.</li>
                <li>Created logic for displaying (with the intent of editing later) the config.yaml file.</li>
                <li>This should set the stage for all further UI changes and make that substantially easier now that I can pass the entire app back and forth instead of just random flex primitives.</li>
            </ul>
            </div>

            <div class='panel' style='margin-bottom: 1rem;'>
            <h2 class='text-primary'>11/22/2025</h2>
            <ul class='text-white'>
                <li>This is a stub update, life happened and not much progress the last couple days.</li>
                <li>Initial stages of new agent registration are done.</li>
                <li>Needs to be automated and populate the config.yaml file with that token still but it should at least get us registered for now even if the token piece is still manual.</li>
                <li>Added logic to display leaderboards in the TUI. Still working out what data needs a menu in there as opposed to just running under the hood...</li>
            </ul>
            </div>

            <div class='panel' style='margin-bottom: 1rem;'>
            <h2 class='text-primary'>11/18/2025</h2>
            <ul class='text-white'>
                <li>Big changes to the UI builder & associated functions.</li>
                <li>Can now display contracts, systems, waypoints, and ship fleet.</li>
                <li>It's not the cleanest and needs lots of work, especially with resizing and scaling properly to the window size.</li>
                <li>I also added lots of logging.</li>
            </ul>
            </div>

            <div class='panel' style='margin-bottom: 1rem;'>
            <h2 class='text-primary'>11/16/2025</h2>
            <ul class='text-white'>
                <li>Another big refactor while I am learning which packages I like.</li>
                <li>Filling in gaps in data collection routines.</li>
            </ul>
            </div>

            <div class='panel' style='margin-bottom: 1rem;'>
            <h2 class='text-primary'>11/16/2025</h2>
            <ul class='text-white'>
                <li>Complete overhaul of postgres connections. Simplified and trimmed lots of fat.</li>
                <li>Added a config file to standardize DB connection info and account/agent tokens.</li>
                <li>Fixed all the timestamp checking (UTC vs Local time) and rebuilt the database.</li>
                <li>Abandoned Bubbletea in favor of Tview. Much simpler, easier to understand and less abstract.</li>
                <li>Shifting gears from data collection to TUI building then to start actually playing the game.</li>
                <li>Also added a quick feature milestone tracker to the README. Swiped that idea from someone else's project.</li>
            </ul>
            </div>

            <div class='panel' style='margin-bottom: 1rem;'>
            <h2 class='text-primary'>11/13/2025</h2>
            <ul class='text-white'>
                <li>New agent registration function so I don't have to keep going to the web page on reset. Requires testing still.</li>
                <li>Basic leaderboard data collection started.</li>
            </ul>
            </div>

            <div class='panel' style='margin-bottom: 1rem;'>
            <h2 class='text-primary'>11/11/2025</h2>
            <ul class='text-white'>
                <li>Created structs and function template for shipyard data collection but can't test until next reset or I move systems.</li>
            </ul>
            </div>

            <div class='panel' style='margin-bottom: 1rem;'>
            <h2 class='text-primary'>11/10/2025</h2>
            <ul class='text-white'>
                <li>Fixed my ingest routines to properly use json.Unmarshal() with wrappers to get rid of the "data" fields.</li>
                <li>Rebuilt my structs for all data types for ease of using interfaces and SQL tuning later on.</li>
                <li>Gave up on SQLite and switched back to PostgreSQL. I need the ability to store arrays in a single column for cross reference later.</li>
                <li>Rebuilt the PSQL DB and updated all functions to write data there again.</li>
            </ul>
            </div>

            <div class='panel' style='margin-bottom: 1rem;'>
            <h2 class='text-primary'>11/09/2025</h2>
            <ul class='text-white'>
                <li>Added system and waypoint data collection routines.</li>
                <li>Built DB tables to contain that data.</li>
            </ul>
            </div>

            <div class='panel' style='margin-bottom: 1rem;'>
            <h2 class='text-primary'>11/04/2025</h2>
            <ul class='text-white'>
                <li>Added contract tables to DB and built associated structs.</li>
                <li>Built Get,Update,Display routine framework for the contract data.</li>
            </ul>
            </div>

            <div class='panel' style='margin-bottom: 1rem;'>
            <h2 class='text-primary'>11/03/2025</h2>
            <ul class='text-white'>
                <li>Moved development environment to a local machine from a remote server.</li>
                <li>Switched from Tview to BubbleTea.</li>
                <li>Abandoned the server performance data for now, not really relevant and distracting.</li>
            </ul>
            </div>

            <div class='panel' style='margin-bottom: 1rem;'>
            <h2 class='text-primary'>11/01/2025</h2>
            <ul class='text-white'>
                <li>Switched from PostgreSQL to SQLite.</li>
            </ul>
            </div>

            <div class='panel' style='margin-bottom: 1rem;'>
            <h2 class='text-primary'>10/20/2025</h2>
            <ul class='text-white'>
                <li>Project start</li>
                <li>Registering an agent</li>
                <li>Laying the basic foundation for how the TUI is going to work.</li>
                <li>Gathering the information required from both the host server and the SpaceTraders web API.</li>
                <li>Found and installed the base packages I'll be using for this project.</li>
                <li>Created and set up the PGSQL DB and all permissions.</li>
                <li>Started pre-defining some of the tables I knew I'd need.</li>
            </ul>
            </div>
        </div>
    </body>
</html>
