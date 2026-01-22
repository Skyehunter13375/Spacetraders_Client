<?php
    $thisFileName = basename($_SERVER['SCRIPT_FILENAME']);
    require_once("Main.php");
    require_once("Sql.php");
    $symbol = $_GET['symbol'];
    $desire = $_GET['type'];

    switch ($desire) {
        case 'System':
            $WAYP = SELECT("SELECT system,symbol,type,x,y,orbits,construction FROM waypoints WHERE system = 'X1-TT41' ORDER BY type");
            print("
                        <table class='table'>
                            <tbody>
                                <tr>
                                    <td>System</td>
                                    <td>Symbol</td>
                                    <td>Type</td>
                                    <td>X</td>
                                    <td>Y</td>
                                    <td>Orbits</td>
                                    <td>Construction</td>
                                </tr>
            ");
            foreach ($WAYP as $key => $val) {
                print("
                                <tr>
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
