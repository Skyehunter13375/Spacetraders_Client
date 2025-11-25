package General

import "github.com/rivo/tview"

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
	// Ships map[string]*Fleet.Ship
	// Systems map[string]*System
	// etc.
}

// This now contains the app itself and all it's constituant components
type App struct {
	UI      *tview.Application
	UIState *UIState
	State   *GlobalState
}

