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
    Theme_BgBase        tcell.Color
    Theme_BgPanel       tcell.Color
    Theme_BgAlt         tcell.Color
    Theme_FgBase        tcell.Color
    Theme_FgMuted       tcell.Color
    Theme_AccentBlue    tcell.Color
    Theme_AccentGreen   tcell.Color
    Theme_AccentYellow  tcell.Color
    Theme_AccentRed     tcell.Color
    Theme_AccentMagenta tcell.Color
}

// This now contains the app itself and all it's constituant components
type App struct {
	UI      *tview.Application
	UIState *UIState
	State   *GlobalState
}

