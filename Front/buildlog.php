<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>NULL SKYE</title>
        <script src="https://unpkg.com/@tailwindcss/browser@4"></script>
        <script src="https://unpkg.com/flowbite@1.6.5/dist/flowbite.min.js"></script>
        <?php $thisFileName = basename($_SERVER['SCRIPT_FILENAME']); ?>
        <?php 
            require_once("subroutines/Main.php"); 
        ?>
    </head>

    <body>
        <?php print_header($thisFileName); ?>
        <div class="min-h-screen bg-neutral-800 bg-opacity-70 text-white font-mono text-sm">
            <div class="w-full mx-auto px-4 py-8">
                <div class="flex flex-wrap gap-4 justify-center">
                    <div class="w-7/8 bg-neutral-800 rounded-lg shadow-md p-4 mb-4 border border-green-800">
                        <?php print(nl2br(file_get_contents("../Logs/CHANGELOG.md"))); ?>
                    </div>
                </div>
            </div>
        </div>
        <script src="https://unpkg.com/flowbite@1.6.5/dist/flowbite.min.js"></script>
    </body>
</html>
