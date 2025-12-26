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

            function renderModulesTable($ShipName, $ModuleData) {
                if (isset($ModuleData[$ShipName])) {
                    print("
                        <table class='w-full border-collapse border border-gray-700 text-white'>
                            <thead>
                                <tr class='bg-neutral-800 bg-opacity-70'>
                                    <th class='border border-gray-600 px-4 py-2 text-left whitespace-nowrap'>Name</th>
                                    <th class='border border-gray-600 px-4 py-2 text-left whitespace-nowrap'>Capacity</th>
                                    <th class='border border-gray-600 px-4 py-2 text-left whitespace-nowrap'>Power</th>
                                    <th class='border border-gray-600 px-4 py-2 text-left whitespace-nowrap'>Crew</th>
                                    <th class='border border-gray-600 px-4 py-2 text-left whitespace-nowrap'>Slots</th>
                                    <th class='border border-gray-600 px-4 py-2 text-left whitespace-nowrap'>Description</th>
                                </tr>
                            </thead>
                            <tbody>
                    ");
                    foreach ($ModuleData[$ShipName] as $mKey => $mVals) {
                        print("
                                <tr class='odd:bg-neutral-900 even:bg-neutral-950'>
                                    <td class='border border-gray-600 px-4 py-2'>{$mVals['name']}</td>
                                    <td class='border border-gray-600 px-4 py-2'>{$mVals['capacity']}</td>
                                    <td class='border border-gray-600 px-4 py-2'>{$mVals['power_required']}</td>
                                    <td class='border border-gray-600 px-4 py-2'>{$mVals['crew_required']}</td>
                                    <td class='border border-gray-600 px-4 py-2'>{$mVals['slots']}</td>
                                    <td class='border border-gray-600 px-4 py-2'>{$mVals['description']}</td>
                                </tr>"
                        );
                    }
                    print("
                            </tbody>
                        </table>"
                    );
                } else {
                    print("<li>No modules installed");
                }
            }

            function renderMountsTable($ShipName, $MountData) {
                if (isset($MountData[$ShipName])) {
                    print("
                        <table class='w-full border-collapse border border-gray-700 text-white'>
                            <thead>
                                <tr class='bg-neutral-800 bg-opacity-70'>
                                    <th class='border border-gray-600 px-4 py-2 text-left whitespace-nowrap'>Name</th>
                                    <th class='border border-gray-600 px-4 py-2 text-left whitespace-nowrap'>Strength</th>
                                    <th class='border border-gray-600 px-4 py-2 text-left whitespace-nowrap'>Power</th>
                                    <th class='border border-gray-600 px-4 py-2 text-left whitespace-nowrap'>Crew</th>
                                    <th class='border border-gray-600 px-4 py-2 text-left whitespace-nowrap'>Deposits</th>
                                    <th class='border border-gray-600 px-4 py-2 text-left whitespace-nowrap'>Description</th>
                                </tr>
                            </thead>
                            <tbody>"
                    );
                    foreach ($MountData[$ShipName] as $mKey => $mVals) {
                        print("<tr class='odd:bg-neutral-900 even:bg-neutral-950'>
                                    <td class='border border-gray-600 px-4 py-2'>{$mVals['name']}</td>
                                    <td class='border border-gray-600 px-4 py-2'>{$mVals['strength']}</td>
                                    <td class='border border-gray-600 px-4 py-2'>{$mVals['power_required']}</td>
                                    <td class='border border-gray-600 px-4 py-2'>{$mVals['crew_required']}</td>
                                    <td class='border border-gray-600 px-4 py-2'>".str_replace(",", "<br>", str_replace(['{','}'], '', $mVals['deposits']))."</td>
                                    <td class='border border-gray-600 px-4 py-2'>{$mVals['description']}</td>
                                </tr>"
                        );
                    }
                    print("
                            </tbody>
                        </table>"
                    );
                } else {
                    print("<li>No mounted devices");
                }
            }

        ?>
    </head>

    <style>
        .scanline-container::after {
            content: '';
            position: absolute;
            top: 0;
            left: -10%;
            height: 100%;
            width: 2px;
            background: rgba(255, 255, 255, 0.6);
            box-shadow: 0 0 6px rgba(255, 255, 255, 0.8);
            animation: scanline 3s linear infinite;
            z-index: 2;
        }

        .scanline::after { left: -10%; width: 2px; animation: scanline 3s linear infinite; }

        @keyframes scanline {
            0%   { left: -10%; }
            100% { left: 130%; }
        }
    </style>


    <body>
        <?php 
            print_header($thisFileName); 
//             $SHIPS = pg_raw(FALSE, "SELECT 
//                                         ship.ship,ship.name,ship.role,ship.fuel_max,ship.fuel_current,ship.fuel_consumed_last_voyage,ship.cargo_max,ship.cargo_current,
//                                         navg.status,navg.waypoint,navg.flight_mode,navg.origin,navg.origin_type,navg.destination,navg.destination_type,navg.arrival,navg.departure,
//                                         crew.current AS crew_current,crew.capacity AS crew_capacity,crew.required,crew.rotation,crew.morale,crew.wages,
//                                         fram.name AS frame
//                                     FROM ships           AS ship
//                                     JOIN ship_navigation AS navg ON ship.ship = navg.ship
//                                     JOIN ship_crew       AS crew ON ship.ship = crew.ship
//                                     JOIN ship_frame      AS fram ON ship.ship = fram.ship
//             ");
//
//             $MODULES = pg_raw(FALSE, "SELECT * from ship_modules;");
//             foreach ($MODULES as $mKey => $mVals) {
//                 $FIX_MODS[$mVals['ship']][] = $mVals;
//             }
// 
//             $MOUNTS  = pg_raw(FALSE, "SELECT * FROM ship_mounts;");
//             foreach ($MOUNTS as $mKey => $mVals) {
//                 $FIX_MOUNT[$mVals['ship']][] = $mVals;
//             }

            // $INVENTORY = pg_raw(FALSE, "SELECT * FROM ship_inventory;");
            // DEBUG_printr($FIX_MODS);
        ?>
        <div class="min-h-screen bg-neutral-800 bg-opacity-70 text-green-400 font-mono">
            <div class="max-w-7/8 mx-auto px-4 py-8 space-y-8">
                <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
                    <?php 
                        foreach ($SHIPS as $ship) {
                            print("
                                <div class='bg-neutral-800 bg-opacity-70 border border-green-800 p-4 rounded shadow-lg overflow-x-hidden'>
                                    <div class='flex items-center justify-between border-b border-green-800 pb-1 mb-2'>
                                        <h3 class='text-lg'>{$ship['ship']}</h3>
                                        <button data-modal-target='{$ship['ship']}' data-modal-toggle='{$ship['ship']}' type='button' class='px-3 mb-1 h-4 text-sm bg-green-600 text-white hover:bg-green-800'>Details</button>
                                    </div>
                                    <table class='w-full text-sm'>
                                        <tbody>
                                            <tr><td class='py-1'>Role:</td><td class='text-right text-white'>". ucfirst(strtolower($ship['role'])) ."</td></tr>
                                            <tr><td class='py-1'>Type:</td><td class='text-right text-white'>{$ship['frame']}</td></tr>
                                            <tr><td class='py-1'>Status:</td><td class='text-right text-white'>". ucfirst(strtolower($ship['status'])) ."</td></tr>
                                            <tr><td class='py-1'>Location:</td><td class='text-right text-white'>{$ship['waypoint']} ({$ship['origin_type']})</td></tr>
                            ");
                            if ($ship['waypoint'] == $ship['destination']) {
                                print("     <tr><td class='py-1'>Destination:</td><td class='text-right text-white'>----</td></tr>
                                            <tr><td class='py-1'>ETA:</td><td class='text-right text-white'>----</td></tr>
                                ");
                            } else {
                                print("     <tr><td class='py-1'>Destination:</td><td class='text-right text-white'>{$ship['destination']} ({$ship['destination_type']})</td></tr>
                                            <tr><td class='py-1'>ETA:</td><td class='text-right text-white'>{$ship['arrival']}</td></tr>
                                ");
                            }


                            $fuel_percent    = ($ship['fuel_max'] != 0 ) ? (( $ship['fuel_current'] / $ship['fuel_max']) * 100) : 100;
                            $crew_percent    = ($ship['crew_capacity'] != 0 ) ? (( $ship['crew_current'] / $ship['crew_capacity']) * 100) : 100;
                            print("
                                            <tr>
                                                <td class='py-1'>Fuel:</td>
                                                <td class='text-right text-white'>
                                                    <div class='w-full bg-gray-800 bg-opacity-70 rounded h-5 overflow-hidden shadow-inner relative scanline-container'>
                                                    <div class='bg-green-700 h-full transition-all duration-500 pl-2 text-left text-sm text-white flex items-center style='width: {$fuel_percent}%;'>
                                                        {$ship['fuel_current']} / {$ship['fuel_max']}
                                                    </div>
                                                    </div>
                                                </td>
                                            </tr>

                                            <tr>
                                                <td class='py-1'>Crew:</td>
                                                <td class='text-right text-white'>
                                                    <div class='w-full bg-gray-800 bg-opacity-70 rounded h-5 overflow-hidden shadow-inner relative scanline-container'>
                                                    <div class='bg-violet-700 h-full transition-all duration-500 pl-2 text-left text-sm text-white flex items-center' style='width: {$crew_percent}%;'>
                                                        {$ship['crew_current']} / {$ship['crew_capacity']}
                                                    </div>
                                                    </div>
                                                </td>
                                            </tr>

                                            <tr>
                                                <td class='py-1'>Morale:</td>
                                                <td class='text-right text-white'>
                                                    <div class='w-full bg-gray-800 rounded h-5 overflow-hidden shadow-inner relative scanline-container'>
                                                    <div class='bg-amber-700 bg-opacity-80 h-full transition-all duration-500 pl-2 text-left text-sm text-white flex items-center style='width: {$ship['morale']}%;'>
                                                        {$ship['morale']}
                                                    </div>
                                                    </div>
                                                </td>
                                            </tr>
                                        </tbody>
                                    </table>
                                </div>
                            ");
                        }


                        $remainder = count($SHIPS) % 4;
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
        <?php
            foreach ($SHIPS as $sKey => $sVal) {
                print("
                    <div id='{$sVal['ship']}' tabindex='-1' aria-hidden='true' class='text-green-400 hidden overflow-y-auto overflow-x-hidden fixed top-0 right-0 left-0 z-50 justify-center items-center w-full md:inset-0 h-[calc(100%-1rem)] max-h-full'>
                        <div class='relative p-4 w-7/8 max-h-full'>
                            <div class='relative bg-white rounded-lg shadow-sm'>
                                <div class='bg-neutral-800 bg-opacity-70 border border-green-800 p-4 rounded shadow-lg overflow-x-hidden'>
                                    <div class='flex text-center justify-between p-4 md:p-5 border-b rounded-t border-gray-200'>
                                        <h3 class='text-xl'>{$sVal['ship']} Detail</h3>
                                        <button type='button' class='text-red-600 bg-transparent hover:bg-gray-200 hover:text-red-900 rounded-lg text-sm w-8 h-8 ms-auto inline-flex justify-center items-center' data-modal-hide='{$sVal['ship']}'>
                                            <svg class='w-3 h-3' aria-hidden='true' xmlns='http://www.w3.org/2000/svg' fill='none' viewBox='0 0 14 14'>
                                                <path stroke='currentColor' stroke-linecap='round' stroke-linejoin='round' stroke-width='2' d='m1 1 6 6m0 0 6 6M7 7l6-6M7 7l-6 6'/>
                                            </svg>
                                            <span class='sr-only'>Close modal</span>
                                        </button>
                                    </div>

                                    <div class='overflow-x-auto w-full p-4 md:p-5 space-y-4'>
                                        <h3 class='text-center text-xl'>Modules</h3>");
                                        renderModulesTable($sVal['ship'], $FIX_MODS);
                print("             </div>
                                    <hr class='my-4 border-gray-600'>
                                    <div class='overflow-x-auto w-full p-4 md:p-5 space-y-4'>
                                        <h3 class='text-center text-xl'>Mounts</h3>");
                                        renderMountsTable($sVal['ship'], $FIX_MOUNT);
                print("             </div>
                                    <hr class='my-4 border-gray-600'>
                                    <div class='overflow-x-auto w-full p-4 md:p-5 space-y-4'>
                                        <h3 class='text-center text-xl'>Current Inventory</h3>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                ");
            }
        ?>

        <script src="https://unpkg.com/flowbite@1.6.5/dist/flowbite.min.js"></script>
    </body>
</html>
