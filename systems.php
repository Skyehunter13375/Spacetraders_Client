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
                <div class="panel" style="max-width: 600px; max-height: 600px;">
                    <?php
                        $WAYPOINTS = SELECT('SELECT symbol,x,y FROM waypoints'); 
                        $jsonData = json_encode($WAYPOINTS);
                    ?>
                    <canvas id="myChart" style='width: 250px; height: 250px;'></canvas>
                </div>
            </div>
        </div>
    </body>

    <script>
        // Create plugin for building orbital path tracking
        const orbitPlugin = {
            id: 'orbitPlugin',
            beforeDatasetsDraw(chart) {
                const { ctx, scales } = chart;
                const xScale = scales.x;
                const yScale = scales.y;

                // Pixel location of origin
                const cx = xScale.getPixelForValue(0);
                const cy = yScale.getPixelForValue(0);

                ctx.save();
                ctx.strokeStyle = 'rgba(255,255,255,0.15)';
                ctx.lineWidth = 1;

                xyValues.forEach(p => {
                    if (p.x === 0 && p.y === 0) return;

                    // Pixel location of the planet
                    const px = xScale.getPixelForValue(p.x);
                    const py = yScale.getPixelForValue(p.y);

                    // TRUE pixel-space Euclidean radius
                    const radius = Math.hypot(px - cx, py - cy);

                    ctx.beginPath();
                    ctx.arc(cx, cy, radius, 0, Math.PI * 2);
                    ctx.stroke();
                });

                ctx.restore();
            }
        };

        // Get data from PHP array
        const xyValues = <?php echo $jsonData; ?>;

        // Store data for concentric orbital data
        const orbitRadii = [...new Set(
            xyValues
                .map(p => Math.sqrt(p.x * p.x + p.y * p.y))
                .filter(r => r > 0)
        )].sort((a, b) => a - b);

        // Create the chart using the data captured above
        const maxOrbitRadius = Math.max(...orbitRadii);
        document.addEventListener('DOMContentLoaded', function () {
            const ctx = document.getElementById('myChart').getContext('2d');
            new Chart(ctx, {
                type: 'scatter',
                plugins: [orbitPlugin],
                data: {
                    datasets: [{
                        data: xyValues,
                        pointRadius: 5,
                        pointBackgroundColor: '#4ade80'
                    }]
                },
                options: {
                    responsive: true,
                    plugins: {
                        legend:  { display: false },
                        tooltip: { enabled: true }
                    },
                    scales: {
                        x: {
                            display: false,
                            suggestedMin: -maxOrbitRadius,
                            suggestedMax:  maxOrbitRadius
                        },
                        y: {
                            display: false,
                            suggestedMin: -maxOrbitRadius,
                            suggestedMax:  maxOrbitRadius
                        }
                    }
                }
            });
        });
    </script>
</html>

