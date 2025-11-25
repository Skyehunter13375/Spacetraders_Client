package General

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

