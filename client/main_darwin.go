package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/tomrittervg/pond/client/system"
)

func main() {
	devFlag := flag.Bool("dev", false, "Is this a development environment?")
	stateFile := flag.String("state-file", "", "File in which to save persistent state")
	cliFlag := flag.Bool("cli", false, "If true, the CLI will be used, even if the GUI is available")
	flag.Parse()

	runtime.LockOSThread()

	dev := os.Getenv("POND") == "dev" || *devFlag
	runtime.GOMAXPROCS(4)

	if len(*stateFile) == 0 && dev {
		*stateFile = "state"
	}

	if len(*stateFile) == 0 {
		home := os.Getenv("HOME")
		if len(home) == 0 {
			fmt.Fprintf(os.Stderr, "$HOME not set. Please export $HOME to set the directory for the state file.\n")
			os.Exit(1)
		}
		*stateFile = filepath.Join(home, ".pond")
	}

	defer system.Shutdown()

	if !haveGUI || *cliFlag {
		client := NewCLIClient(*stateFile, rand.Reader, false /* testing */, true /* autoFetch */)
		client.disableV2Ratchet = true
		client.dev = dev
		client.Start()
	} else {
		ui := NewGTKUI()
		client := NewGUIClient(*stateFile, ui, rand.Reader, false /* testing */, true /* autoFetch */)
		client.disableV2Ratchet = true
		client.dev = dev
		client.Start()
		ui.Run()
	}
}
