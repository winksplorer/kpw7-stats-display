package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// used for humanReadable functions
const (
	kb = 1024
	mb = kb * 1024
	gb = mb * 1024
	tb = gb * 1024
)

// get unixtime since boot
func getBootTime() (int64, error) {
	file, err := os.Open("/proc/stat")
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "btime") {
			fields := strings.Fields(line)
			if len(fields) == 2 {
				return strconv.ParseInt(fields[1], 10, 64)
			}
			break
		}
	}
	return 0, fmt.Errorf("btime not found")
}

// round to decimal place
func roundTo(x float64, places int) float64 {
	pow := math.Pow(10, float64(places))
	return math.Round(x*pow) / pow
}

// human readable byte sizes
func humanReadable(bytes uint64) string {
	switch {
	case bytes >= tb:
		return fmt.Sprintf("%.1ft", roundTo(float64(bytes)/float64(tb), 1))
	case bytes >= gb:
		return fmt.Sprintf("%.1fg", roundTo(float64(bytes)/float64(gb), 1))
	case bytes >= mb:
		return fmt.Sprintf("%.1fm", roundTo(float64(bytes)/float64(mb), 1))
	case bytes >= kb:
		return fmt.Sprintf("%.1fk", roundTo(float64(bytes)/float64(kb), 1))
	default:
		return fmt.Sprintf("%d", bytes)
	}
}
