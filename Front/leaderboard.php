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
        <?php print_header($thisFileName); ?>
        <div class="min-h-screen bg-neutral-800 bg-opacity-70 text-green-400 font-mono">
            <div class="flex justify-center mx-auto px-4 py-8 space-y-8">
                <div class="overflow-x-auto w-fit bg-neutral-800 rounded-lg shadow-md p-4 mb-4 border border-green-800">
                        <table class="mx-auto text-center text-white max-w-screen-xl border-collapse border border-gray-700">
                            <thead>
                                <tr class="bg-neutral-800 bg-opacity-70">
                                    <th class='text-center border border-gray-600 px-4 py-2 whitespace-nowrap'>Agent</th>
                                    <th class='text-center border border-gray-600 px-4 py-2 whitespace-nowrap'>Credits</th>
                                </tr>
                            </thead>
                            <tbody>
                                <?php 
                                    $Leaderboard = SELECT("SELECT * FROM leaderboard_creds");
                                    foreach ($Leaderboard as $key => $vals) {
                                        print("<tr class='odd:bg-neutral-900 even:bg-neutral-950'>");
                                        print("<td class='border border-gray-600 px-4 py-2'>{$vals['agent']}</td>");
                                        print("<td class='border border-gray-600 px-4 py-2'>{$vals['credits']}</td>");
                                        print("</tr>");
                                    }
                                ?>
                            </tbody>
                        </table>
                </div>
                <div class="overflow-x-auto w-fit bg-neutral-800 rounded-lg shadow-md p-4 mb-4 border border-green-800">
                        <table class="mx-auto text-center text-white max-w-screen-xl border-collapse border border-gray-700">
                            <thead>
                                <tr class="bg-neutral-800 bg-opacity-70">
                                    <th class='text-center border border-gray-600 px-4 py-2 whitespace-nowrap'>Agent</th>
                                    <th class='text-center border border-gray-600 px-4 py-2 whitespace-nowrap'>Charts</th>
                                </tr>
                            </thead>
                            <tbody>
                                <?php 
                                    $Leaderboard = SELECT("SELECT * FROM leaderboard_charts");
                                    foreach ($Leaderboard as $key => $vals) {
                                        print("<tr class='odd:bg-neutral-900 even:bg-neutral-950'>");
                                        print("<td class='border border-gray-600 px-4 py-2'>{$vals['agent']}</td>");
                                        print("<td class='border border-gray-600 px-4 py-2'>{$vals['charts']}</td>");
                                        print("</tr>");
                                    }
                                ?>
                            </tbody>
                        </table>
                </div>
            </div>
        </div>
        <script src="https://unpkg.com/flowbite@1.6.5/dist/flowbite.min.js"></script>
    </body>
</html>
