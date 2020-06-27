package desktop

import "github.com/wrzfeijianshen/fyne2"

// Driver represents the extended capabilities of a desktop driver
type Driver interface {
	// Create a new borderless window that is centered on screen
	CreateSplashWindow() fyne.Window
}
