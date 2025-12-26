package Model

import "github.com/rivo/tview"
import "github.com/gdamore/tcell/v2"

type GlobalTheme struct {
	BgBase        tcell.Color
	BgBorder      tcell.Color
	FgBase        tcell.Color
	FgMuted       tcell.Color
	AccentBlue    tcell.Color
	AccentGreen   tcell.Color
	AccentYellow  tcell.Color
	AccentRed     tcell.Color
	AccentMagenta tcell.Color
}

var Theme = GlobalTheme {
	BgBase:        tcell.GetColor("#2A2E2A"),
	BgBorder:      tcell.GetColor("#556B2F"),
	FgBase:        tcell.GetColor("#D0D2C9"),
	FgMuted:       tcell.GetColor("#7C8573"),
	AccentBlue:    tcell.GetColor("#4A6068"),
	AccentGreen:   tcell.GetColor("#6BA64B"),
	AccentYellow:  tcell.GetColor("#C2A447"),
	AccentRed:     tcell.GetColor("#B0493C"),
	AccentMagenta: tcell.GetColor("#8B4A5A"),
}

type UIState struct {
	MainMenu *tview.List
	SubMenu  *tview.List
	Output   *tview.Flex
}

// Global data shared across screens (expand over time).
type GlobalState struct {
}

// This now contains the app itself and all it's constituant components
type App struct {
	UI      *tview.Application
	UIState *UIState
	State   *GlobalState
}
