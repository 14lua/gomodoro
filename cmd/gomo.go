package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	padding  = 2
	maxWidth = 60
)

var currentPhase int = 0

func Start(phase int) {
	switch phase {
	case 0: // short break
		currentPhase = 0
		StartTimer(5)
	case 1: // long break
		currentPhase = 1
		StartTimer(30)
	case 2: // work
		currentPhase = 2
		StartTimer(25)
	}
}

func StartTimer(minutes int) {

	var startColour string
	var endColour string
	fmt.Println(currentPhase)
	switch currentPhase {
	case 0:
		startColour = "#AA5555"
		endColour = "#FFAAAA"
	case 1:
		startColour = "#55AA55"
		endColour = "#AAFFAA"
	case 2:
		startColour = "#5555FF"
		endColour = "#AAAAFF"
	}
	prog := progress.New(progress.WithScaledGradient(startColour, endColour))

	if _, err := tea.NewProgram(model{percent: 0.0, currentSeconds: 0, duration: 1, progress: prog}).Run(); err != nil {
		fmt.Println("Oh no!", err)
		os.Exit(1)
	}
}

type tickMsg time.Time

type model struct {
	percent        float64
	currentSeconds int
	duration       int
	progress       progress.Model
}

func (m model) Init() tea.Cmd {
	return tickCmd()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m, tea.Quit

	case tea.WindowSizeMsg:
		m.progress.Width = msg.Width - padding*2 - 4
		if m.progress.Width > maxWidth {
			m.progress.Width = maxWidth
		}
		return m, nil

	case tickMsg:
		m.percent += (1 / float64(m.duration))
		m.currentSeconds++
		if m.percent > 1.0 {
			m.percent = 1.0
			return m, tea.Quit
		}
		return m, tickCmd()

	default:
		return m, nil
	}
}

func (m model) View() string {
	pad := strings.Repeat(" ", padding)
	return fmt.Sprintf("\n%s%s\n\n%sElapsed Time: %d seconds\n",
		pad, m.progress.ViewAs(m.percent), pad, m.duration-m.currentSeconds)
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
