package cron

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/barryz/rmqmonitor/g"
	"github.com/barryz/rmqmonitor/utils"

	"github.com/barryz/cron"
)

func logRotateRun() {
	log.Println("[INFO] start to rotate rabbitmq log file ......")
	suffix := fmt.Sprintf(".%s", utils.GetYesterdayDate())
	logRotateComman