package cointop

import (
	"os"

	"github.com/miguelmota/gocui"
)

// Quit quites the program
func (ct *Cointop) Quit() error {
	return gocui.ErrQuit
}

// QuitView exists the current view
func (ct *Cointop) QuitView() error {
	ct.debuglog("quitView()")
	if ct.State.portfolioVisible {
		ct.State.portfolioVisible = false
		return ct.UpdateTable()
	}
	if ct.State.filterByFavorites {
		ct.State.filterByFavorites = false
		return ct.UpdateTable()
	}
	if ct.ActiveViewName() == ct.Views.Table.Name() {
		return ct.Quit()
	}

	return nil
}

// Exit safely exits the program
func (ct *Cointop) Exit() {
	ct.debuglog("exit()")
	if ct.g != nil {
		ct.g.Close()
	} else {
		os.Exit(0)
	}
}
