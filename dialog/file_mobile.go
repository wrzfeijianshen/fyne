// +build ios android

package dialog

import (
	fyne "github.com/wrzfeijianshen/fyne2"
	"github.com/wrzfeijianshen/fyne2/internal/driver/gomobile"
)

func (f *fileDialog) loadPlaces() []fyne.CanvasObject {
	return nil
}

func isHidden(file, _ string) bool {
	return false
}

func fileOpenOSOverride(f *FileDialog) bool {
	gomobile.ShowFileOpenPicker(f.callback.(func(fyne.URIReadCloser, error)), f.filter)
	return true
}

func fileSaveOSOverride(f *FileDialog) bool {
	ShowInformation("File Save", "File save not available on mobile", f.parent)

	callback := f.callback.(func(fyne.URIWriteCloser, error))
	if callback != nil {
		callback(nil, nil)
	}

	return true
}
