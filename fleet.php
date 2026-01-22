<!DOCTYPE html>
<html lang="en">
    <head>
        <title>NULL SKY</title>
        <meta name="viewport">
        <meta charset="UTF-8">
        <meta content="width=device-width, initial-scale=1.0">
        <link rel="stylesheet" href="Style/style.css">
        <!-- No caching -->
        <meta http-equiv="Cache-Control" content="no-cache, no-store, must-revalidate"/>
        <meta http-equiv="Pragma" content="no-cache"/>
        <meta http-equiv="Expires" content="0"/>

        <?php
            $thisFileName = basename($_SERVER['SCRIPT_FILENAME']);
            require_once("subroutines/Main.php");
            require_once("subroutines/Sql.php");
        ?>
    </head>

    <body>
        <?php print_header($thisFileName); ?>
        <div class="min-h-screen text-primary">
            <div class="container">
                <div class="grid grid-4">
                    <?php 
                        $SHIPS = SELECT("SELECT
                                            ship.symbol,
                                            ship.role,
                                            ship.last_updated,
                                            fram.name AS frame,
                                            navg.waypoint,
                                            wayp.type AS waypoint_type,
                                            navg.status,
                                            navg.flight_mode,
                                            crew.current  AS crew_cur,
                                            crew.required AS crew_req,
                                            crew.capacity AS crew_cap,
                                            crew.morale,
                                            fuel.current  AS fuel_cur,
                                            fuel.capacity AS fuel_cap
                                        FROM ships AS ship
                                        JOIN ship_nav   AS navg ON navg.ship = ship.symbol
                                        JOIN ship_crew  AS crew ON crew.ship = ship.symbol
                                        JOIN ship_fuel  AS fuel ON fuel.ship = ship.symbol
                                        JOIN ship_frame AS fram ON fram.ship = ship.symbol
                                        JOIN waypoints  AS wayp ON navg.waypoint = wayp.symbol
                        ");
                        foreach ($SHIPS as $key => $vals) {
                            print("
                                <div class='panel'>
                                    <h3 class='text-lg section-title'>{$vals['symbol']}</h3>
                                    <table class='table'>
                                        <tbody>
                                            <tr> <td>Frame:</td>    <td class='text-right'>{$vals['frame']}</td></tr>
                                            <tr> <td>Role:</td>     <td class='text-right'>{$vals['role']} </td></tr>
                                            <tr> <td>Crew:</td>     <td class='text-right'>{$vals['crew_cur']} / {$vals['crew_cap']} (req: {$vals['crew_req']})</td></tr>
                                            <tr> <td>Fuel:</td>     <td class='text-right'>{$vals['fuel_cur']} / {$vals['fuel_cap']}</td></tr>
                                            <tr> <td>Waypoint:</td> <td class='text-right'><button class='open-modal' data-symbol='{$vals['waypoint']}'  data-type='System'>{$vals['waypoint']} ({$vals['waypoint_type']})</button></td></tr>
                                            <tr> <td>Contract:</td> <td class='text-right'><button class='open-modal' data-symbol='PLACEHOLDER_CONTRACT' data-type='Contract'>View Contract</td></tr>
                                            <tr> <td>Cargo:</td>    <td class='text-right'><button class='open-modal' data-symbol='{$vals['symbol']}'    data-type='Cargo'>View Cargo</button></td></tr>
                                        </tbody>
                                    </table>
                                </div>
                            ");
                        }


//                         $remainder = count($SHIPS) % 4;
//                         $loopCount = $remainder == 0 ? 0 : (4 - $remainder);
//                         for ($i = 0; $i < $loopCount; $i++) {
//                             print("<div class='panel'></div>");
//                         }
                    ?>
                </div>
            </div>
        </div>
        <div id="modal-backdrop" class="modal-backdrop hidden">
            <div class="modal-window">
                <div id="modal-content"></div>
            </div>
        </div>


        <script>
            document.addEventListener('DOMContentLoaded', function () {

                const MODAL_ENDPOINT = '/subroutines/Modal.php';

                const backdrop = document.getElementById('modal-backdrop');
                const content  = document.getElementById('modal-content');

                // Open modal (event delegation)
                document.addEventListener('click', function (e) {
                    const btn = e.target.closest('.open-modal');
                    if (!btn) return;

                    const params = new URLSearchParams();

                    // Collect all data-* attributes from the button
                    for (const [key, value] of Object.entries(btn.dataset)) {
                        params.append(key, value);
                    }

                    content.innerHTML = 'Loading…';
                    backdrop.classList.remove('hidden');

                    fetch(MODAL_ENDPOINT + '?' + params.toString())
                        .then(r => {
                            if (!r.ok) throw new Error('Request failed');
                            return r.text();
                        })
                        .then(html => {
                            content.innerHTML = html;
                        })
                        .catch(err => {
                            console.error(err);
                            content.innerHTML = 'Error loading data';
                        });
                });

                // Close modal when clicking outside modal window
                backdrop.addEventListener('click', function (e) {
                    if (e.target === backdrop) {
                        backdrop.classList.add('hidden');
                        content.innerHTML = '';
                    }
                });

            });
        </script>
    </body>
</html>

