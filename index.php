<!DOCTYPE html>
<html lang="en">
    <head>
        <title>NULL SKYE</title>
        <meta   name="viewport">
        <meta   charset="UTF-8">
        <meta   content="width=device-width, initial-scale=1.0">
        <link   rel="stylesheet" href="style.css">
        <!-- <script src="https://unpkg.com/@tailwindcss/browser@4"></script> -->
        <!-- <script src="https://unpkg.com/flowbite@1.6.5/dist/flowbite.min.js"></script> -->
        <?php 
            $thisFileName = basename($_SERVER['SCRIPT_FILENAME']);
            require_once("subroutines/Main.php");
            require_once("subroutines/Sql.php");
        ?>
    </head>

    <body>
        <!-- Generate Header Bar -->
        <?php print_header($thisFileName); ?>

        <div class="min-h-screen text-primary">
            <div class="container px-4 py-8 space-y-8">
                <!-- Welcome Section -->
                <div class="grid">
                    <div class="panel">
                        <h2 class="text-xl font-bold mb-4">Welcome to NULLSKY, adventurer</h2>
                        <blockquote class="italic text-sm text-secondary mb-4" cite="https://api.spacetraders.io/v2">
                            “SpaceTraders is an open-universe space-themed game that offers a set of HTTP endpoints to control a fleet of ships. All players operate their fleet in the same universe, and the game is currently in alpha.”<br>
                            – SpaceTraders Dev Team
                        </blockquote>
                    </div>
                </div>

                <!-- Grid 1 - Server Status + Global Stats -->
                <div class="grid grid-2">
                    <div class="panel">
                        <h3 class="text-lg section-title">Server Status</h3>
                        <table class="table">
                            <?php $ServerStatus = SELECT('SELECT server_up, game_version, agents, ships, systems, waypoints, accounts, reset_freq, reset_date, next_reset, last_updated FROM server')[0]; ?>
                            <tbody>
                                <tr> <td>Status:</td>            <td class="text-right"><?= $ServerStatus['server_up'] ?>    </td></tr>
                                <tr> <td>Version:</td>           <td class="text-right"><?= $ServerStatus['game_version'] ?> </td></tr>
                                <tr> <td>Next Reset:</td>        <td class="text-right"><?= $ServerStatus['next_reset'] ?>   </td></tr>
                            </tbody>
                        </table>
                    </div>
                    <div class="panel">
                        <h3 class="text-lg section-title">Global Stats</h3>
                        <table class="table">
                            <tbody>
                                <tr> <td>Agents Registered:</td> <td class="text-right"><?= $ServerStatus['agents'] ?>       </td></tr>
                                <tr> <td>Ships Owned:</td>       <td class="text-right"><?= $ServerStatus['ships'] ?>        </td></tr>
                                <tr> <td>Systems Found:</td>     <td class="text-right"><?= $ServerStatus['systems'] ?>      </td></tr>
                                <tr> <td>Waypoints Scanned:</td> <td class="text-right"><?= $ServerStatus['waypoints'] ?>    </td></tr>
                            </tbody>
                        </table>
                    </div>
                </div>

                <!-- Grid 2 - Leaderboards -->
                <div class="grid grid-2">
                    <div class="panel">
                        <h3 class="text-lg section-title">Leaderboard (Credits)</h3>
                        <table class="table">
                            <thead>
                                <tr>
                                    <th>Agent</th>
                                    <th>Credits</th>
                                </tr>
                            </thead>
                            <tbody>
                                <?php 
                                    $Leaderboard = SELECT("SELECT * FROM leaderboard_creds");
                                    foreach ($Leaderboard as $key => $vals) {
                                        print("<tr>");
                                        print("<td>{$vals['agent']}</td>");
                                        print("<td>{$vals['credits']}</td>");
                                        print("</tr>");
                                    }
                                ?>
                            </tbody>
                        </table>
                    </div>
                    <div class="panel">
                        <h3 class="text-lg section-title">Leaderboard (Charts)</h3>
                        <table class="table">
                            <thead>
                                <tr>
                                    <th>Agent</th>
                                    <th>Charts</th>
                                </tr>
                            </thead>
                            <tbody>
                                <?php 
                                    $Leaderboard = SELECT("SELECT * FROM leaderboard_charts");
                                    foreach ($Leaderboard as $key => $vals) {
                                        print("<tr");
                                        print("<td>{$vals['agent']}</td>");
                                        print("<td>{$vals['charts']}</td>");
                                        print("</tr>");
                                    }
                                ?>
                            </tbody>
                        </table>
                    </div>
                </div>

            </div>
        </div>
    </body>
</html>
