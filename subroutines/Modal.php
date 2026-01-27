<?php
    $thisFileName = basename($_SERVER['SCRIPT_FILENAME']);
    require_once("Main.php");
    require_once("Sql.php");
    $symbol = $_GET['symbol'];
    $desire = $_GET['type'];

    switch ($desire) {
        case 'System':
            $SYST = SELECT("SELECT system FROM waypoints WHERE symbol = '{$symbol}'")[0]['system'];
            $WAYP = SELECT("SELECT system,symbol,type,x,y,orbits,construction FROM waypoints WHERE system = '{$SYST}' ORDER BY type");
            print("
                    <div class='panel'>
                        <table class='table'>
                            <thead>
                                <tr>
                                    <th>System</th>
                                    <th>Symbol</th>
                                    <th>Type</th>
                                    <th>X</th>
                                    <th>Y</th>
                                    <th>Orbits</th>
                                    <th>Construction</th>
                                </tr>
                            </thead>
                            <tbody>
            ");
            foreach ($WAYP as $key => $val) {
                if ($val['symbol'] == $symbol) {
                    print("     <tr style='background-color: var(--border-accent);'>");
                } else {
                    print("     <tr>");
                }
                print("
                                    <td>{$val['system']}</td>
                                    <td>{$val['symbol']}</td>
                                    <td>{$val['type']}</td>
                                    <td>{$val['x']}</td>
                                    <td>{$val['y']}</td>
                                    <td>{$val['orbits']}</td>
                                    <td>{$val['construction']}</td>
                                </tr>
                ");
            }
            print("
                            </tbody>
                        </table>
                    </div>
            ");
            break;


        case 'Contract':
            $AVAIL = SELECT("SELECT id,faction,type,pay_on_accept,pay_on_complete,accepted,fulfilled,deadline,expiration,deadline_to_accept FROM contracts WHERE fulfilled = 0");
            $COMPL = SELECT("SELECT id,faction,type,pay_on_accept,pay_on_complete,accepted,fulfilled,deadline,expiration,deadline_to_accept FROM contracts WHERE fulfilled_by = '{$symbol}'");
            print("
                    <h3 class='text-primary text-lg section-title'>Available Contracts</h3>
                    <div class='grid grid-4'>
            ");
            foreach ($AVAIL as $key => $vals) {
                print("
                    <div class='panel'>
                        <h3 class='text-primary text-lg section-title'>{$vals['id']}</h3>
                        <table class='table'>
                            <tbody>
                                <tr> <td>Type:</td>         <td class='text-right'>{$vals['type']}</td></tr>
                                <tr> <td>Upfront Pay:</td>  <td class='text-right'>{$vals['pay_on_accept']}</td></tr>
                                <tr> <td>Complete Pay:</td> <td class='text-right'>{$vals['pay_on_complete']}</td></tr>
                                <tr> <td>Deadline:</td>     <td class='text-right'>{$vals['deadline']}</td></tr>
                                <tr> <td>Expiration:</td>   <td class='text-right'>{$vals['deadline_to_accept']}</td></tr>
                            </tbody>
                        </table>
                    </div>
                ");
            }
            print("
                </div><br>
                <h3 class='text-primary text-lg section-title'>Completed Contracts</h3>
                    <table class='table'>
                        <thead>
                            <tr>
                                <th>ID</th>
                                <th>Completed By</th>
                                <th>Payment</th>
                                <th>Faction</th>
                                <th>Date</th>
                            </tr>
                        </thead>
                        <tbody>
                        </tbody>
                    </table>
            </div>
            ");
            break;


        case 'Components':
            $ENGINE = SELECT("SELECT * FROM ship_engine WHERE ship = '{$symbol}'");
            $FRAME  = SELECT("SELECT * FROM ship_frame WHERE ship = '{$symbol}'");
            $REACTR = SELECT("SELECT * FROM ship_reactor WHERE ship = '{$symbol}'");
            $CREW   = SELECT("SELECT * FROM ship_crew WHERE ship = '{$symbol}'");
            $MODULE = SELECT("SELECT * FROM ship_modules WHERE ship = '{$symbol}'");
            $MOUNTS = SELECT("SELECT * FROM ship_mounts WHERE ship = '{$symbol}'");
            print("
                <div class='grid grid-4'>
                    <div class='panel'>
                        <h3 class=text-primary text-lg section-title'>Frame</h3>
                        <table class='table'>
                            <tbody>
                                <tr> <td>Frame</td>          <td class='text-right'>{$FRAME[0]['name']}</td></tr>
                                <tr> <td>Module Slots</td>   <td class='text-right'>{$FRAME[0]['module_slots']}</td></tr>
                                <tr> <td>Mount Points</td>   <td class='text-right'>{$FRAME[0]['mount_points']}</td></tr>
                                <tr> <td>Fuel Capacity</td>  <td class='text-right'>{$FRAME[0]['fuel_capacity']}</td></tr>
                                <tr> <td>Condition</td>      <td class='text-right'>{$FRAME[0]['condition']}</td></tr>
                                <tr> <td>Integrity</td>      <td class='text-right'>{$FRAME[0]['integrity']}</td></tr>
                                <tr> <td>Quality</td>        <td class='text-right'>{$FRAME[0]['quality']}</td></tr>
                                <tr> <td>Power Required</td> <td class='text-right'>{$FRAME[0]['power_required']}</td></tr>
                                <tr> <td>Crew Required</td>  <td class='text-right'>{$FRAME[0]['crew_required']}</td></tr>
                            </tbody>
                        </table>
                        <p class='text-white'>{$FRAME[0]['description']}</p>
                    </div>

                    <div class='panel'>
                        <h3 class=text-primary text-lg section-title'>Reactor</h3>
                        <table class='table'>
                            <tbody>
                                <tr> <td>Reactor</td>        <td class='text-right'>{$REACTR[0]['name']}</td></tr>
                                <tr> <td>Condition</td>      <td class='text-right'>{$REACTR[0]['condition']}</td></tr>
                                <tr> <td>Integrity</td>      <td class='text-right'>{$REACTR[0]['integrity']}</td></tr>
                                <tr> <td>Quality</td>        <td class='text-right'>{$REACTR[0]['quality']}</td></tr>
                                <tr> <td>Power Output</td>   <td class='text-right'>{$REACTR[0]['power_output']}</td></tr>
                                <tr> <td>Crew Required</td>  <td class='text-right'>{$REACTR[0]['crew_required']}</td></tr>
                            </tbody>
                        </table>
                        <p class='text-white'>{$REACTR[0]['description']}</p>
                    </div>

                    <div class='panel'>
                        <h3 class=text-primary text-lg section-title'>Engine</h3>
                        <table class='table'>
                            <tbody>
                                <tr> <td>Engine</td>         <td class='text-right'>{$ENGINE[0]['name']}</td></tr>
                                <tr> <td>Condition</td>      <td class='text-right'>{$ENGINE[0]['condition']}</td></tr>
                                <tr> <td>Integrity</td>      <td class='text-right'>{$ENGINE[0]['integrity']}</td></tr>
                                <tr> <td>Speed</td>          <td class='text-right'>{$ENGINE[0]['speed']}</td></tr>
                                <tr> <td>Quality</td>        <td class='text-right'>{$ENGINE[0]['quality']}</td></tr>
                                <tr> <td>Power Required</td> <td class='text-right'>{$ENGINE[0]['power_required']}</td></tr>
                                <tr> <td>Crew Required</td>  <td class='text-right'>{$ENGINE[0]['crew_required']}</td></tr>
                            </tbody>
                        </table>
                        <p class='text-white'>{$ENGINE[0]['description']}</p>
                    </div>

                    <div class='panel'>
                        <h3 class=text-primary text-lg section-title'>Crew</h3>
                        <table class='table'>
                            <tbody>
                                <tr> <td>Current</td>  <td class='text-right'>{$CREW[0]['current']} / {$CREW[0]['capacity']}</td></tr>
                                <tr> <td>Required</td> <td class='text-right'>{$CREW[0]['required']}</td></tr>
                                <tr> <td>Rotation</td> <td class='text-right'>{$CREW[0]['rotation']}</td></tr>
                                <tr> <td>Morale</td>   <td class='text-right'>{$CREW[0]['morale']}</td></tr>
                                <tr> <td>Wages</td>    <td class='text-right'>{$CREW[0]['wages']}</td></tr>
                            </tbody>
                        </table>
                    </div>
                </div>

                <div class='panel'>
                    <h3 class=text-primary text-lg section-title'>Modules Installed</h3>
                    <table class='table'>
                        <thead>
                            <th>Name</th>
                            <th>Capacity</th>
                            <th>Power Required</th>
                            <th>Crew Required</th>
                            <th>Slots Required</th>
                        </thead>
                        <tbody>
            ");
            foreach ($MODULE as $key => $val) {
                print("
                            <tr>
                                <td>{$val['name']}</td>
                                <td>{$val['capacity']}</td>
                                <td>{$val['power_required']}</td>
                                <td>{$val['crew_required']}</td>
                                <td>{$val['slots_required']}</td>
                            </tr>
                ");
            }
            print("
                        </tbody>
                    </table>
                </div>

                <br>
                <div class='panel'>
                    <h3 class=text-primary text-lg section-title'>Mounted Components</h3>
                    <table class='table'>
                        <thead>
                            <th>Name</th>
                            <th>Strength</th>
                            <th>Power Required</th>
                            <th>Crew Required</th>
                            <th>Slots Required</th>
                            <th>Deposits</th>
                        </thead>
                        <tbody>
            ");
            foreach ($MOUNTS as $key => $val) {
                $val['deposits'] = str_replace(',', '<br>', $val['deposits']);
                print("
                            <tr>
                                <td>{$val['name']}</td>
                                <td>{$val['strength']}</td>
                                <td>{$val['power_required']}</td>
                                <td>{$val['crew_required']}</td>
                                <td>{$val['slots_required']}</td>
                                <td>{$val['deposits']}</td>
                            </tr>
                ");
            }
            print("
                        </tbody>
                    </table>
                </div>
            ");

            break;


        case 'Cargo':
            print("You've discovered a placeholder! I'm still working on creating this...come back later!!");
            break;
    }
?>
