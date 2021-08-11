package handlers

import (
	"github.com/aleperaltabazas/dtp/cli"
	"os"
	"os/signal"
	"syscall"
)

func Init() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cli.Exit()
	}()
}
