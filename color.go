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
		<svg width="256" height="256" viewBox="0 0 256 256" xmlns="http://www.w3.org/2000/svg">
			<rect stroke="#000" stroke-width="4" fill="` + color + `" x="2" y="2" width="252" height="252" rx="0" ry="0" />
		</svg>`
	))
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
