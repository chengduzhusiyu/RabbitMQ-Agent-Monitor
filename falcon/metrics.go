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
	Endpoint    string      `json:"endpoint"`
	Metric      string      `json:"metric"`
	Value       interface{} `json:"value"`
	CounterType string      `json:"counterType"`
	Tags        string      `json:"tags"`
	Timestamp   int64       `json:"timestamp"`
	Step        int64       `json:"step"`
}

// NewMetric create an new metric
func NewMetric(name string, value interface{}, tags string) *MetaData {
	host := g.GetHost()
	return &MetaData{
		Metric:      name,
		Endpoint:    host,
		CounterType: fmt.Sprintf("GAUGE"),
		Tags:        tags,
		Timestamp:   time.Now().Unix(),
		Step:        g.Config().Interval,
		Value:       value,
	}
}

func (m *MetaData) String() string {
	s := fmt.Sprintf("MetaData Metric:%s Endpoint:%s Value:%v CounterType:%s Tags:%s Timestamp:%d Step:%d",
		m.Metric, m.Endpoint, m.Value, m.CounterType, m.Tags, m.Timestamp, m.Step)
	return s
}

// SetValue value setter
func (m *MetaData) SetValue(v interface{}) {
	m.Value = v
}

func trimFloat(s float64) float64 {
	if s, err := strconv.ParseFloat(fmt.Sprintf("%.3f", s), 64); err == nil {
		return s
	}
	return s
}

func calcPercentage(l, t int64) (pct float64) {
	if t == 0 {
		return
	}
	pct = float64(l) / float64(t) * 100.00
	pct = trimFloat(pct)
	return
}

func qStats(s string) int64 {
	var aliveQueue = g.Config().Qrunning
	for _, i := range aliveQueue {
		if strings.Contains(strings.ToLower(s), i) {
			return 1
		}
	}
	return 0
}

func isAliveness(s string) int64 {
	switch s {
	case "ok":
		return 1
	default:
		return 0
	}
}

func partitions(s []string) int64 {
	switch len(s) {
	case 0:
		return 1
	default:
		return 0
	}
}

func consumerUtil(c interface{}) float64 {
	if vv, ok := c.(float64); ok {
		return trimFloat(vv * 100.00)
	} else if _, ok := c.(bool); ok {
		return 0.0
	} else if _, ok := c.(string); ok {
		return 0.0
	}
	return 0.0
}

func updateCurrentStatsDB(db string) {
	statsDB.SetCurrentLocate(db)
}

// GetCurrentStatsDB get current stats management database
func GetCurrentStatsDB() *g.StatsDB {
	return statsDB
}

// handleJudge
func handleJudge() (data []*MetaData) {
	data = make([]*MetaData, 0)
	nd, err := funcs.GetNode()
	if err != nil {
		log.Println(err.Error())
		return
	}

	data = append(data, NewMetric(overviewPrefix+"ioReadawait", nd.Rawait, ""))    // io_read_avg_wait_time
	data = append(data, NewMetric(overviewPrefix+"ioWriteawait", nd.Wawait, ""))   // io_write_avg_wait_time
	data = append(data, NewMetric(overviewPrefix+"ioSyncawait", nd.Syncawait, "")) // io_sync_avg_wait_time
	data = append(data, NewMetric(overviewPrefix+"memConnreader", nd.ConnectionReaders, ""))
	data = append(data, NewMetric(overviewPrefix+"memConnwriter", nd.ConnectionWriters, ""))
	data = append(data, NewMetric(overviewPrefix+"memConnchannels