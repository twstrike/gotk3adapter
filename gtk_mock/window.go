package gtk_mock

import (
	"github.com/coyim/gotk3adapter/gdki"
	"github.com/coyim/gotk3adapter/gtki"
	"github.com/gotk3/gotk3/gdk"
)

type MockWindow struct {
	MockBin
}

func (*MockWindow) AddAccelGroup(v2 gtki.AccelGroup) {
}

func (*MockWindow) GetTitle() string {
	return ""
}

func (*MockWindow) HasToplevelFocus() bool {
	return false
}

func (*MockWindow) Fullscreen() {
}

func (*MockWindow) Unfullscreen() {
}

func (*MockWindow) IsActive() bool {
	return false
}

func (*MockWindow) Resize(v1, v2 int) {
}

func (*MockWindow) SetApplication(v2 gtki.Application) {
}

func (*MockWindow) SetIcon(v2 gdki.Pixbuf) {
}

func (*MockWindow) SetTitle(v1 string) {
}

func (*MockWindow) SetTitlebar(v2 gtki.Widget) {
}

func (*MockWindow) SetTransientFor(v2 gtki.Window) {
}

func (*MockWindow) SetUrgencyHint(v2 bool) {
}

func (*MockWindow) Present() {
}

func (*MockWindow) AddMnemonic(v1 uint, v2 gtki.Widget) {
}

func (*MockWindow) RemoveMnemonic(v1 uint, v2 gtki.Widget) {
}

func (*MockWindow) ActivateMnemonic(v1 uint, v2 gdki.ModifierType) bool {
	return true
}

func (*MockWindow) GetMnemonicModifier() gdk.ModifierType {
	return gdk.ModifierType(gdki.GDK_SHIFT_MASK)
}

func (*MockWindow) SetMnemonicModifier(v1 gdki.ModifierType) {
}

func (*MockWindow) GetSize() (int, int) {
	return 0, 0
}
