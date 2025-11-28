package General

import "github.com/rivo/tview"
import "github.com/gdamore/tcell/v2"

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
	frame.SetBorders(1, 1, 1, 1, 1, 1)
	frame.SetBorder(true)
	if title != "" {
		frame.SetTitle(" " + title + " ")
	}
	return &CardButton{
		Frame:    frame,
		selected: false,
		onSelect: onSelect,
	}
}

func (c *CardButton) Draw(screen tcell.Screen) {
	if c.selected {
		c.Frame.SetBorderColor(tcell.ColorYellow)
	} else {
		c.Frame.SetBorderColor(tcell.ColorWhite)
	}
	c.Frame.Draw(screen)
}

func (c *CardButton) InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
	return func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
		switch event.Key() {
		case tcell.KeyEnter:
			if c.onSelect != nil {
				c.onSelect()
			}
			return
		}
		if h := c.Frame.InputHandler(); h != nil {
			h(event, setFocus)
		}
	}
}

func (c *CardButton) Focus(delegate func(p tview.Primitive)) {
	c.selected = true
	delegate(c.Frame)
}

func (c *CardButton) Blur() {
	c.selected = false
}

