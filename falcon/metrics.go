package falcon

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/barryz/rmqmonitor/funcs"
	"github.com/barryz/rmqmonitor/g"
)

var (
	statsDB = g.NewStatsDB()
)

const (
	overviewPrefix = "rabbitmq.overview."
	queuePrefix    = "rabbitmq.queue."
	exchangePrefix = "rabbitmq.exchange."
)

// MetaData meta data
type MetaData struct {
	Endpoint    string      `