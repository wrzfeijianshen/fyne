// +build ci !linux,!darwin,!windows,!freebsd,!openbsd,!netbsd

package app

import (
	"errors"
	"net/url"

	fyne "github.com/wrzfeijianshen/fyne2"
	"github.com/wrzfeijianshen/fyne2/theme"
)

func defaultTheme() fyne.Theme {
	return theme.DarkTheme()
}

func rootConfigDir() string {
	return "/tmp/fyne-test/"
}

func (app *fyneApp) OpenURL(_ *url.URL) error {
	return errors.New("Unable to open url for unknown operating system")
}

func (app *fyneApp) SendNotification(_ *fyne.Notification) {
	fyne.LogError("Refusing to show notification for unknown operating system", nil)
}
