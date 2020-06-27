// Package main loads a very basic Hello World graphical application
package main

import (
	"net/url"

	"github.com/wrzfeijianshen/fyne2/app"
	"github.com/wrzfeijianshen/fyne2/widget"
)

func main() {
	a := app.New()
	tourURL, _ := url.Parse("https://tour.fyne.io")

	w := a.NewWindow("Hello")
	w.SetContent(widget.NewVBox(
		widget.NewLabel("Hello Fyne!"),
		widget.NewHyperlink("tour.fyne.io", tourURL),
	))

	w.ShowAndRun()
}
