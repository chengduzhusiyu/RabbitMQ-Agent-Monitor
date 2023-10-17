package witch

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/barryz/rmqmonitor/g"
	"github.com/barryz/rmqmonitor/witch/system"
)

const (
	commandBuildIn    = "buildin"
	commandSupervisor = "supervisor"
	commandSystemd    = "systemd"
)

func handleSignals(exitFunc func()) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SI