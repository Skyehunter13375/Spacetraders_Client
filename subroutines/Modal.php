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

        case 'Cargo':
            printf("You asked for {$desire} data on {$symbol}");
            break;

        case 'Contract':
            printf("You asked for {$desire} data on {$symbol}");
            break;
    }
?>
