package main

import (
	"log"
	"net/http"
	"os"
)

var frontendDir string = "/tmp/kpw7-stats-display/frontend"
var port string = ":42799"

func main() {
	if len(os.Args) > 1 && os.Args[1] == "dev" {
		frontendDir = "frontend"
	}

	http.Handle("/", http.FileServer(http.Dir(frontendDir)))
	http.HandleFunc("/hostname", hostname)
	http.HandleFunc("/boot-time", bootTime)
	http.HandleFunc("/cpu-usage", cpuUsage)
	http.HandleFunc("/nvidia-usage", nvidiaUsage)
	http.HandleFunc("/mem-usage", memUsage)
	http.HandleFunc("/ping", ping)

	log.Printf("kpw7-stats-display on port %s\n", port)
	if err := http.ListenAndServe(":42799", nil); err != nil {
		log.Println("http error:", err)
	}
}
