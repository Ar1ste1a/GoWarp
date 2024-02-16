package warp

import "github.com/rivo/tview"

func startCLI() {
	defer func() {
		if r := recover(); r != nil {
			// Handle the panic
		}
	}()

	// Create a new application
	tv := tview.Application{}
	go func() {
		err := tv.Run()
		if err != nil {
			panic(err)
		}
	}()
}
