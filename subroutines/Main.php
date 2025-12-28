<?php
// <------------------------------------------------------------------------------------------------------------------------------>
// Purpose: Define basic web UI functionality
// Changes:
//     - 12/26/2025 - Simplified for new data structure
//     - 04/20/2025 - Program created
// <------------------------------------------------------------------------------------------------------------------------------>
error_reporting(E_ALL);
ini_set('display_errors', 1);
ini_set('log_errors', 1);

function print_header($current) {
    $baseUrl   = "http://10.10.1.26/SpaceTraders";
    $currClass = "class='nav-link nav-link-active'";
    $notCurrnt = "class='nav-link'";
    $menuItems = [
        'index.php'       => 'Home',
        'fleet.php'       => 'Fleet',
        'contracts.php'   => 'Contracts',
        'systems.php'     => 'Systems',
        'changelog.php'   => 'Changelog',
        'aboutme.php'     => 'AboutMe'
    ];

    print("
        <div class='container flex'>
            <nav>
                <div class='flex' style='padding-top: 1rem; margin-bottom: -10px'>
                    <a href='{$baseUrl}/index.php'><img src='Media/images/NullSkyLight.png' style='width: 3.75rem; height: 3.75rem;' alt='Logo'></a>
                    <div class='hidden md:block'>
                        <ul class='flex md:flex-row md:space-x-8'>
    ");

    foreach ($menuItems as $file => $title) {
        $class = ($current === $file) ? $currClass : $notCurrnt;
        $fpath = $baseUrl . "/" . $file;
        print("<a href='{$fpath}' {$class}>{$title}</a>");
    }

    print("
                        </ul>
                    </div>
                </div>
            </nav>
        </div>
    ");
}


function create_table($HEADERS, $DATA) {
    print("
        <table class='w-full border-collapse border border-gray-700'>
            <thead>
                <tr class='bg-purple-700 text-white'>
    ");
    foreach ($HEADERS as $key) {
        print("     <th class='border border-gray-600 px-4 py-2 text-left whitespace-nowrap'>{$key}</th>");
    }

    print("
                </tr>
            </thead>
            <tbody>
    ");

    foreach ($DATA as $dKey => $dVals) {
        print("<tr class='odd:bg-gray-700 even:bg-gray-800'>");
        foreach ($HEADERS as $hKey) {
            print("<td class='border border-gray-600 px-4 py-2'>{$dVals[$hKey]}</td>");
        }
        print("</tr>");
    }

    print("
                </tr>
            </tbody>
        </table>
    ");
}

function DEBUG_printr($array = []) {
    print("<pre style='text-align: left;'>");
    print_r($array);
    print("</pre>");
}
?>
