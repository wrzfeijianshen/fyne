// +build ios

package gomobile

import fyne "github.com/wrzfeijianshen/fyne2"

func (*device) SystemScaleForWindow(_ fyne.Window) float32 {
	if currentDPI >= 450 {
		return 3
	} else if currentDPI >= 340 {
		return 2.5
	}
	return 2
}
