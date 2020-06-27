// +build ci nacl

package app

import (
	fyne "github.com/wrzfeijianshen/fyne2"
	"github.com/wrzfeijianshen/fyne2/sdk/painter/software"
	"github.com/wrzfeijianshen/fyne2/test"
)

// NewWithID returns a new app instance using the test (headless) driver.
// The ID string should be globally unique to this app.
func NewWithID(id string) fyne.App {
	return newAppWithDriver(test.NewDriverWithPainter(software.NewPainter()), id)
}
