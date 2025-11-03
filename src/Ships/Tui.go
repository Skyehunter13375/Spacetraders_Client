package Ships

import (
	"database/sql"
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func BuildShipBox(symbol string, width int) string {
	ship := GetShipState(symbol)

	// Build the ship’s info text
	info := fmt.Sprintf("Role:     %s\n", ship.Registration.Role)
	info = info + fmt.Sprintf("Status:   %s\n", ship.Nav.Status)
	info = info + fmt.Sprintf("Frame:    %s\n", ship.Frame.Name)
	info = info + fmt.Sprintf("Reactor:  %s\n", ship.Reactor.Name)
	info = info + fmt.Sprintf("Engine:   %s\n", ship.Engine.Name)
	info = info + fmt.Sprintf("Mode:     %s\n", ship.Nav.FlightMode)
	info = info + fmt.Sprintf("Waypoint: %s\n", ship.Nav.WaypointSymbol)
	info = info + fmt.Sprintf("Crew:     %d/%d (Req: %d)\n", ship.Crew.Current, ship.Crew.Capacity, ship.Crew.Required)
	info = info + fmt.Sprintf("Fuel:     %d/%d\n", ship.Fuel.Current, ship.Fuel.Capacity)
	info = info + fmt.Sprintf("Morale:   %d%%", ship.Crew.Morale)

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
	db, _ := sql.Open("sqlite3", "SpaceTraders.db")
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

		// Once a row hits 3 boxes or we reach the end, join and add the row
		if (i+1)%colsPerRow == 0 || i == len(symbols)-1 {
			rows = append(rows, lipgloss.JoinHorizontal(lipgloss.Top, currentRow...))
			currentRow = []string{}
		}
	}

	return lipgloss.JoinVertical(lipgloss.Left, rows...)
}
