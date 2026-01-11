package run

import (
	"fmt"
	"goping/internal/ping"
)

func RunOnce(urls []string) {
	for _, url := range urls {
		fmt.Printf("→ %-30s %s\n", url, ping.Ping(url))
	}
}
