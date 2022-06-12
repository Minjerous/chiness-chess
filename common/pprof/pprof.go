package pprof

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func InitPprof() {
	go func() {
		log.Println(http.ListenAndServe(":9990", nil))
	}()
}
