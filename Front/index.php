<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>NULL SKYE</title>
        <script src="https://unpkg.com/@tailwindcss/browser@4"></script>
        <?php $thisFileName = basename($_SERVER['SCRIPT_FILENAME']); ?>
        <?php 
            require_once("subroutines/Main.php");
            require_once("subroutines/Sql.php");
        ?>
    </head>

    <body>
        <!-- Generate Header Bar -->
        <?php print_header($thisFileName); ?>

        <div class="min-h-screen bg-neutral-800 bg-opacity-70 text-green-400 font-mono">
            <div class="max-w-7/8 mx-auto px-4 py-8 space-y-8">
                <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                    <!-- Server Detail -->
                    <?php $ServerStatus = SELECT('SELECT server_up, game_version, agents, ships, systems, waypoints, accounts, reset_freq, reset_date, next_reset, last_updated FROM server')[0]; ?>
                    <div class="bg-neutral-800 bg-opacity-70 border border-green-800 p-4 rounded shadow-lg overflow-x-hidden">
                        <h3 class="text-lg mb-2 border-b border-green-800 pb-1">Server Status</h3>
                        <table class="w-full text-sm">
                            <tbody>
                                <tr><td class="py-1">Status:</td><td class="text-right text-white"><?= $ServerStatus['server_up'] ?></td></tr>
                                <tr><td class="py-1">Version:</td><td class="text-right text-white"><?= $ServerStatus['game_version'] ?></td></tr>
                                <tr><td class="py-1">Next Reset:</td><td class="text-right text-white"><?= $ServerStatus['next_reset'] ?></td></tr>
                                <tr><td class="py-1">Agents Registered:</td><td class="text-right text-white"><?= $ServerStatus['agents'] ?></td></tr>
                                <tr><td class="py-1">Ships Owned:</td><td class="text-right text-white"><?= $ServerStatus['ships'] ?></td></tr>
                                <tr><td class="py-1">Systems Found:</td><td class="text-right text-white"><?= $ServerStatus['systems'] ?></td></tr>
                                <tr><td class="py-1">Waypoints Scanned:</td><td class="text-right text-white"><?= $ServerStatus['waypoints'] ?></td></tr>
                            </tbody>
                        </table>
                    </div>

                <!-- Welcome Section -->
                <div class="bg-neutral-800 border border-green-800 rounded-lg p-6 shadow-md">
                    <h2 class="text-2xl font-bold mb-4">Welcome to NULLSKY, adventurer</h2>
                    <blockquote class="italic text-sm text-gray-400 border-l-4 border-green-800 pl-4 mb-4" cite="https://api.spacetraders.io/v2">
                        "SpaceTraders is a headless API and fleet-management game where players can work together or against each other to trade, explore, expand, and conquer in a dynamic and growing universe. Build your own UI, write automated scripts, or just play the game from the comfort of your terminal."<br>- SpaceTraders API Dev Team
                    </blockquote>
                </div>
            </div>
        </div>
        <script src="https://unpkg.com/flowbite@1.6.5/dist/flowbite.min.js"></script>
    </body>
</html>
