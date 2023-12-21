package test

import (
	"github.com/unix-world/smart-fyne"
	"github.com/unix-world/smart-fyne/theme"
)

type mockCloud struct {
	configured bool
}

func (c *mockCloud) Cleanup(_ fyne.App) {
	c.configured = false
}

func (c *mockCloud) ProviderDescription() string {
	return "Mock cloud implementation"
}

func (c *mockCloud) ProviderIcon() fyne.Resource {
	return theme.ComputerIcon()
}

func (c *mockCloud) ProviderName() string {
	return "mock"
}

func (c *mockCloud) Setup(_ fyne.App) error {
	c.configured = true
	return nil
}
