package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func AboutView() fyne.CanvasObject {
	content := container.NewVBox(
		widget.NewLabelWithStyle("About Termnia", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewSeparator(),
		widget.NewLabel("version: 0.1.0"),
	)

	return container.NewCenter(content)
}
