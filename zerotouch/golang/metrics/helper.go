package metrics

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

type Helper struct {
	SERVICE_NAME string
	DATABASE string
}

func (m *Helper) PushToSummarytMetrics() func(*prometheus.SummaryVec,string,*error,context.Context) {
	startTime := time.Now()
	return func(request *prometheus.SummaryVec,methodName string,err *error,ctx context.Context) {

		if *err != nil {
			request.WithLabelValues(m.SERVICE_NAME,methodName,"ko").Observe(float64(time.Now().Sub(startTime).Milliseconds()))
		} else {
			request.WithLabelValues(m.SERVICE_NAME,methodName,"ok").Observe(float64(time.Now().Sub(startTime).Milliseconds()))
		}
	}
}

func (m *Helper) PushToErrorCounterMetrics() func(*prometheus.CounterVec,error,context.Context) {
	return func(request *prometheus.CounterVec,err error,ctx context.Context) {
		request.WithLabelValues(m.SERVICE_NAME,m.DATABASE,err.Error()).Inc()
	}
}
