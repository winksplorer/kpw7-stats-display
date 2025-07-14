package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	probing "github.com/prometheus-community/pro-bing"
)

func hostname(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	hostname, err := os.Hostname()
	if err != nil {
		log.Println("couldn't get hostname:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, hostname)
}

func bootTime(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	bootTime, err := getBootTime()
	if err != nil {
		log.Println("couldn't get system boot time:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, bootTime)
}

func ping(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.Header.Get("X-Custom-Token") == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	pinger, err := probing.NewPinger(r.Header.Get("X-Custom-Token"))
	if err != nil {
		log.Println("couldn't set up pinger:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	pinger.Count = 3
	pinger.Timeout = time.Second * 6

	err = pinger.Run() // Blocks until finished.
	if err != nil {
		log.Println("couldn't ping:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	stats := pinger.Statistics() // get send/receive/duplicate/rtt stats

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, stats.PacketLoss)
}
