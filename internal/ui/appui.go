package ui

import (
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/mateus-luciano/termnia/internal/config"
)

func Run() error {
	a := app.NewWithID(config.AppID)

	switch config.Get().Theme {
	case "dark":
		a.Settings().SetTheme(theme.DarkTheme())
	default:
		a.Settings().SetTheme(theme.LightTheme())
	}

	w := a.NewWindow("Termnia")
	w.Resize(fyne.NewSize(1000, 650))

	items := []string{"Terminal", "Settings", "About"}
	contentStack := container.NewStack()

	// term := NewTerminalView(w)
	settings := NewSettingsView(a, w)
	about := AboutView()

	list := container.NewVBox()
	for i, item := range items {
		index := i
		btn := widget.NewButton(item, func() {
			switch index {
			case 0:
				contentStack.Objects = []fyne.CanvasObject{}
			case 1:
				contentStack.Objects = []fyne.CanvasObject{settings}
			case 2:
				contentStack.Objects = []fyne.CanvasObject{about}
			}
			contentStack.Refresh()
		})
		list.Add(btn)
	}

	title := canvas.NewText("Termnia", color.NRGBA{R: 255, G: 255, B: 255, A: 255})
	title.TextSize = 18
	top := container.NewVBox(title, canvas.NewLine(color.NRGBA{A: 40}))
	root := container.NewBorder(top, nil, list, nil, contentStack)

	w.SetContent(root)

	w.SetOnClosed(func() {
		log.Println("director window closed")
	})

	w.ShowAndRun()

	return nil
}

func NewSettingsView(a fyne.App, w fyne.Window) fyne.CanvasObject {
	return container.NewVBox(
		canvas.NewText("Settings", color.NRGBA{R: 50, G: 50, B: 50, A: 255}),
	)
}
