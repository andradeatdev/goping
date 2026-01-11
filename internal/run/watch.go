package run

import (
	"fmt"
	"os"

	"goping/internal/config"
	"goping/internal/ping"
	"time"
)

func RunWatch(cfg config.Config) {
	printHeader(cfg)

	logFile, err := ping.GetLogFile(cfg.LogPath, cfg.LogFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer logFile.Close()

	// w := bufio.NewWriter(logFile)
	// defer w.Flush()

	i := 1
	for shouldContinue(i, cfg.NumPings) {
		for _, url := range cfg.URLs {
			printPing(i, cfg.NumPings, url, logFile)
		}

		i++
		if cfg.NumPings != 0 && i > cfg.NumPings {
			break
		}

		sleepInterval(cfg.Interval)
	}
}

func shouldContinue(iteration, max int) bool {
	if max == 0 {
		return true
	}

	return iteration <= max
}

func printHeader(cfg config.Config) {
	pings := "∞"
	if cfg.NumPings > 0 {
		pings = fmt.Sprintf("%d", cfg.NumPings)
	}

	fmt.Printf(
		"watching %d urls | %s pings | interval %ds\n",
		len(cfg.URLs), pings, int(cfg.Interval.Seconds()),
	)
	fmt.Println("--------------------------------------------------")
}

func printPing(iteration, max int, url string, w *os.File) {
	status := ping.Ping(url)

	total := "∞"
	if max > 0 {
		total = fmt.Sprintf("%d", max)
	}

	logLine := fmt.Sprintf(
		"%s | [%d/%s] → %-30s %s\n",
		ping.Now(), iteration, total, url, status,
	)

	fmt.Fprint(w, logLine)
	fmt.Printf("%s", logLine)
}

func sleepInterval(d time.Duration) {
	fmt.Printf("next ping in %ds\n\n", int(d.Seconds()))
	time.Sleep(d)
}
