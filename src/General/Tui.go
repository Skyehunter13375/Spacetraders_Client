package General

import "github.com/rivo/tview"
import "github.com/gdamore/tcell/v2"
import "strings"
import "fmt"

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
func FocusMain(app *App) {
	ui := app.UIState
	ui.SubMenu.Clear()
	ui.Output.Clear()
	app.UI.SetFocus(app.UIState.MainMenu)
}

func FocusSub(app *App) {
	ui := app.UIState
	ui.Output.Clear()
	app.UI.SetFocus(app.UIState.MainMenu)
}

type CardButton struct {
	*tview.Frame
	selected bool
	onSelect func()
}

func NewCardButton(inner tview.Primitive, title string, onSelect func()) *CardButton {
	frame := tview.NewFrame(inner)
	frame.SetBorder(true)
	frame.SetBorders(0, 0, 0, 0, 0, 0)
	frame.SetBorderColor(Theme.BgBorder)

	if title != "" {frame.SetTitle(" " + title + " ")}

	return &CardButton{
		Frame:    frame,
		selected: false,
		onSelect: onSelect,
	}
}

func (c *CardButton) InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
	return func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
		switch event.Key() {
		case tcell.KeyEnter:
			if c.onSelect != nil {c.onSelect()}
			return
		}
		if h := c.Frame.InputHandler(); h != nil {h(event, setFocus)}
	}
}

func (c *CardButton) Focus(delegate func(p tview.Primitive)) {
	c.selected = true
	delegate(c.Frame)
}

func (c *CardButton) Blur() {
	c.selected = false
}

func ProgressBar(curr, req, max int) string {
	if max <= 0 { return "N/A" }

	var bar strings.Builder
	totalWidth := 20

	// Clamp values so we don't go below 0 or above max
	if curr < 0   { curr = 0 }
	if curr > max { curr = max }
	if req  < 0   { req = 0 }
	if req  > max { req = max }

	// Compute percentages
	percent := int(float64(curr) / float64(max) * 100)
	filled := (percent * totalWidth) / 100
	reqPos := (int(float64(req) / float64(max) * float64(totalWidth)))

	// Build the bar
	for i := range totalWidth {
		switch {
			case i < filled:
				bar.WriteString("■")
			case i == reqPos:
				bar.WriteString("|") // marker for required threshold
			default:
				bar.WriteString(" ")
			}
	}

	return fmt.Sprintf("[[green]%s[-]] %d%%", bar.String(), percent)
}
