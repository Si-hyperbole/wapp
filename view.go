// fix out put layout
// add refresh
// fix error handleing miss spelt city

package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"charm.land/bubbles/v2/cursor"
	tea "charm.land/bubbletea/v2"
	owm "github.com/briandowns/openweathermap"
	"github.com/joho/godotenv"
)

// update handles msg and typing
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	//var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			m.quitting = true
			return m, tea.Quit

		// Change cursor mode
		case "ctrl+r":
			m.cursorMode++
			if m.cursorMode > cursor.CursorHide {
				m.cursorMode = cursor.CursorBlink
			}
			cmds := make([]tea.Cmd, len(m.inputs))
			for i := range m.inputs {
				s := m.inputs[i].Styles()
				s.Cursor.Blink = m.cursorMode == cursor.CursorBlink
				m.inputs[i].SetStyles(s)
			}
			return m, tea.Batch(cmds...)

		// Set focus to next input
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			// Did the user press enter while the submit button was focused?
			if s == "enter" && m.focusIndex == len(m.inputs) { //index the length of the input of the cussor with out it reads the data that cursor is on

				//return m, tea.Quit // when reachest ent of list its at submit

				return m, WeatherResults(m.inputs[0].Value())

			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > len(m.inputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs)
			}

			cmds := make([]tea.Cmd, len(m.inputs))
			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == m.focusIndex {
					// Set focused state
					cmds[i] = m.inputs[i].Focus()
					continue
				}
				// Remove focused state
				m.inputs[i].Blur()
			}

			return m, tea.Batch(cmds...)
		}
	}
	// Handle character input and blinking
	cmd := m.updateInputs(msg)
	return m, cmd

} // end update

func (m *model) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}
	return tea.Batch(cmds...)
} // end updateInput

// view : return a string based on the state of the model
func (m model) View() tea.View {
	var b strings.Builder
	var c *tea.Cursor

	// Render all text input
	for i, in := range m.inputs {
		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
		if m.cursorMode != cursor.CursorHide && in.Focused() {
			c = in.Cursor()
			if c != nil {
				c.Y += i
			}
		}
	}

	// Render the Submit button

	button := &blurredButton
	if m.focusIndex == len(m.inputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	// Render the help menu
	b.WriteString(helpStyle.Render("cursor mode is "))
	b.WriteString(cursorModeHelpStyle.Render(m.cursorMode.String())) // brings up what ever is up blink, solid ...
	b.WriteString(helpStyle.Render(" (ctrl+r to change style)"))

	if m.quitting {
		b.WriteRune('\n')
	}

	// Return the view directly without manually overriding the cursor
	v := tea.NewView(b.String())
	v.Cursor = c
	return v
} //end View

// API
func WeatherApi(q string) tea.Cmd {
	return func() tea.Msg {
		apiKey := os.Getenv("OWM_API_KEY")
		w, err := owm.NewCurrent("C", "EN", apiKey) // C - Celcious with english output
		if apiKey == "" {
			fmt.Println("Error: OWM_API_KEY enviorment variable not set.")
			fmt.Println("Set it using: export OWM_API_KEY=your_api_key_here")
			//return tea.Cmd
		}
		if err != nil {
			log.Fatalln(err)
		}

		w.CurrentByName(q)
		fmt.Printf("Temperature: %.1f°C\n", w.Main.Temp)
		fmt.Printf("Temperature Feels like: %.1f°C\n", w.Main.FeelsLike)
		fmt.Printf("Humidity: %d%%\n", w.Main.Humidity)
		fmt.Printf("Conditions: %s\n", w.Weather[0].Description)
		fmt.Printf("Wind Speed: %.1f m/s\n", w.Wind.Speed)

		//if os.Args[0] == "" {
		//	fmt.Println("error os.Args[1] ")
		//}
		return nil

	}

} // WeatherApi

// (os.Args[]) os is the name of a standard Go package that gives your program access to the operating system.
func WeatherResults(c string) tea.Cmd { // needs to creat a string to call, to reconize the input
	err := godotenv.Load() //loads .env automatically
	if err != nil {
		log.Fatalln("Error loading .env file")
	}
	//if len(os.Args) < 2 {
	//	log.Fatal("Usage: wapp.exe <city name>")
	//}

	return WeatherApi(c) //calls the api and the string c

}

// msg
type ErrorMsg error
