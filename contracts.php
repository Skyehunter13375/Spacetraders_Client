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
        <?php 
            print_header($thisFileName); 
            $CONTRACTS = SELECT("SELECT id,faction,type,pay_on_accept,pay_on_complete,accepted,fulfilled,deadline,expiration,deadline_to_accept FROM contracts");
        ?>
        <div class="min-h-screen text-primary">
            <div class="container">
                <div class="grid grid-4">
                    <?php
                        foreach ($CONTRACTS as $key => $vals) {
                            print("
                                <div class='panel'>
                                    <h3 class='text-lg section-title'>{$vals['id']}</h3>
                                    <table class='table'>
                                        <tbody>
                                            <tr> <td>Type:</td>  <td class='text-right'>{$vals['type']}</td></tr>
                                            <tr> <td>Upfront Pay:</td>  <td class='text-right'>{$vals['pay_on_accept']}</td></tr>
                                            <tr> <td>Complete Pay:</td>  <td class='text-right'>{$vals['pay_on_complete']}</td></tr>
                                            <tr> <td>Deadline:</td>  <td class='text-right'>{$vals['deadline']}</td></tr>
                                            <tr> <td>Expiration:</td>  <td class='text-right'>{$vals['deadline_to_accept']}</td></tr>
                                        </tbody>
                                    </table>
                                </div>
                            ");
                        }


                        $remainder = count($CONTRACTS) % 4;
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
