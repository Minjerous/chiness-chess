package pprof

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

//  pprof
func InitPprof() {
	go func() {
		log.Println(http.ListenAndServe(":", nil))
		log.Println(http.ListenAndServe(":", nil))
		log.Println(http.ListenAndServe(":", nil))
		log.Println(http.ListenAndServe(":", nil))
		log.Println(http.ListenAndServe(":", nil))
	}()
}
