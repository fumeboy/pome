package trace

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	MaxTraceId = 100000000
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenTraceId() (traceId string) {
	now := time.Now()
	traceId = fmt.Sprintf("%04d%02d%02d%02d%02d%02d%08d", now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second(), rand.Int31n(MaxTraceId))
	return
}
