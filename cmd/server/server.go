package main

import (
	"flag"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/fagongzi/log"
)

var (
	addr = flag.String("addr", "127.0.0.1:9092", "Addr: client entrypoint")
)

func main() {
	flag.Parse()

	log.InitLog()
	runtime.GOMAXPROCS(runtime.NumCPU())

	waitStop()
}

func waitStop() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	sig := <-sc
	log.Infof("exit: signal=<%d>.", sig)
	switch sig {
	case syscall.SIGTERM:
		log.Infof("exit: bye :-).")
		os.Exit(0)
	default:
		log.Infof("exit: bye :-(.")
		os.Exit(1)
	}
}
