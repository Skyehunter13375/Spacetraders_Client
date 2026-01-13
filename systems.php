<!DOCTYPE html>
<html lang="en">
    <head>
        <title>NULL SKY</title>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <link rel="stylesheet" href="Style/style.css">
        <script src="https://cdn.jsdelivr.net/npm/chart.js"></script>
        <?php
            $thisFileName = basename($_SERVER['SCRIPT_FILENAME']);
            require_once("subroutines/Main.php");
            require_once("subroutines/Sql.php");
        ?>
    </head>

    <body>
        <?php print_header($thisFileName); ?>
        <div class="min-h-screen text-primary">
            <div class="container">
                <?php
                    $WAYPOINTS = SELECT('SELECT symbol AS id,x,y FROM waypoints'); 
                    $jsonData = json_encode($WAYPOINTS);
                ?>
                <div id="orbitChart" style="width:900px; height: 900px;"></div>
            </div>
        </div>
    </body>

    <script src="https://cdn.plot.ly/plotly-2.30.0.min.js"></script>
    <script>
        const planets = <?php echo $jsonData; ?>;

        // Calculate max radii so the data is plotted correctly
        const radii = planets.map(p =>
          Math.sqrt(p.x * p.x + p.y * p.y)
        );

        const maxRadius = Math.max(...radii);
        const padding = 50;
        const chartLimit = maxRadius + padding;

        // Only build unique orbit radii lines
        const uniqueRadii = [...new Set(radii.map(r => Math.round(r)))];
        const orbits = uniqueRadii.map(r => ({
          type: 'circle',
          x0: -r,
          y0: -r,
          x1: r,
          y1: r,
          line: { dash: 'dot', width: 1, color: 'rgba(255,255,255,0.3)' }
        }));

        const trace = {
          x:            planets.map(p => p.x),
          y:            planets.map(p => p.y),
          text:         planets.map(p => p.id),
          mode:         'markers+text',
          textposition: 'top center',
          type:         'scatter',
          marker:       { size: 10 },
          customdata:   planets
        };

        const layout = {
          paper_bgcolor: 'rgba(38, 38, 38, 0.7)',
          plot_bgcolor:  'rgba(38, 38, 38, 0.7)',

          xaxis: {
            range: [-chartLimit, chartLimit],
            zeroline: true,
            scaleanchor: 'y'
          },
          yaxis: {
            range: [-chartLimit, chartLimit],
            zeroline: true
          },
          shapes: orbits,
          showlegend: false,
        };

        Plotly.newPlot('orbitChart', [trace], layout);
    </script>
</html>

