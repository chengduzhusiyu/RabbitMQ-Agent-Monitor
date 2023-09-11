package system

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strconv"
	"syscall"
	"time"
)

// Launcher supervises the process status, start, stop and restart.
type Launcher struct {
	pidFile string
	cmd     string
}

// NewLauncher creates new system.
func NewLauncher(pidFile, cmd string) *Launcher {
	return &Launcher{
		pidFile: pidFile,
		cmd:     cmd,
	}
}

func (s *Launcher) writePid(pid int) {
	if err := WriteFile(s.pidFile, []byte(strconv.FormatInt(int64(pid), 10)), 0644); err != nil {
		log.Fatalf("Failed to write pid file: %s", err)
	}
}

func (s *Launcher) readPid() (int, bool) {
	f, err := ioutil.ReadFile(s.pidFile)
	if err !