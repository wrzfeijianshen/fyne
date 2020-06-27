package playground

import (
	"github.com/wrzfeijianshen/fyne2/sdk/painter/software"
	"github.com/wrzfeijianshen/fyne2/test"
)

// NewSoftwareCanvas creates a new canvas in memory that can render without hardware support
func NewSoftwareCanvas() test.WindowlessCanvas {
	return test.NewCanvasWithPainter(software.NewPainter())
}
