package main

import (
	"fmt"
	"net/http"
	"encoding/json"

)

type Header struct {
	IpAddress string
	Language string
	Software string
}
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the header parser api!\n" +
		"From your browser visit the url 0.0.0.0:8080/api/whoami\n" +
		"You should receive your Header information returned as JSON.")
}

func WhoAmI(w http.ResponseWriter, r *http.Request) {
	h := r.Header

	ip := h.Get("X-Real-Ip")
	if ip == "" {
		ip = h.Get("X-Forward-For")
	}
	if ip == "" {
		ip = r.RemoteAddr
	}

	lang := h.Get("Accept-Language")
	software := h.Get("User-Agent")

	header := Header {
		IpAddress: ip,
		Language: lang,
		Software: software,
	}

	json,err := json.MarshalIndent(&header,"","\t\t")
	if err != nil {
		return
	}
	w.Write(json)
	return

	fmt.Fprintln(w, header)
}

