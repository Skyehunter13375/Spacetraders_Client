package Settings

import "github.com/rivo/tview"
import "Spacetraders/src/General"

func ShowSettingsMenu(app *General.App) {
	app.UIState.SubMenu.Clear()
	app.UIState.Output.Clear()
	app.UIState.Output.AddItem(DisplaySettings(app), 0, 1, false)

	app.UIState.SubMenu.AddItem("Edit", "", 0, func() { app.UI.SetFocus(app.UIState.Output.GetItem(0)) })
	app.UIState.SubMenu.AddItem("Back", "", 0, func() { General.FocusMain(app) })
}

func DisplaySettings(app *General.App) *tview.Form {
	s, _:= General.GetConfig()

	data := tview.NewForm()
	data.SetBorder(true)
	data.SetTitle("  Settings  ")
	data.SetFieldBackgroundColor(app.State.Theme_BgBase)
	data.SetFieldTextColor(app.State.Theme_FgBase)
	data.SetButtonBackgroundColor(app.State.Theme_BgPanel)
	data.SetButtonTextColor(app.State.Theme_FgBase)

	data.AddInputField("Database Host:",  s.DB.Host, 30, nil, func(v string) { s.DB.Host = v })
	data.AddInputField("Database Name:",  s.DB.Name, 30, nil, func(v string) { s.DB.Name = v })
	data.AddInputField("Database User:",  s.DB.User, 30, nil, func(v string) { s.DB.User = v })
	data.AddInputField("Database Pass:",  s.DB.Pass, 30, nil, func(v string) { s.DB.Pass = v })
	data.AddInputField("Database SSL:",   s.DB.SSL,  30, nil, func(v string) { s.DB.SSL  = v })
	data.AddInputField("Database Type:",  s.DB.Type, 30, nil, func(v string) { s.DB.Type = v })

	data.AddTextView(  "──────────────", "", 0, 1, false, false)

	data.AddInputField("Account Token:",  "Hidden", 30, nil, func(v string) { s.Tokens.Account = v })
	data.AddInputField("Agent Token:",    "Hidden", 30, nil, func(v string) { s.Tokens.Agent   = v })

	data.AddTextView(  "──────────────", "", 0, 1, false, false)
	
	data.AddButton("Save", nil)
	data.AddButton("Cancel", func() { ShowSettingsMenu(app); app.UI.SetFocus(app.UIState.SubMenu) })

	return data
}

