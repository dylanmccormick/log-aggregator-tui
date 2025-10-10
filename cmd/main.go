package main

import (
	"fmt"
	"os"

	lg "github.com/dylanmccormick/log-aggregator-tui/internal/log"
	ui "github.com/dylanmccormick/log-aggregator-tui/internal/ui"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct{}

func main() {
	// This is going to have 2 parts.

	// 1: Start the reader/ aggregator to run as a goroutine... I think
	// 2: Load the TUI

	// m := model{}
	rm := ui.RootModel{
		LogList: &ui.LogListComponent{},
	}
	p := tea.NewProgram(rm)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there has been an error: %v", err)
		os.Exit(1)
	}

	logMessages, err := lg.ReadLogFile("../testdata/sample.log")
	if err != nil {
		panic(err)
	}
	for _, msg := range logMessages {
		fmt.Printf("%#v\n", msg.Level)
	}
}

func (m model) View() string {
	return "Hello world"
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m model) Init() tea.Cmd {
	return nil
}
