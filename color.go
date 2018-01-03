package main

import (
	"net/http"
	"regexp"
	"log"
)

var triplets = regexp.MustCompile(`^(?i)([0-9A-F]{3})|([0-9A-F]{6})$`)

func handler(w http.ResponseWriter, r *http.Request) {
	color := r.URL.Path[1:]
	if triplets.Match([]byte(color)) {
		color = "#" + color
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	w.Write([]byte(`
		<svg width="16" height="16" viewBox="0 0 16 16" xmlns="http://www.w3.org/2000/svg">
			<rect stroke="#000" stroke-width="2" fill="` + color + `" x="1" y="1" width="14" height="14" rx="0" ry="0" />
		</svg>`))
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
