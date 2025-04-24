package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type menuModel struct {
	choices  []option
	cursor   int
	selected map[int]struct{}
}

type option struct {
	label   string
	execute func()
}

const padding = "    "

var styleSelected = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FAFAFA")).
	Background(lipgloss.Color("#7D56F4")).
	Width(15).
	PaddingLeft(2)

func StartMenu() {
	p := tea.NewProgram(initialMenuModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Oh no! An error while starting the menu: %v", err)
		os.Exit(1)
	}
}

func initialMenuModel() menuModel {
	return menuModel{
		choices: []option{
			option{"work", func() { startProgress(25) }},
			option{"short break", func() { startProgress(5) }},
			option{"long break", func() { startProgress(25) }},
		},
		selected: make(map[int]struct{}),
	}
}

func (m menuModel) Init() tea.Cmd {
	return tea.ClearScreen
}

func (m menuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
		case "enter":
			m.choices[m.cursor].execute()
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m menuModel) View() string {
	s := "\n" + padding + "Which phase to start?\n\n"
	for i, choice := range m.choices {
		s += padding
		if m.cursor == i {
			s += styleSelected.Render(choice.label) + "\n"
		} else {
			s += choice.label + "\n"
		}
	}
	s += "\n" + padding + "Press q to quit.\n\n"
	return s
}

func startProgress(minutes int) { // TODO
	m := progressModel{
		progress.New(progress.WithDefaultGradient()),
	}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Oh no! An error while starting the timer: %v", err)
		os.Exit(1)
	}
}

type tickMsg time.Time

type progressModel struct {
	progress progress.Model
}

func (p progressModel) Init() tea.Cmd {
	return tea.Batch(tickCmd(), tea.ClearScreen)
}

func (p progressModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" || msg.String() == "q" {
			return p, tea.Quit
		}
	case tickMsg:
		if p.progress.Percent() == 1.0 {
			return p, tea.Quit
		}

		cmd := p.progress.IncrPercent(0.25)
		return p, tea.Batch(tickCmd(), cmd)
	case progress.FrameMsg:
		progressModel, cmd := p.progress.Update(msg)
		p.progress = progressModel.(progress.Model)
		return p, cmd
	}
	return p, nil
}

func (p progressModel) View() string {
	return "\n" + p.progress.View() + "\n\n" + "Press q to exit."
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second * 1, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
