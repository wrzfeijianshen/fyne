package gomobile

import (
	"github.com/wrzfeijianshen/fyne2"
	"github.com/wrzfeijianshen/fyne2/driver/mobile"
	"github.com/fyne-io/mobile/app"
)

func showVirtualKeyboard(keyboard mobile.KeyboardType) {
	if driver, ok := fyne.CurrentApp().Driver().(*mobileDriver); ok {
		driver.app.ShowVirtualKeyboard(app.KeyboardType(keyboard))
	}
}

func hideVirtualKeyboard() {
	if driver, ok := fyne.CurrentApp().Driver().(*mobileDriver); ok {
		driver.app.HideVirtualKeyboard()
	}
}
