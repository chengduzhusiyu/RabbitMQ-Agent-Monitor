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

// New