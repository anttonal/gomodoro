package main

import (
	"fmt"
	"time"
)

func countdown(seconds int) {
	for ; seconds > 0; seconds-- {
		minutes := seconds / 60
		fmt.Printf("\r%02d:%02d", minutes, seconds)
		time.Sleep(time.Second)
	}
}

// TODO: define runPomodoro(cfg Config) — work -> short break -> long break cycle
func runPomodoro(cfg Config) {
	for {
		for i := cfg.Cycles; i > 0; i-- {
			countdown(cfg.WorkDuration)
			countdown(cfg.ShortBreak)
		}
		countdown(cfg.LongBreak)
	}
}
