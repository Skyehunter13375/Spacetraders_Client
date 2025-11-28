package General

import "github.com/rivo/tview"
import "github.com/gdamore/tcell/v2"

// ────────────────────────────────────────────────────────────────────────────────
// APPLICATION STRUCTURES
// ────────────────────────────────────────────────────────────────────────────────
// Holds dynamic UI component references that screens will update.
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

type Configs struct {
	DB struct {
		Type string `yaml:"type"`
		Name string `yaml:"name"`
		Host string `yaml:"host"`
		Port int32  `yaml:"port"`
		User string `yaml:"user"`
		Pass string `yaml:"pass"`
		SSL  string `yaml:"SSL"`
	} `yaml:"database"`
	Tokens struct {
		Account string `yaml:"accnt"`
		Agent   string `yaml:"agent"`
	} `yaml:"tokens"`
}

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
