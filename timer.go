package main

import (
	"fmt"
	"time"
)

func countdown(seconds int) {
	for ; seconds > 0; seconds-- {
		minutes := seconds / 60
		fmt.Printf("\r%02d:%02d", minutes, seconds%60)
		time.Sleep(time.Second)
	}
	fmt.Printf("\r%02d:%02d", seconds/60, seconds%60)
}

func runPomodoro(cfg Config) {
	for {
		for i := cfg.Cycles; i > 0; i-- {
			countdown(cfg.WorkDuration)
			notify("work session done, take a break")
			countdown(cfg.ShortBreak)
			notify("break over, back to work")
		}
		countdown(cfg.LongBreak)
		notify("long break over, starting new cycle")
	}
}
