package main

import (
	"fmt"

	"charm.land/bubbles/v2/cursor"
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var (
	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle         = focusedStyle
	noStyle             = lipgloss.NewStyle()
	helpStyle           = blurredStyle
	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	focusedButton = focusedStyle.Render("[ Submit ]")                    // where the cuurser is on it
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit")) // shows when curser not on it [%s] shows submit

)

type model struct {
	focusIndex int
	inputs     []textinput.Model
	cursorMode cursor.Mode
	quitting   bool
}

func initialModel() model {
	m := model{
		inputs: make([]textinput.Model, 1), // SWITCH CITY [] text input from the area on switch i case 0 = weather (takes only the first letter)
	}

	var ti textinput.Model
	for i := range m.inputs {
		ti = textinput.New()
		ti.CharLimit = 32

		s := ti.Styles()
		s.Cursor.Color = lipgloss.Color("205")
		s.Focused.Prompt = focusedStyle
		s.Focused.Text = focusedStyle
		s.Blurred.Prompt = blurredStyle
		s.Focused.Text = focusedStyle
		ti.SetStyles(s)
		ti.SetWidth(20) // with out this it will only set the place holder to the first letter
		//ti.Placeholder = "City Location"
		//ti.Focus() //start here for key input

		// i = counter
		switch i {
		case 0:
			ti.Placeholder = "City Location"
			ti.Focus() //start here for key input
		//not being used
		case 1:
			ti.Placeholder = "City 2"
			//ti.CharLimit = 64
		case 2:
			ti.Placeholder = "City 3"
			//ti.EchoMode = textinput.EchoPassword
			//ti.EchoCharacter = '•'
		}
		m.inputs[i] = ti // array counter
	}

	return m
}

// kiscks off the event loop
func (m model) Init() tea.Cmd {
	return textinput.Blink // initiats cursor loop
}
