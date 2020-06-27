package widget_test

import (
	"testing"

	"github.com/wrzfeijianshen/fyne2"
	"github.com/wrzfeijianshen/fyne2/layout"
	"github.com/wrzfeijianshen/fyne2/test"
	"github.com/wrzfeijianshen/fyne2/theme"
	"github.com/wrzfeijianshen/fyne2/widget"
)

func TestCheck_Layout(t *testing.T) {
	test.NewApp()
	defer test.NewApp()
	test.ApplyTheme(t, theme.LightTheme())

	for name, tt := range map[string]struct {
		text     string
		checked  bool
		disabled bool
	}{
		"checked": {
			text:    "Test",
			checked: true,
		},
		"unchecked": {
			text: "Test",
		},
		"checked_disabled": {
			text:     "Test",
			checked:  true,
			disabled: true,
		},
		"unchecked_disabled": {
			text:     "Test",
			disabled: true,
		},
	} {
		t.Run(name, func(t *testing.T) {
			check := &widget.Check{
				Text:    tt.text,
				Checked: tt.checked,
			}
			if tt.disabled {
				check.Disable()
			}

			window := test.NewWindow(fyne.NewContainerWithLayout(layout.NewCenterLayout(), check))
			window.Resize(check.MinSize().Max(fyne.NewSize(150, 200)))

			test.AssertImageMatches(t, "check/layout_"+name+".png", window.Canvas().Capture())

			window.Close()
		})
	}
}
