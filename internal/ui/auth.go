// Copyright © 2026 Alex Fu <alexfu@fastmail.com>

package ui

import (
	"errors"
	"os"

	"ynab/internal/auth"
	"ynab/internal/ynab"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type loginUIModel struct {
	textInput textinput.Model
	canceled  bool
}

func newLoginUIModel() loginUIModel {
	textInput := textinput.New()
	textInput.Placeholder = "Enter YNAB Token"
	textInput.Focus()
	textInput.EchoMode = textinput.EchoPassword
	return loginUIModel{
		textInput: textInput,
	}
}

func (m loginUIModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m loginUIModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "enter":
			return m, tea.Quit
		case "esc", "ctrl+c":
			m.canceled = true
			return m, tea.Quit
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)

	// Keep width at least placeholder-length, grow with the typed value
	m.textInput.SetWidth(max(
		len(m.textInput.Placeholder),
		len(m.textInput.Value())+1, // +1 reserves a column for the cursor
	))
	return m, cmd
}

func (m loginUIModel) View() tea.View {
	title := "Logging into YNAB"

	helpText := lipgloss.NewStyle().
		Italic(true).
		Foreground(lipgloss.Color("8")).
		Render("Don't have a token? Get one at https://app.ynab.com/settings/developer")

	view := tea.NewView(
		lipgloss.NewStyle().
			Padding(1).
			Render(
				lipgloss.JoinVertical(lipgloss.Left, title, "", m.textInput.View(), helpText, "", "esc to cancel"),
			),
	)

	view.AltScreen = true
	return view
}

func NewLoginUI() error {
	m, err := tea.NewProgram(newLoginUIModel(), tea.WithOutput(os.Stderr)).Run()
	if err != nil {
		return err
	}
	model := m.(loginUIModel)
	if model.canceled {
		return errors.New("login cancelled")
	}
	token := model.textInput.Value()
	if !ynab.TokenValid(token) {
		return errors.New("token not valid")
	}
	return auth.SaveToken(token)
}
