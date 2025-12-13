package Task

import "github.com/rivo/tview"
import "github.com/gdamore/tcell/v2"
import "strings"
import "fmt"
import "Spacetraders/src/Model"

func FocusMain(app *Model.App) {
	ui := app.UIState
	ui.SubMenu.Clear()
	ui.Output.Clear()
	app.UI.SetFocus(app.UIState.MainMenu)
}

func FocusSub(app *Model.App) {
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
	frame.SetBorderColor(Model.Theme.BgBorder)

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

