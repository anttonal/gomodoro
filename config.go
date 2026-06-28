package main

type Config struct {
	WorkDuration int
	ShortBreak   int
	LongBreak    int
	Cycles       int
}

func defaultConfig() Config {
	return Config{
		WorkDuration: 50 * 60,
		ShortBreak:   10 * 60,
		LongBreak:    30 * 60,
		Cycles:       2,
	}
}
