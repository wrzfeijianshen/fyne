package widget_test

import (
	"testing"

	fyne "github.com/wrzfeijianshen/fyne2"
	"github.com/wrzfeijianshen/fyne2/sdk/widget"
	"github.com/wrzfeijianshen/fyne2/test"
	"github.com/wrzfeijianshen/fyne2/theme"

	"github.com/stretchr/testify/assert"
)

var shadowLevel = widget.ElevationLevel(5)

func TestShadow_ApplyTheme(t *testing.T) {
	app := test.NewApp()
	defer test.NewApp()
	app.Settings().SetTheme(theme.DarkTheme())

	s := widget.NewShadow(widget.ShadowAround, shadowLevel)
	w := test.NewWindow(s)
	defer w.Close()
	w.Resize(fyne.NewSize(50, 50))

	s.Resize(fyne.NewSize(30, 30))
	s.Move(fyne.NewPos(10, 10))
	test.AssertImageMatches(t, "shadow_theme_dark.png", w.Canvas().Capture())

	test.ApplyTheme(t, theme.LightTheme())
	test.AssertImageMatches(t, "shadow_theme_light.png", w.Canvas().Capture())

	test.ApplyTheme(t, theme.DarkTheme())
	test.AssertImageMatches(t, "shadow_theme_dark.png", w.Canvas().Capture())
}

func TestShadow_AroundShadow(t *testing.T) {
	app := test.NewApp()
	defer test.NewApp()
	app.Settings().SetTheme(theme.LightTheme())

	s := widget.NewShadow(widget.ShadowAround, shadowLevel)
	w := test.NewWindow(s)
	defer w.Close()
	w.Resize(fyne.NewSize(50, 50))

	s.Resize(fyne.NewSize(30, 30))
	s.Move(fyne.NewPos(10, 10))
	test.AssertImageMatches(t, "shadow_around.png", w.Canvas().Capture())
}

func TestShadow_Transparency(t *testing.T) {
	app := test.NewApp()
	defer test.NewApp()
	app.Settings().SetTheme(theme.LightTheme())

	s := widget.NewShadow(widget.ShadowAround, shadowLevel)
	w := test.NewWindow(s)
	defer w.Close()
	w.Resize(fyne.NewSize(50, 50))

	s.Resize(fyne.NewSize(40, 20))
	s.Move(fyne.NewPos(5, 15))
	s2 := widget.NewShadow(widget.ShadowAround, shadowLevel)
	w.Canvas().Overlays().Add(s2)
	s2.Resize(fyne.NewSize(20, 40))
	s2.Move(fyne.NewPos(15, 5))
	test.AssertImageMatches(t, "shadow_transparency.png", w.Canvas().Capture())
}

func TestShadow_BottomShadow(t *testing.T) {
	app := test.NewApp()
	defer test.NewApp()
	app.Settings().SetTheme(theme.LightTheme())

	s := widget.NewShadow(widget.ShadowBottom, shadowLevel)
	w := test.NewWindow(s)
	defer w.Close()
	w.Resize(fyne.NewSize(50, 50))

	s.Resize(fyne.NewSize(30, 30))
	s.Move(fyne.NewPos(10, 10))
	test.AssertImageMatches(t, "shadow_bottom.png", w.Canvas().Capture())
}

func TestShadow_MinSize(t *testing.T) {
	assert.Equal(t, fyne.NewSize(0, 0), widget.NewShadow(widget.ShadowAround, 1).MinSize())
}

func TestShadow_TopShadow(t *testing.T) {
	app := test.NewApp()
	defer test.NewApp()
	app.Settings().SetTheme(theme.LightTheme())

	s := widget.NewShadow(widget.ShadowTop, shadowLevel)
	w := test.NewWindow(s)
	defer w.Close()
	w.Resize(fyne.NewSize(50, 50))

	s.Resize(fyne.NewSize(30, 30))
	s.Move(fyne.NewPos(10, 10))
	test.AssertImageMatches(t, "shadow_top.png", w.Canvas().Capture())
}
