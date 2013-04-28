package thecarcle

import (
	"log"
	"net/http"
	"runtime"
)

func NewApp() (r *TheCarcleApp) {
	return &TheCarcleApp{}
}

type TheCarcleApp struct {
}

func (this *TheCarcleApp) Run() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	log.Println("Starting server on %s\n", HttpPort)
	m := Mux()
	err := http.ListenAndServe(HttpPort, m)
	if err != nil {
		panic("Http ListenAndServe: " + err.Error())
	}
}
