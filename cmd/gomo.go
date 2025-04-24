package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

type option struct {
	label   string
	execute func() // TODO startTimer(n minutes)
}

const padding = "  "

var styleSelected = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FAFAFA")).
	Background(lipgloss.Color("#7D56F4")).
	Width(15).
	PaddingLeft(2)

func StartMenu() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Oh no! An error while starting the menu: %v", err)
		os.Exit(1)
	}
}

func initialModel() model {
	return model{
		choices:  []string{"work", "short break", "long break"},
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	s := padding + "Which phase to start?\n\n"
	for i, choice := range m.choices {
		s += padding
		if m.cursor == i {
			s += styleSelected.Render(choice) + "\n"
		} else {
			s += choice + "\n"
		}
	}
	s += "\n" + padding + "Press q to quit.\n"
	return s
}
