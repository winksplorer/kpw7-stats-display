package main

import (
	"log"
	"net/http"
)

var frontendDir string = "./frontend"
var port string = ":42799"

func main() {
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
