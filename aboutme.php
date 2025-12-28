<!DOCTYPE html>
<html lang="en">
    <head>
        <title>NULL SKYE</title>
        <meta   name="viewport">
        <meta   charset="UTF-8">
        <meta   content="width=device-width, initial-scale=1.0">
        <link   rel="stylesheet" href="style.css">
        <script src="javascript/marked.min.js"></script>
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
        <div class="container px-4 py-8 space-y-8">
            <div class="panel text-white">
                <div id="preview"></div>
            </div>
        </div>
        <script src="https://unpkg.com/flowbite@1.6.5/dist/flowbite.min.js"></script>
    </body>

    <script>
        const preview = document.getElementById("preview");
        async function loadMarkdown() {
            const response = await fetch("Logs/ABOUTME.md");
            const markdown = await response.text();
            preview.innerHTML = marked.parse(markdown);
        }
        loadMarkdown();
    </script>
</html>


