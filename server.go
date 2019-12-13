package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var jsonCommands = [...]string{
	"termux-battery-status",
	"termux-camera-info",
	"termux-contact-list",
	"termux-infrared-frequencies",
	"termux-location",
	"termux-sms-inbox",
	"termux-telephony-cellinfo",
	"termux-telephony-deviceinfo",
	"termux-tts-engines",
}

func json(w http.ResponseWriter, r *http.Request) {
	var arg []string
	cmd := chi.URLParam(r, "cmd")
	if !strings.HasPrefix(cmd, "termux-") {
		cmd = "termux-" + cmd
	}
	cmdFound := false
	for _, c := range jsonCommands {
		if cmd == c {
			cmdFound = true
			break
		}
	}
	if !cmdFound {
		http.Error(w, "not supported", 404)
		return
	}

	if cmd == "termux-location" {
		p := r.FormValue("p")
		if p == "network" || p == "passive" {
			arg = append(arg, "-p", p)
		}
	}
	if cmd == "termux-sms-inbox" {
		if dates := r.FormValue("d"); dates == "true" {
			arg = append(arg, "-d")
		}
		if limit := r.FormValue("l"); limit != "" {
			arg = append(arg, "-l", limit)
		}
		if numbers := r.FormValue("n"); numbers == "true" {
			arg = append(arg, "-n")
		}
		if offset := r.FormValue("o"); offset != "" {
			arg = append(arg, "-o", offset)
		}
	}

	out, err := exec.Command(cmd, arg...).Output()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(out)
}

func camera(w http.ResponseWriter, r *http.Request) {
	photo := "termux-photo.jpg"
	cameraID := r.FormValue("c")
	if cameraID == "" {
		cameraID = "0"
	}

	out, err := exec.Command("termux-camera-photo",
		"-c", cameraID, photo).Output()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if string(out) != "" {
		http.Error(w, string(out), 500)
		return
	}

	jpg, err := ioutil.ReadFile(photo)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "image/jpeg")
	w.Write(jpg)
}

func notification(w http.ResponseWriter, r *http.Request) {
	var arg []string

	if content := r.FormValue("c"); content != "" {
		arg = append(arg, "-c", content)
	}
	if id := r.FormValue("i"); id != "" {
		arg = append(arg, "-i", id)
	}
	if title := r.FormValue("t"); title != "" {
		arg = append(arg, "-t", title)
	}
	if url := r.FormValue("u"); url != "" {
		arg = append(arg, "-u", url)
	}

	out, err := exec.Command("termux-notification", arg...).Output()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if string(out) != "" {
		http.Error(w, string(out), 500)
		return
	}
	w.Write([]byte("OK"))
}

func sendSMS(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("t")
	r.ParseForm()
	var numbers string
	for _, number := range r.Form["n"] {
		numbers = numbers + number + ","
	}

	out, err := exec.Command("termux-sms-send", "-n", numbers, text).Output()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if string(out) != "" {
		http.Error(w, string(out), 500)
		return
	}
	w.Write([]byte("OK"))
}

func speak(w http.ResponseWriter, r *http.Request) {
	var arg []string

	if engine := r.FormValue("e"); engine != "" {
		arg = append(arg, "-e", engine)
	}
	if language := r.FormValue("l"); language != "" {
		arg = append(arg, "-l", language)
	}
	if pitch := r.FormValue("p"); pitch != "" {
		arg = append(arg, "-p", pitch)
	}
	if rate := r.FormValue("r"); rate != "" {
		arg = append(arg, "-r", rate)
	}
	if stream := r.FormValue("s"); stream != "" {
		arg = append(arg, "-s", stream)
	}
	if text := r.FormValue("t"); text != "" {
		arg = append(arg, text)
	}

	out, err := exec.Command("termux-tts-speak", arg...).Output()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if string(out) != "" {
		http.Error(w, string(out), 500)
		return
	}
	w.Write([]byte("OK"))
}

func vibrate(w http.ResponseWriter, r *http.Request) {
	var arg []string

	if duration := r.FormValue("d"); duration != "" {
		arg = append(arg, "-d", duration)
	}
	if force := r.FormValue("d"); force == "true" {
		arg = append(arg, "-f")
	}

	out, err := exec.Command("termux-vibrate", arg...).Output()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if string(out) != "" {
		http.Error(w, string(out), 500)
		return
	}
	w.Write([]byte("OK"))
}

func main() {
	var port string
	flag.StringVar(&port, "p", "8000", "port")
	flag.Parse()

	r := chi.NewRouter()
	r.Use(cors)
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(index))
	})
	r.Get("/camera-photo", camera)
	r.Get("/notification", notification)
	r.Get("/sms-send", sendSMS)
	r.Get("/tts-speak", speak)
	r.Get("/vibrate", vibrate)
	r.Get("/:cmd", json)

	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, r))
}
