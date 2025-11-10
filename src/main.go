package main

import (
	"Spacetraders/src/Agents"
	"Spacetraders/src/Contracts"
	"Spacetraders/src/Fleet"
	"Spacetraders/src/Server"
	Waypoints "Spacetraders/src/Systems"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	menuItems     []string
	selectedIndex int
	content       string
	width, height int
	quitting      bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {

		case "right", "l":
			m.selectedIndex = (m.selectedIndex + 1) % len(m.menuItems)

		case "left", "h":
			m.selectedIndex = (m.selectedIndex - 1 + len(m.menuItems)) % len(m.menuItems)

		case "enter":
			item := m.menuItems[m.selectedIndex]
			switch item {
			case "Server":
				m.content = Server.ServerView(m.width)
			case "Agents":
				m.content = Agents.AgentView(m.width)
			case "Ships":
				Fleet.UpdateShipState()
				m.content = Fleet.ShipsView(m.width)
			case "Contracts":
				Contracts.NegotiateNewContract("NULL_SKY-1")
				Contracts.UpdateContracts()
				m.content = Contracts.ContractView(m.width)
			case "Systems":
				Waypoints.UpdateSystem("X1-XQ13")
				m.content = "Coming Soon..."
			case "Exit":
				m.quitting = true
				return m, tea.Quit
			default:
				m.content = "Coming Soon..."
			}

		case "q", "esc":
			m.quitting = true
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	if m.quitting {
		return ""
	}

	var menuParts []string
	for i, item := range m.menuItems {
		style := lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			Padding(0, 2).
			BorderForeground(lipgloss.Color("240")).
			Foreground(lipgloss.Color("250"))
		if i == m.selectedIndex {
			style = style.
				Foreground(lipgloss.Color("46")).
				BorderForeground(lipgloss.Color("46"))
		}
		menuParts = append(menuParts, style.Render(item))
	}

	menuBar := lipgloss.JoinHorizontal(lipgloss.Left, menuParts...)

	title := lipgloss.NewStyle().
		Foreground(lipgloss.Color("63")).
		Bold(true).
		Render("Welcome to Null Sky")

	footer := lipgloss.NewStyle().
		Foreground(lipgloss.Color("240")).
		Render("(← → navigate, Enter select, q to quit)")

	body := lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		menuBar,
		m.content,
		footer,
	)

	return lipgloss.Place(
		m.width, m.height,
		lipgloss.Left, lipgloss.Top,
		body,
	)
}

func main() {
	m := model{
		menuItems: []string{"Server", "Agents", "Ships", "Contracts", "Systems", "Exit"},
		content:   "Use ← → to navigate, Enter to select.",
	}

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error:", err)
	}
}
