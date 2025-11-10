package Fleet

import (
	"Spacetraders/src/General"
	"database/sql"
	"fmt"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/lipgloss"
)

func ProgressBar(value, max int) string {
	bar := progress.New(
		progress.WithGradient("#045f04ff", "#28b328ff"),
	)
	bar.Width = 30

	var ratio float64
	if max > 0 {
		ratio = float64(value) / float64(max)
	} else {
		ratio = 1
	}

	// Clamp between 0 and 1 just in case
	if ratio < 0 {
		ratio = 0
	} else if ratio > 1 {
		ratio = 1
	}

	return bar.ViewAs(ratio)
}

func BuildShipBox(symbol string, width int) string {
	ship := GetShipState(symbol)

	crewBar := ProgressBar(ship.Crew.Current, ship.Crew.Capacity)
	fuelBar := ProgressBar(ship.Fuel.Current, ship.Fuel.Capacity)
	moraleBar := ProgressBar(ship.Crew.Morale, 100)

	info := fmt.Sprintf("Role:     %s\n", ship.Registration.Role)
	info += fmt.Sprintf("Status:   %s\n", ship.Nav.Status)
	info += fmt.Sprintf("Frame:    %s\n", ship.Frame.Name)
	info += fmt.Sprintf("Reactor:  %s\n", ship.Reactor.Name)
	info += fmt.Sprintf("Engine:   %s\n", ship.Engine.Name)
	info += fmt.Sprintf("Mode:     %s\n", ship.Nav.FlightMode)
	info += fmt.Sprintf("Waypoint: %s\n", ship.Nav.WaypointSymbol)
	info += fmt.Sprintf("Crew:     %d/%d (Req: %d)\n%s\n", ship.Crew.Current, ship.Crew.Capacity, ship.Crew.Required, crewBar)
	info += fmt.Sprintf("Fuel:     %d/%d\n%s\n", ship.Fuel.Current, ship.Fuel.Capacity, fuelBar)
	info += fmt.Sprintf("Morale:   %d%%\n%s", ship.Crew.Morale, moraleBar)

	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Padding(1, 2).
		Width(width).
		BorderForeground(lipgloss.Color("240"))

	title := lipgloss.NewStyle().
		Foreground(lipgloss.Color("63")).
		Bold(true).
		Underline(true).
		Render(symbol)

	return boxStyle.Render(title + "\n\n" + info)
}

func ShipsView(width int) string {
	db, err := sql.Open("postgres", "user=skyehunter dbname=spacetraders sslmode=disable")
	if err != nil {
		General.LogErr(fmt.Sprintf("DB open failed: %v", err))
	}
	defer db.Close()
	ships, _ := db.Query(`SELECT symbol FROM ship`)
	defer ships.Close()

	var symbols []string
	for ships.Next() {
		var s string
		ships.Scan(&s)
		symbols = append(symbols, s)
	}

	const colsPerRow = 3
	boxWidth := (width / colsPerRow) - 2

	var rows []string
	var currentRow []string

	for i, symbol := range symbols {
		box := BuildShipBox(symbol, boxWidth)
		currentRow = append(currentRow, box)

		if (i+1)%colsPerRow == 0 || i == len(symbols)-1 {
			rows = append(rows, lipgloss.JoinHorizontal(lipgloss.Top, currentRow...))
			currentRow = []string{}
		}
	}

	return lipgloss.JoinVertical(lipgloss.Left, rows...)
}
