package Settings

import "github.com/rivo/tview"
import "Spacetraders/src/General"
import "github.com/gdamore/tcell/v2"

func DisplaySettings() tview.Primitive {
	s, _:= General.GetConfig()

	window := tview.NewFlex()
	
	data := tview.NewForm()
	data.SetBorder(true)
	data.SetTitle("  Settings  ")
	data.SetFieldBackgroundColor(tcell.ColorBlack)
	data.SetFieldTextColor(tcell.ColorTeal)

	data.AddInputField("Database Host:",  s.DB.Host, 50, nil, func(v string) { s.DB.Host = v })
	data.AddInputField("Database Name:",  s.DB.Name, 50, nil, func(v string) { s.DB.Name = v })
	data.AddInputField("Database User:",  s.DB.User, 50, nil, func(v string) { s.DB.User = v })
	data.AddInputField("Database Pass:",  s.DB.Pass, 50, nil, func(v string) { s.DB.Pass = v })
	data.AddInputField("Database SSL:",   s.DB.SSL,  50, nil, func(v string) { s.DB.SSL  = v })
	data.AddInputField("Database Type:",  s.DB.Type, 50, nil, func(v string) { s.DB.Type = v })

	data.AddTextView(  "──────────────", "", 0, 1, false, false)

	data.AddInputField("Account Token:",  "Hidden", 0, nil, func(v string) { s.Tokens.Account = v })
	data.AddInputField("Agent Token:",    "Hidden", 0, nil, func(v string) { s.Tokens.Agent   = v })

	window.AddItem(data, 0, 1, false)

	return window

}

