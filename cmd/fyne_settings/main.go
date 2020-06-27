package main

import (
	fyne "github.com/wrzfeijianshen/fyne2"
	"github.com/wrzfeijianshen/fyne2/app"
	"github.com/wrzfeijianshen/fyne2/cmd/fyne_settings/settings"
	"github.com/wrzfeijianshen/fyne2/widget"
)

func main() {
	s := settings.NewSettings()

	a := app.New()
	w := a.NewWindow("Fyne Settings")

	appearance := s.LoadAppearanceScreen(w)
	tabs := widget.NewTabContainer(
		&widget.TabItem{Text: "Appearance", Icon: s.AppearanceIcon(), Content: appearance})
	tabs.SetTabLocation(widget.TabLocationLeading)
	w.SetContent(tabs)

	w.Resize(fyne.NewSize(480, 320))
	w.ShowAndRun()
}
