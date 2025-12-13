package Tui 

import "github.com/rivo/tview"
import "Spacetraders/src/Task"
import "Spacetraders/src/Model"

func ShowSettingsMenu(app *Model.App) {
	app.UIState.SubMenu.Clear()
	app.UIState.Output.Clear()
	app.UIState.Output.AddItem(DisplaySettings(app), 0, 1, false)
	app.UIState.SubMenu.AddItem("Edit", "", 0, func() { app.UI.SetFocus(app.UIState.Output.GetItem(0)) })
	app.UIState.SubMenu.AddItem("Back", "", 0, func() { Task.FocusMain(app) })
	app.UI.SetFocus(app.UIState.SubMenu)
}

func DisplaySettings(app *Model.App) *tview.Form {
	s, _:= Task.GetConfig()

	data := tview.NewForm()
	data.SetBorder(true)
	data.SetTitle("  Settings  ")
	data.SetBorderColor(Model.Theme.BgBorder)
	data.SetFieldBackgroundColor(Model.Theme.BgBase)
	data.SetFieldTextColor(Model.Theme.FgBase)
	data.SetButtonBackgroundColor(Model.Theme.BgBase)
	data.SetButtonTextColor(Model.Theme.FgBase)

	data.AddInputField("Database:",  s.DB.DbPath, 30, nil, func(v string) { s.DB.DbPath = v })
	// data.AddTextView(  "──────────────", "", 0, 1, false, false)
	
	data.AddInputField("Error Log:",    s.LOG.ErrPath, 30, nil, func(v string) { s.LOG.ErrPath = v })
	data.AddInputField("Activity Log:", s.LOG.ActPath, 30, nil, func(v string) { s.LOG.ActPath = v })
	// data.AddTextView(  "──────────────", "", 0, 1, false, false)

	data.AddInputField("Account Token:",  "Hidden", 30, nil, func(v string) { s.API.AccntToken = v })
	data.AddInputField("Agent Token:",    "Hidden", 30, nil, func(v string) { s.API.AgentToken = v })
	// data.AddTextView(  "──────────────", "", 0, 1, false, false)
	
	data.AddButton("Save", nil)
	data.AddButton("Cancel", func() { ShowSettingsMenu(app); app.UI.SetFocus(app.UIState.SubMenu) })

	return data
}

