package Settings

import "github.com/rivo/tview"
import "Spacetraders/src/General"

func ShowSettingsMenu(app *General.App) {
	app.UIState.SubMenu.Clear()
	app.UIState.Output.Clear()
	app.UIState.Output.AddItem(DisplaySettings(app), 0, 1, false)
	app.UIState.SubMenu.AddItem("Edit", "", 0, func() { app.UI.SetFocus(app.UIState.Output.GetItem(0)) })
	app.UIState.SubMenu.AddItem("Back", "", 0, func() { General.FocusMain(app) })
	app.UI.SetFocus(app.UIState.SubMenu)
}

func DisplaySettings(app *General.App) *tview.Form {
	s, _:= General.GetConfig()

	data := tview.NewForm()
	data.SetBorder(true)
	data.SetTitle("  Settings  ")
	data.SetBorderColor(General.Theme.BgBorder)
	data.SetFieldBackgroundColor(General.Theme.BgBase)
	data.SetFieldTextColor(General.Theme.FgBase)
	data.SetButtonBackgroundColor(General.Theme.BgBase)
	data.SetButtonTextColor(General.Theme.FgBase)

	data.AddInputField("Database:",  s.DB.DbPath, 30, nil, func(v string) { s.DB.DbPath = v })

	data.AddTextView(  "──────────────", "", 0, 1, false, false)

	data.AddInputField("Account Token:",  "Hidden", 30, nil, func(v string) { s.API.AccntToken = v })
	data.AddInputField("Agent Token:",    "Hidden", 30, nil, func(v string) { s.API.AgentToken = v })

	data.AddTextView(  "──────────────", "", 0, 1, false, false)
	
	data.AddButton("Save", nil)
	data.AddButton("Cancel", func() { ShowSettingsMenu(app); app.UI.SetFocus(app.UIState.SubMenu) })

	return data
}

