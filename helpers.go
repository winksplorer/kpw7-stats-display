package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
