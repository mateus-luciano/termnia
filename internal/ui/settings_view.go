package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func SettingsView(w fyne.Window, app fyne.App) fyne.CanvasObject {
	content := container.NewVBox(
		canvas.NewText("Settings", color.White),
		widget.NewSeparator(),
	)

	return content
}
