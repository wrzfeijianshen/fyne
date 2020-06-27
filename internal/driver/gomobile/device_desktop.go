// +build !ios,!android

package gomobile

import fyne "github.com/wrzfeijianshen/fyne2"

func (*device) SystemScaleForWindow(_ fyne.Window) float32 {
	return 2 // this is simply due to the high number of pixels on a mobile device - just an approximation
}
