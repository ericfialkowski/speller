package main

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	wordMap := map[string]bool{}
	file, err := os.Open("words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		file.Close()
	}()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		wordMap[strings.ToLower(scanner.Text())] = true
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		word := strings.ToLower(r.RequestURI)[1:] // strip the leading /
		if len(word) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		found := wordMap[word]
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(found)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting: %v", err)
	}
}
