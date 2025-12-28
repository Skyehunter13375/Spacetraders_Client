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
        <?php print_header($thisFileName); ?>
        <div class="min-h-screen bg-neutral-800 bg-opacity-70 text-green-400 font-mono">
            <div class="max-w-7/8 mx-auto px-4 py-8 space-y-8">
                <div class="overflow-x-auto w-7/8 bg-neutral-800 rounded-lg shadow-md p-4 mb-4 border border-green-800">
                    <table class='w-full border-collapse border border-gray-700 text-white'>
                        <thead>
                            <tr class='bg-neutral-800 bg-opacity-70'>
                                <th class='border border-gray-600 px-4 py-2 text-left whitespace-nowrap'>Waypoint</th>
                                <th class='border border-gray-600 px-4 py-2 text-left whitespace-nowrap'>Type</th>
                                <th class='border border-gray-600 px-4 py-2 text-left whitespace-nowrap'>X Coordinate</th>
                                <th class='border border-gray-600 px-4 py-2 text-left whitespace-nowrap'>Y Coordinate</th>
                                <th class='border border-gray-600 px-4 py-2 text-left whitespace-nowrap'>Orbitals</th>
                                <th class='border border-gray-600 px-4 py-2 text-left whitespace-nowrap'>Traits</th>
                                <th class='border border-gray-600 px-4 py-2 text-left whitespace-nowrap'>Faction</th>
                                <th class='border border-gray-600 px-4 py-2 text-left whitespace-nowrap'>Under Construction</th>
                            </tr>
                        </thead>
                        <tbody>
                            <?php
                                // $Systems = SELECT("SELECT symbol,type,x_coord,y_coord,ARRAY_TO_STRING(orbitals, '<br>') AS orbitals,ARRAY_TO_STRING(traits, '<br>') AS traits,faction,construction FROM systems;");
                                foreach ($Systems as $sKey => $sVals) {
                                    print("<tr class='odd:bg-neutral-900 even:bg-neutral-950'>");
                                    print("<td class='border border-gray-600 px-4 py-2'>{$sVals['symbol']}</td>");
                                    print("<td class='border border-gray-600 px-4 py-2'>".ucwords(strtolower(str_replace('_', ' ', $sVals['type'])))."</td>");
                                    print("<td class='border border-gray-600 px-4 py-2'>{$sVals['x_coord']}</td>");
                                    print("<td class='border border-gray-600 px-4 py-2'>{$sVals['y_coord']}</td>");
                                    print("<td class='border border-gray-600 px-4 py-2'>{$sVals['orbitals']}</td>");
                                    print("<td class='border border-gray-600 px-4 py-2'>{$sVals['traits']}</td>");
                                    print("<td class='border border-gray-600 px-4 py-2'>{$sVals['faction']}</td>");
                                    print("<td class='border border-gray-600 px-4 py-2'>{$sVals['construction']}</td>");
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
