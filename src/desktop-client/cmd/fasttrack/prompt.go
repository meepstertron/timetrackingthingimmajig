// File is 100% AI generated!

package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// namePrompt is a tiny bubbletea program that asks the user to name a project.
type namePrompt struct {
	name string
	done bool // set once the user confirms a non-empty name
}

// promptProjectName runs the naming prompt and returns the entered name.
// The second return value is false if the prompt was cancelled.
func promptProjectName() (string, bool) {
	model, err := tea.NewProgram(namePrompt{}).Run()
	if err != nil {
		return "", false
	}
	result, ok := model.(namePrompt)
	if !ok || !result.done {
		return "", false
	}
	return result.name, true
}

func (m namePrompt) Init() tea.Cmd {
	return nil
}

func (m namePrompt) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if strings.TrimSpace(m.name) != "" {
				m.done = true
				return m, tea.Quit
			}
		case "esc", "ctrl+c":
			return m, tea.Quit
		case "backspace":
			if runes := []rune(m.name); len(runes) > 0 {
				m.name = string(runes[:len(runes)-1])
			}
		case " ":
			m.name += " "
		default:
			if len(msg.Runes) > 0 {
				m.name += string(msg.Runes)
			}
		}
	}
	return m, nil
}

func (m namePrompt) View() string {
	if m.done {
		return ""
	}
	return fmt.Sprintf("Name your project: %s█\n(enter to confirm, esc to cancel)\n", m.name)
}
