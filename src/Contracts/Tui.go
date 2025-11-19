package Contracts

import (
	"fmt"
	"strconv"

	"github.com/rivo/tview"
)

func DisplayContract(data Contract) tview.Primitive {
	box := tview.NewForm()
	box.SetBorder(true)
	box.SetTitle("  " + data.ID + "  ")

	box.AddTextView("Contract:",  data.ID,       0, 1, true, true)
	box.AddTextView("Faction:",   data.Faction,  0, 1, true, true)
	box.AddTextView("Type:",      data.Type,     0, 1, true, true)
	box.AddTextView("Payment:",   strconv.Itoa(data.Terms.Payment.OnAccepted + data.Terms.Payment.OnFulfilled),  0, 1, true, true)
	box.AddTextView("Accepted:",  strconv.FormatBool(data.Accepted),  0, 1, true, true)
	box.AddTextView("Fulfilled:", strconv.FormatBool(data.Fulfilled), 0, 1, true, true)
	box.AddTextView("Expires:",   data.Expiration, 0, 1, true, true)
	box.AddTextView("Deadline:",  data.DeadlineToAccept, 0, 1, true, true)

	for i,v := range data.Terms.Deliver {
		box.AddTextView("Material " + strconv.Itoa(i+1) + ":", fmt.Sprintf("%d/%d | %s | %s", v.UnitsFulfilled, v.UnitsRequired, v.Material, v.Destination), 0, 1, true, true)
	}

	return box
}