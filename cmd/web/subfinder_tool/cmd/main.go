package livesub

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/hacker301et/sec1/cmd/web/subfinder_tool/view"
)

func Start() {
	m := view.NewView()
	p := tea.NewProgram(m)

	if _, err := p.Run(); err != nil {
		fmt.Println("could not start program:", err)
		os.Exit(1)
	}

}
