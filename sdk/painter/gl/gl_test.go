// +build !ci

package gl

import (
	"runtime"
	"testing"

	fyne "github.com/wrzfeijianshen/fyne2"
	"github.com/wrzfeijianshen/fyne2/canvas"
	"github.com/wrzfeijianshen/fyne2/theme"
)

func init() {
	runtime.LockOSThread()
}

func TestDrawImage_Ratio(t *testing.T) {
	//	d := NewGLDriver()
	//	win := d.CreateWindow("Test")
	//	c := win.Canvas().(*glCanvas)

	img := canvas.NewImageFromResource(theme.FyneLogo())
	img.Resize(fyne.NewSize(10, 10))
	//	c.newGlImageTexture(img)
	//	assert.Equal(t, float32(1.0), c.aspects[img])
}

func TestDrawImage_Ratio2(t *testing.T) {
	//	d := NewGLDriver()
	//	win := d.CreateWindow("Test")
	//	c := win.Canvas().(*glCanvas)

	// make sure we haven't used the visual ratio
	img := canvas.NewImageFromResource(theme.FyneLogo())
	img.Resize(fyne.NewSize(20, 10))
	//	c.newGlImageTexture(img)
	//	assert.Equal(t, float32(1.0), c.aspects[img])
}
