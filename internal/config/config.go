package config

import (
	"flag"
	"fmt"
	"time"
)

type Config struct {
	NumPings int
	Interval time.Duration
	URLs     []string
	LogPath  string
	LogFile  string
}

func ParseWatchFlags(args []string) (Config, error) {
	watch := flag.NewFlagSet("watch", flag.ExitOnError)

	count := watch.Int("count", 0, "number of pings to send (0 = infinite)")
	interval := watch.Int("interval", 5, "interval between pings in seconds")

	watch.Parse(args)

	urls := watch.Args()
	if len(urls) == 0 {
		return Config{}, fmt.Errorf("no urls provided")
	}

	logFile := fmt.Sprintf("watch-%s.log", time.Now().Format("2006-01-02-15-04-05"))

	return Config{
		NumPings: *count,
		Interval: time.Duration(*interval) * time.Second,
		URLs:     urls,
		LogPath:  "logs",
		LogFile:  logFile,
	}, nil
}
