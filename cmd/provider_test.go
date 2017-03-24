package cmd

import (
	"testing"

	"github.com/KyleBanks/goggles"
	"github.com/KyleBanks/goggles/conf"
	"github.com/KyleBanks/goggles/pkg/sys"
)

type mockProvider struct {
	ListFn    func() ([]*goggles.Package, error)
	DetailsFn func(string) (*goggles.Package, error)

	OpenFileExplorerFn func(string)
	OpenTerminalFn     func(string)
	OpenBrowserFn      func(string)

	preferencesFn       func() *conf.Config
	updatePreferencesFn func(*conf.Config)
}

func (m *mockProvider) List() ([]*goggles.Package, error)          { return m.ListFn() }
func (m *mockProvider) Details(n string) (*goggles.Package, error) { return m.DetailsFn(n) }
func (m *mockProvider) OpenFileExplorer(n string)                  { m.OpenFileExplorerFn(n) }
func (m *mockProvider) OpenTerminal(n string)                      { m.OpenTerminalFn(n) }
func (m *mockProvider) OpenBrowser(n string)                       { m.OpenBrowserFn(n) }
func (m *mockProvider) Preferences() *conf.Config                  { return m.preferencesFn() }
func (m *mockProvider) UpdatePreferences(c *conf.Config)           { m.updatePreferencesFn(c) }

func TestProvider_Preferences(t *testing.T) {
	var p provider

	c := p.Preferences()
	if c == nil {
		t.Fatal("Unexpected nil Config")
	}
}

func TestProvider_UpdatePreferences(t *testing.T) {
	var p provider

	c := p.Preferences()
	sys.SetGopath("/not/a/real/gopath")
	p.UpdatePreferences(c)

	if sys.RawGopath() != c.Gopath {
		t.Fatalf("Unexpected RawGopath, expected=%v, got=%v", c.Gopath, sys.RawGopath())
	}
}
