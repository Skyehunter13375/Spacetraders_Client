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
    // First log cloudflared stuff
//     $cfd_detail = getallheaders();
//     $LogString  = "Page: {$current}|<>|From: {$cfd_detail['X-Forwarded-For']}|<>|Proto: {$cfd_detail['X-Forwarded-Proto']}|<>|OS: {$cfd_detail['Sec-Ch-Ua-Platform']}|<>|Mobile: {$cfd_detail['Sec-Ch-Ua-Mobile']}|<>|Region: {$cfd_detail['Cf-Ipcountry']}|<>|Agent: {$cfd_detail['User-Agent']}";
//     error_log($LogString);

    // Now actually generate the header
    $baseUrl   = "http://10.10.1.26/SpaceTraders/Front";
    $currClass = "class='block py-2 px-3 rounded-sm text-green-600 bg-transparent md:bg-transparent md:text-green-600 md:p-0' aria-current='page'";
    $notCurrnt = "class='block py-2 px-3 rounded-sm text-white bg-transparent hover:bg-neutral-800 md:hover:bg-transparent md:border-0 md:hover:text-green-800 md:p-0'";
    $menuItems = [
        'index.php'       => 'Home',
        'fleet.php'       => 'Fleet',
        'contracts.php'   => 'Contracts',
        'systems.php'     => 'Systems',
        'leaderboard.php' => 'Leaderboard',
        'buildlog.php'    => 'Build Log'
    ];

    print("
        <div class='flex-full flex-wrap items-center justify-center'>
            <nav class='bg-neutral-800 bg-opacity-70 text-white font-mono'>
                <div class='max-w-screen-xl flex flex-wrap items-center justify-between mx-auto p-4'>
                    <a href='{$baseUrl}/index.php' class='flex items-left space-x-3 rtl:space-x-reverse'>
                        <img src='images/NullSkyeLight.png' class='h-15 w-15' alt='Logo' />
                        <!-- <span class='text-white text-2xl font-semibold whitespace-nowrap'>NULL SKYE</span> -->
                    </a>
                    <button data-collapse-toggle='navbar-default' type='button' class='inline-flex p-2 w-10 h-10 text-sm text-gray-500 rounded-lg md:hidden hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200' aria-controls='navbar-default' aria-expanded='false'>
                        <span class='sr-only'>Open main menu</span>
                        <svg class='w-5 h-5' aria-hidden='true' xmlns='http://www.w3.org/2000/svg' fill='none' viewBox='0 0 17 14'><path stroke='currentColor' stroke-linecap='round' stroke-linejoin='round' stroke-width='2' d='M1 1h15M1 7h15M1 13h15'/></svg>
                    </button>
                    <div class='hidden w-full md:block md:w-auto' id='navbar-default'>
                        <ul class='font-medium flex flex-col p-4 md:p-0 mt-4 md:flex-row md:space-x-8 text-green-400 font-mono'>

    ");

    // Generate menu items dynamically
    foreach ($menuItems as $file => $title) {
        $class = ($current === $file) ? $currClass : $notCurrnt;
        printf("<li><a href='%s%s' %s>%s</a></li>", $baseUrl."/", $file, $class, $title);
    }

    print("
                        </ul>
                    </div>
                </div>
            </nav>
        </div>
        <div class='w-full flex justify-center'>
            <hr class='w-full border-green-600'>
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
