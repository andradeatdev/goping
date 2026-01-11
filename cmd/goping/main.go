package main

import (
	"fmt"
	"goping/internal/config"
	"goping/internal/run"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}

	switch os.Args[1] {
	case "watch":
		cfg, err := config.ParseWatchFlags(os.Args[2:])
		if err != nil {
			fmt.Println(err)
			return
		}
		run.RunWatch(cfg)

	default:
		run.RunOnce(os.Args[1:])
	}
}

func usage() {
	fmt.Println("usage: goping [watch] [options] urls...")
}
