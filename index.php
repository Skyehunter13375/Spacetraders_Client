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
        <!-- Generate Header Bar -->
        <?php print_header($thisFileName); ?>

        <div class="min-h-screen text-primary">
            <div class="container">
                <!-- Welcome Section -->
                <div class="grid">
                    <div class="panel">
                        <h2 class="text-xl section-title">Welcome to Null Sky</h2>
                        <p class='text-white'>
                            Welcome to my passion project / fever dream / insomnia fueled effort to build something cool for myself! First off, if you came here looking for tips on playing this wonderful game you have sadly come to the wrong place adventurer!<br>
                            I work on this purely in my spare time as a fun way to kill time and also build some skills while doing it. To be brutally honest it's a good escape from video games until Resident Evil 9 releases in Feb 2026!<br><br>

                            I have no words of wisdom here yet so feel free to check these pages out if you'd like. I have been working on this project for some time, starting over from scratch several times, but only the bones of the project are complete at this time.<br>
                            More will be added over time! I refuse to give up on this project and I WILL eventually get to play this incredible game someday!
                        </p><br>
                        <blockquote class="italic text-sm text-secondary" cite="">
                            “Between quiet stars, the abyss invites thought. I look inward and find room to grow, echoes turning into possibility. The universe holds its silence gently, and I meet it with calm awareness, learning beneath the Null Sky.”<br>
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
                                <tr> <td>Last Reset:</td>        <td class="text-right">PLACEHOLDER TEXT</td></tr>
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
                                        print("<tr>");
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
