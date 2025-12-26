<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>NULL SKYE</title>
        <script src="https://unpkg.com/@tailwindcss/browser@4"></script>
        <link href="https://cdn.jsdelivr.net/npm/flowbite@3.1.2/dist/flowbite.min.css" rel="stylesheet" />
        <?php $thisFileName = basename($_SERVER['SCRIPT_FILENAME']); ?>
        <?php 
            require_once("subroutines/Main.php");
        ?>
    </head>

    <body>
        <?php 
            print_header($thisFileName); 
//             $CONTRACTS = pg_raw(FALSE, "SELECT
//                                             *,
//                                             array_to_string(deliver_material, ',') AS materials,
//                                             array_to_string(deliver_destination, ',') AS destinations,
//                                             array_to_string(deliver_required, ',') AS required_quantity,
//                                             array_to_string(deliver_fulfilled, ',') AS fulfilled_quantity
//                                         FROM contracts;"
//             );
        ?>
        <div class="min-h-screen bg-neutral-800 bg-opacity-70 text-green-400 font-mono">
            <div class="max-w-7/8 mx-auto px-4 py-8 space-y-8">
                <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
                    <?php 
                        foreach ($CONTRACTS as $cKey => $cVals) {
                            $materials = explode(",", str_replace("'", '', $cVals['materials']));
                            $destinate = explode(",", str_replace("'", '', $cVals['destinations']));
                            $required  = explode(",", str_replace("'", '', $cVals['required_quantity']));
                            $fulfilled = explode(",", str_replace("'", '', $cVals['fulfilled_quantity']));

                            print("
                                <div class='bg-neutral-800 bg-opacity-70 border border-green-800 p-4 rounded shadow-lg overflow-x-hidden'>
                                    <table class='w-full text-sm'>
                                        <tbody>
                                            <tr><td class='py-1'>Faction:</td><td class='text-right text-white'>{$cVals['faction']}</td></tr>
                                            <tr><td class='py-1'>Type:</td><td class='text-right text-white'>{$cVals['type']}</td></tr>
                                            <tr><td class='py-1'>Payment:</td><td class='text-right text-white'>". ($cVals['payment_upfront'] + $cVals['payment_completion']) ."</td></tr>
                            ");
                            foreach ($materials as $mKey => $garbage) {
                                print("<tr><td class='py-1'>Material:</td><td class='text-right text-white'>{$materials[$mKey]} ({$fulfilled[$mKey]}/{$required[$mKey]})</td></tr>");
                            }
                            print("
                                            <tr><td class='py-1'>Accepted:</td><td class='text-right text-white'>". ($cVals['accepted'] == 1 ? 'Yes' : 'No') ."</td></tr>
                                            <tr><td class='py-1'>Fulfilled:</td><td class='text-right text-white'>".($cVals['fulfilled'] == 1 ? 'Yes' : 'No') ."</td></tr>
                                            <tr><td class='py-1'>Expires:</td><td class='text-right text-white'>{$cVals['expires']}</td></tr>
                                            <tr><td class='py-1'>Deadline:</td><td class='text-right text-white'>{$cVals['deadline']}</td></tr>
                                        </tbody>
                                    </table>
                                </div>
                            ");
                        }


                        $remainder = count($CONTRACTS) % 4;
                        $loopCount = $remainder == 0 ? 0 : (4 - $remainder);
                        for ($i = 0; $i < $loopCount; $i++) {
                            print("
                                <div class='bg-neutral-800 bg-opacity-70 border border-green-800 p-4 rounded shadow-lg overflow-x-hidden'>
                                    <div class='flex items-center justify-between border-b border-green-800 pb-1 mb-2'>
                                        <h3 class='text-lg'>To Be Announced...</h3>
                                    </div>
                                </div>
                            ");
                        }
                    ?>
                </div>
            </div>
        </div>

        <!------------------------------------------------------------------ BEGIN MODAL STUFF HERE ------------------------------------------------------------------>
        <script src="https://unpkg.com/flowbite@1.6.5/dist/flowbite.min.js"></script>
    </body>
</html>
