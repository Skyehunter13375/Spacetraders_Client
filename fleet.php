<!DOCTYPE html>
<html lang="en">
    <head>
        <title>NULL SKYE</title>
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
                        ");
                        foreach ($SHIPS as $key => $vals) {
                            print("
                                <div class='panel'>
                                    <h3 class='text-lg section-title'>{$vals['symbol']}</h3>
                                    <table class='table'>
                                        <tbody>
                                            <tr> <td>Frame:</td>  <td class='text-right'>{$vals['frame']}</td></tr>
                                            <tr> <td>Role:</td>   <td class='text-right'>{$vals['role']} </td></tr>
                                            <tr> <td>Locale:</td> <td class='text-right'>{$vals['waypoint']}</td></tr>
                                            <tr> <td>Crew:</td>   <td class='text-right'>{$vals['crew_cur']} / {$vals['crew_cap']} (req: {$vals['crew_req']})</td></tr>
                                            <tr> <td>Fuel:</td>   <td class='text-right'>{$vals['fuel_cur']} / {$vals['fuel_cap']}</td></tr>
                                            <tr> <td>Morale:</td> <td class='text-right'>{$vals['morale']} / 100</td></tr>
                                        </tbody>
                                    </table>
                                </div>
                            ");
                        }


                        $remainder = count($SHIPS) % 4;
                        $loopCount = $remainder == 0 ? 0 : (4 - $remainder);
                        for ($i = 0; $i < $loopCount; $i++) {
                            print("<div class='panel'></div>");
                        }
                    ?>
                </div>
            </div>
        </div>
    </body>
</html>
