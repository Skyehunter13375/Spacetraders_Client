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
        <div class="min-h-screen text-primary">
            <div class="container px-4 py-8 space-y-8">
                <div class="grid">
                    <div class="panel text-white">
                        <?php print(nl2br(file_get_contents("Logs/CHANGELOG.md"))); ?>
                    </div>
                </div>
            </div>
        </div>
        <script src="https://unpkg.com/flowbite@1.6.5/dist/flowbite.min.js"></script>
    </body>
</html>
