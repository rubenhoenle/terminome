package main

import (
	"fmt"
	"log"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type metronomeState uint

var (
	modelStyle = lipgloss.NewStyle().
			Width(15).
			Height(5).
			Align(lipgloss.Center, lipgloss.Center).
			BorderStyle(lipgloss.HiddenBorder()).
			Background(lipgloss.Color("#00ff00")).
			Foreground(lipgloss.Color("#ffffff"))
	focusedModelStyle = lipgloss.NewStyle().
				Width(15).
				Height(5).
				Align(lipgloss.Center, lipgloss.Center).
				BorderStyle(lipgloss.HiddenBorder()).
				Background(lipgloss.Color("#ff0000")).
				Foreground(lipgloss.Color("#ffffff"))
	helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
)

type mainModel struct {
	state metronomeState
}

func newModel(timeout time.Duration) mainModel {
	m := mainModel{state: 0}
	return m
}

func (m mainModel) Init() tea.Cmd {
	return tea.Batch()
}

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "tab":
			if m.state >= 3 {
				m.state = 0
			} else {
				m.state += 1
			}
		case "n":
		}
	}
	return m, tea.Batch(cmds...)
}

func (m mainModel) View() string {
	var s string
	contents := []string{
		modelStyle.Render("1"),
		modelStyle.Render("2"),
		modelStyle.Render("3"),
		modelStyle.Render("4"),
	}
	contents[m.state] = focusedModelStyle.Render(fmt.Sprintf("%d", m.state+1))
	s += lipgloss.JoinHorizontal(lipgloss.Top, contents...)
	s += helpStyle.Render("\ntab: focus next • n: new • q: exit\n")
	return s
}

func main() {
	p := tea.NewProgram(newModel(0))

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
