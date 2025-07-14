package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"

	probing "github.com/prometheus-community/pro-bing"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/v3/cpu"
)

// hostname@
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

// uptime@
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

// cpu@
func cpuUsage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	percentages, err := cpu.Percent(0, false)
	if err != nil {
		log.Println("couldn't get cpu info:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, roundTo(percentages[0], 2))
}

// nvidia@
func nvidiaUsage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	out, err := exec.Command("nvidia-smi", "--query-gpu=utilization.gpu", "--format=csv,noheader,nounits").Output()
	if err != nil {
		log.Println("couldn't get gpu info:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, string(out))
}

// mem@
func memUsage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	virtualMem, err := mem.VirtualMemory()
	if err != nil {
		log.Println("couldn't get memory info:", err)
	}

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "%s/%s", humanReadable(virtualMem.Used), humanReadable(virtualMem.Total))
}

// ping@ip@sec
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
