package server

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/D-Zane/img-generator/configs"
	"github.com/D-Zane/img-generator/pkg/img"
)

func render(w http.ResponseWriter, msg string) {
	_, err := w.Write([]byte(msg))
	if err != nil {
		log.Println(err)
	}
}

func renderImg(w http.ResponseWriter, buffer *bytes.Buffer) {
	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println(err)
	}
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {

	buffer, err := img.GenerateFavicon()
	if err != nil {
		log.Println(err)
	}

	renderImg(w, buffer)
}

func imgHandler(w http.ResponseWriter, r *http.Request) {
	buffer, err := img.Generate(strings.Split(r.URL.Path, "/"))
	if err != nil {
		log.Println(err)
	}
	renderImg(w, buffer)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "ping")
}

func robotsHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "robots")
}

func Run(conf configs.ConfI) {
	http.HandleFunc("/", imgHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/robots.txt", robotsHandler)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Println("Server starting...")
		if err := http.ListenAndServe(":"+conf.GetPort(), nil); err != nil {
			log.Fatalln(err)
		}
	}()

	signalValue := <-sigs
	signal.Stop(sigs)

	log.Println("stop signal: ", signalValue)
}
