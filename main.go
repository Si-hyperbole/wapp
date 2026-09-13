
// -------------    References  -------------------------
// https://charm.land/blog/v2/
// https://www.youtube.com/watch?v=3QnwrBqLqnI&t=90s
// Info:   https://www.youtube.com/watch?v=3QnwrBqLqnI&t=90s
// https://dev.to/andyhaskell/intro-to-bubble-tea-in-go-21lg
//============ charmland v2 =============
//https://github.com/charmbracelet/bubbletea/discussions/1374
// https://www.youtube.com/watch?v=DEeFnVj3cv8&t=338s

// -------------- calling api in buble CMD
// https://www.youtube.com/watch?v=DEeFnVj3cv8&t=879s
// ------------------------------------------------------------
// ------------------------------------------------------------
package main

import (
	"fmt"
	"os"

	//"weatherApp/API"

	tea "charm.land/bubbletea/v2"
)

func main() {
	if _, err := tea.NewProgram(initialModel()).Run(); err != nil {
		fmt.Printf("could not start program: %s\n", err)
		os.Exit(1)

	}

	//if err != nil {
	//	fmt.Fprint(os.Stderr, err)
	//	return
	//}

}
