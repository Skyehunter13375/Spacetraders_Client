<!DOCTYPE html>
<html lang="en">
    <head>
        <title>NULL SKY</title>
        <meta name="viewport">
        <meta charset="UTF-8">
        <meta content="width=device-width, initial-scale=1.0">
        <link rel="stylesheet" href="Style/style.css">
        <?php
            $thisFileName = basename($_SERVER['SCRIPT_FILENAME']);
            require_once("subroutines/Main.php");
        ?>
    </head>

    <body>
        <?php print_header($thisFileName); ?>
        <div class="container text-primary">
            <div class="panel">
                <h2 class="text-xl section-title">About Me</h2>
                <p style="color: var(--text-secondary);">
                    I am a Linux system administrator and software engineer with roughly a decade of experience working at the intersection of technical support and software development.<br>
                    Building tools to resolve common issues or prevent them from ever appearing in the first place is my passion.<br><br>

                    There is nothing more permanent than a temporary solution! I cannot accept the "slap a band-aid on it" mentality when a permanent solution is possible.<br>
                    I put a lot of personal time and focus on finding root cause for any particular issue and resolving it permanently, or implementing a solid mitigation process if a permanent solution is not yet possible.<br><br>

                    I am comfortable working independently or as part of a team, mentoring others, and helping translate technical complexity into actionable outcomes.<br>
                    I am also a stickler for documentation! I spent a great deal of my career writing training documentation, building classes for new hires, and mentoring teams from day one until they know the product better than I do.
                </p>
            </div>
        </div>
        <div class="container text-primary">
            <div class="panel">
                <h2 class="text-lx section-title">Links</h3>
                <div class="grid"><div class="panel" style="width: 250px;">LinkedIn: <a target="_blank" style="color: var(--text-secondary);" href=" https://www.linkedin.com/in/patrick-kelley-19490b132/">Patrick Kelley</a></div></div>
                <div class="grid"><div class="panel" style="width: 250px;">Github:   <a target="_blank" style="color: var(--text-secondary);" href="https://github.com/Skyehunter13375">Skyehunter13375</a></div></div>
            </div>
        </div>
    </body>
</html>


