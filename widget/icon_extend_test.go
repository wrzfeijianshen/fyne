package widget

import (
	"testing"

	"github.com/wrzfeijianshen/fyne2/canvas"
	"github.com/wrzfeijianshen/fyne2/internal/cache"
	"github.com/wrzfeijianshen/fyne2/theme"
	"github.com/stretchr/testify/assert"
)

type extendedIcon struct {
	Icon
}

func newExtendedIcon() *extendedIcon {
	icon := &extendedIcon{}
	icon.ExtendBaseWidget(icon)
	return icon
}

func TestIcon_Extended_SetResource(t *testing.T) {
	icon := newExtendedIcon()
	icon.SetResource(theme.FyneLogo())

	objs := cache.Renderer(icon).Objects()
	assert.Equal(t, 1, len(objs))
	assert.Equal(t, theme.FyneLogo(), objs[0].(*canvas.Image).Resource)
}
