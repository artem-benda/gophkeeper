package application

import tea "github.com/charmbracelet/bubbletea"

type App struct {
	App *tea.Program
}

func NewApp() *App {
	app := tea.NewProgram(initialModel())
	return &App{
		App: app,
	}
}
