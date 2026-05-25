// Copyright © 2026 Alex Fu <alexfu@fastmail.com>

package ui

import (
	"os"

	"charm.land/bubbles/v2/spinner"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type loadingUI struct {
	spinner spinner.Model
	loading bool
	work    func() tea.Msg
}

type loadResult struct {
	loading bool
}

func loadCmd(work func()) tea.Cmd {
	return func() tea.Msg {
		work()
		return loadResult{loading: false}
	}
}

func (m loadingUI) Init() tea.Cmd {
	return tea.Batch(m.work, m.spinner.Tick)
}

func (m loadingUI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case loadResult:
		if !msg.loading {
			return m, tea.Quit
		}
	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m loadingUI) View() tea.View {
	return tea.NewView(lipgloss.JoinHorizontal(lipgloss.Left, m.spinner.View(), "Loading..."))
}

func NewLoadingUI(work func()) error {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Blue)
	_, err := tea.NewProgram(loadingUI{spinner: s, work: loadCmd(work), loading: true}, tea.WithOutput(os.Stderr)).Run()
	return err
}
