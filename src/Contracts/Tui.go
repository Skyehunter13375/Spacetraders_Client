package Contracts

import (
	"fmt"
	"strings"

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

func ContractView(width int) string {
	var sb strings.Builder
	CStruct := GetAllContracts()
	fmt.Fprintf(&sb, "ID:        %s\n", CStruct.Data[0].ID)
	fmt.Fprintf(&sb, "UpFront:   %d\n", CStruct.Data[0].Terms.Payment.OnAccepted)
	fmt.Fprintf(&sb, "OnCompl:   %d\n", CStruct.Data[0].Terms.Payment.OnFulfilled)
	fmt.Fprintf(&sb, "Accepted:  %t\n", CStruct.Data[0].Accepted)
	fmt.Fprintf(&sb, "Fulfilled: %t\n", CStruct.Data[0].Fulfilled)
	fmt.Fprintf(&sb, "Deadline:  %s\n", CStruct.Data[0].Terms.Deadline)
	fmt.Fprintf(&sb, "Expires:   %s\n", CStruct.Data[0].Expiration)
	fmt.Fprintf(&sb, "Material:  %d/%d (%s → %s)",
		CStruct.Data[0].Terms.Deliver[0].UnitsFulfilled,
		CStruct.Data[0].Terms.Deliver[0].UnitsRequired,
		CStruct.Data[0].Terms.Deliver[0].Material,
		CStruct.Data[0].Terms.Deliver[0].Destination,
	)

	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Padding(1, 2).
		Width(width / 2).
		BorderForeground(lipgloss.Color("240"))

	title := lipgloss.NewStyle().
		Foreground(lipgloss.Color("63")).
		Bold(true).
		Underline(true).
		Render(CStruct.Data[0].Faction + " | " + CStruct.Data[0].Type + " | " + CStruct.Data[0].DeadlineToAccept)

	return boxStyle.Render(title + "\n\n" + sb.String())
}
