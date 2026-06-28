package main

import (
	"os/exec"
)

func notify(msg string) {
	if _, err := exec.LookPath("notify-send"); err == nil {
		exec.Command("notify-send", "-a", "timer", msg).Run()
	}
}
