package main

import (
	"github.com/unix-world/smart-fyne"
	"github.com/unix-world/smart-fyne/app"
	"github.com/unix-world/smart-fyne/cmd/fyne_settings/settings"
	"github.com/unix-world/smart-fyne/container"
)

func main() {
	s := settings.NewSettings()

	a := app.New()
	w := a.NewWindow("Fyne Settings")

	appearance := s.LoadAppearanceScreen(w)
	tabs := container.NewAppTabs(
		&container.TabItem{Text: "Appearance", Icon: s.AppearanceIcon(), Content: appearance})
	tabs.SetTabLocation(container.TabLocationLeading)
	w.SetContent(tabs)

	w.Resize(fyne.NewSize(520, 520))
	w.ShowAndRun()
}
