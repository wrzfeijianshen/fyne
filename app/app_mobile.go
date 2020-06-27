// +build !ci

// +build android ios mobile

package app

import (
	fyne "github.com/wrzfeijianshen/fyne2"
	"github.com/wrzfeijianshen/fyne2/internal/driver/gomobile"
)

// NewWithID returns a new app instance using the appropriate runtime driver.
// The ID string should be globally unique to this app.
func NewWithID(id string) fyne.App {
	return newAppWithDriver(gomobile.NewGoMobileDriver(), id)
}
