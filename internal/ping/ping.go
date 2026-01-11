package ping

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func Ping(url string) string {
	resp, err := http.Head(url)
	if err != nil {
		return fmt.Sprintf("[ERR] %v", err)
	}
	defer resp.Body.Close()

	return fmt.Sprintf("[%s]", resp.Status)
}

func GetLogFile(dir, logFile string) (*os.File, error) {
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return nil, err
	}

	logFile = fmt.Sprintf("%s/%s", dir, logFile)

	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func Now() string {
	return time.Now().Format("15:04:05")
}
