package main

import (
	"Spacetraders/src/Agent"
	"Spacetraders/src/Server"
	"Spacetraders/src/Ships"
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func handleMenuSelection(choice string, app *tview.Application) tview.Primitive {
	switch choice {
	case "Game Status":
		return Server.DisplayGameServerState()

	case "Agent Status":
		return Agent.DisplayAgentState()

	case "Ships":
		return Ships.DisplayShipState()

	case "Exit":
		app.Stop()
		return nil

	default:
		emptyBox := tview.NewTextView()
		fmt.Fprintf(emptyBox, "You selected the menu: %s\n", choice)
		return emptyBox
	}
}

func main() {
	// ┣━━━━━━━━━━━━━━━━━━━━━━┫ Pre-Define Default Colorscheme ┣━━━━━━━━━━━━━━━━━━━━━━┫
	app := tview.NewApplication()
	tview.Styles.PrimitiveBackgroundColor = tcell.ColorBlack
	tview.Styles.ContrastBackgroundColor = tcell.ColorBlack
	tview.Styles.MoreContrastBackgroundColor = tcell.ColorBlack
	tview.Styles.BorderColor = tcell.ColorGray
	tview.Styles.TitleColor = tcell.ColorWhite
	tview.Styles.GraphicsColor = tcell.ColorGray
	tview.Styles.PrimaryTextColor = tcell.ColorWhite
	tview.Styles.SecondaryTextColor = tcell.ColorGray
	tview.Styles.TertiaryTextColor = tcell.ColorGray
	tview.Styles.InverseTextColor = tcell.ColorGray
	tview.Styles.ContrastSecondaryTextColor = tcell.ColorGray

	// ┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫ Build Menu Bar ┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
	menuItems := []string{"Game Status", "Agent Status", "Ships", "Exit"}
	menuBar := tview.NewFlex().SetDirection(tview.FlexColumn)
	menuBar.SetBorder(false)
	menuBar.SetTitle("Welcome to Null Sky")
	menuBar.SetTitleAlign(0)

	buttons := []*tview.Button{}
	for _, item := range menuItems {
		btn := tview.NewButton(item)
		btn.SetBorder(true)
		buttons = append(buttons, btn)
		if item == "Exit" {
			spacer := tview.NewBox()
			menuBar.AddItem(spacer, 0, 1, false)
		}
		menuBar.AddItem(btn, len(item)+4, 0, false)
	}

	selected := 0
	highlight := func(index int) {
		for i, b := range buttons {
			if i == index {
				b.SetLabelColor(tcell.ColorGreen)
				b.SetBorderColor(tcell.ColorGreen)
			} else {
				b.SetLabelColor(tcell.ColorWhite)
				b.SetBorderColor(tcell.ColorGray)
			}
		}
	}
	highlight(selected)

	// ┣━━━━━━━━━━━━━━━━━━━━━━━━━┫ Building the data field ┣━━━━━━━━━━━━━━━━━━━━━━━━━━┫
	content := tview.NewFlex().SetDirection(tview.FlexRow)
	content.SetBorder(false)

	// ┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫ Define the layout ┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
	layout := tview.NewFlex().SetDirection(tview.FlexRow)
	layout.AddItem(menuBar, 3, 0, false)
	layout.AddItem(content, 0, 1, false)

	app.SetRoot(layout, true)

	// ┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫ Input Key Handler ┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyRight:
			selected = (selected + 1) % len(buttons)
			highlight(selected)

		case tcell.KeyLeft:
			selected = (selected - 1 + len(buttons)) % len(buttons)
			highlight(selected)

		case tcell.KeyEnter:
			content.Clear()
			content_data := handleMenuSelection(menuItems[selected], app)
			if content_data != nil {
				content.AddItem(content_data, 0, 1, false)
			}

		case tcell.KeyESC:
			app.Stop()
		}
		return nil
	})

	// ┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫ Display UI ┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
	if err := app.EnableMouse(false).Run(); err != nil {
		panic(err)
	}
}
