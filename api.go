package main

import (
	"fmt"
	"net/http"
	"time"

	probing "github.com/prometheus-community/pro-bing"
)

func ping(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	pinger, err := probing.NewPinger("192.168.1.174")
	if err != nil {
		panic(err)
	}
	pinger.Count = 3
	pinger.Timeout = time.Second * 10
	err = pinger.Run() // Blocks until finished.
	if err != nil {
		panic(err)
	}
	stats := pinger.Statistics() // get send/receive/duplicate/rtt stats

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, stats.PacketLoss)
}
