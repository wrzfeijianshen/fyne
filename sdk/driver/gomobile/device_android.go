// +build android

package gomobile

import fyne "github.com/wrzfeijianshen/fyne2"

func (*device) SystemScaleForWindow(_ fyne.Window) float32 {
	if currentDPI >= 600 {
		return 4
	} else if currentDPI >= 405 {
		return 3
	} else if currentDPI >= 270 {
		return 2
	} else if currentDPI >= 180 {
		return 1.5
	}
	return 1
}
