package main

import (
	"time"
	"flag"
	//"os"
	//"runtime/pprof"
	_ "net/http/pprof"
	"net/http"

	"github.com/getamis/sirius/log"

)



func main() {
	//profilePath := flag.String("profilePath", "cpu.prof", "path")
	local := flag.Int("delay", 1000, "millisecond")
	delay := flag.Int("delay", 1000, "millisecond")
	flag.Parse()

	//f, err := os.Create(*profilePath)
	//if err != nil {
	//	log.Crit("could not create CPU profile", "err", err)
	//}
	//if err := pprof.StartCPUProfile(f); err != nil {
	//	log.Crit("could not start CPU profile: ", "err", err)
	//}
	//defer pprof.StopCPUProfile()


	go func() {http.ListenAndServe("0.0.0.0:8000", nil)}()

	ticker := time.NewTicker(time.Duration(*delay) *time.Millisecond)
	logger := log.New()
	for {
		for t := range ticker.C {
			logger.Info("Print a message", "time", t.String())
		}
	}
}
