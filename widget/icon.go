package widget

import (
	"image/color"

	"github.com/wrzfeijianshen/fyne2"
	"github.com/wrzfeijianshen/fyne2/canvas"
	"github.com/wrzfeijianshen/fyne2/sdk/widget"
	"github.com/wrzfeijianshen/fyne2/theme"
)

type iconRenderer struct {
	widget.BaseRenderer
	image *Icon
}

func (i *iconRenderer) MinSize() fyne.Size {
	size := theme.IconInlineSize()
	return fyne.NewSize(size, size)
}

func (i *iconRenderer) Layout(size fyne.Size) {
	if len(i.Objects()) == 0 {
		return
	}

	i.Objects()[0].Resize(size)
}

func (i *iconRenderer) BackgroundColor() color.Color {
	return color.Transparent
}

func (i *iconRenderer) Refresh() {
	i.image.propertyLock.RLock()
	i.updateObjects()
	i.image.propertyLock.RUnlock()
	i.Layout(i.image.Size())
	canvas.Refresh(i.image.super())
}

func (i *iconRenderer) updateObjects() {
	var objects []fyne.CanvasObject
	if i.image.Resource != nil {
		// TODO this is recreated every time - perhaps cache as long as i.image.Resource doesn't change
		raster := canvas.NewImageFromResource(i.image.Resource)
		raster.FillMode = canvas.ImageFillContain
		objects = append(objects, raster)
	}
	i.SetObjects(objects)
}

// Icon widget is a basic image component that load's its resource to match the theme.
type Icon struct {
	BaseWidget

	Resource fyne.Resource // The resource for this icon
}

// SetResource updates the resource rendered in this icon widget
func (i *Icon) SetResource(res fyne.Resource) {
	i.Resource = res
	i.Refresh()
}

// MinSize returns the size that this widget should not shrink below
func (i *Icon) MinSize() fyne.Size {
	i.ExtendBaseWidget(i)
	return i.BaseWidget.MinSize()
}

// CreateRenderer is a private method to Fyne which links this widget to its renderer
func (i *Icon) CreateRenderer() fyne.WidgetRenderer {
	i.ExtendBaseWidget(i)
	i.propertyLock.RLock()
	defer i.propertyLock.RUnlock()
	r := &iconRenderer{image: i}
	r.updateObjects()
	return r
}

// NewIcon returns a new icon widget that displays a themed icon resource
func NewIcon(res fyne.Resource) *Icon {
	icon := &Icon{}
	icon.ExtendBaseWidget(icon)
	icon.SetResource(res) // force the image conversion

	return icon
}
