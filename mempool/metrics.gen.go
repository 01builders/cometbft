// Code generated by metricsgen. DO NOT EDIT.

package mempool

import (
	"github.com/go-kit/kit/metrics/discard"
	prometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func PrometheusMetrics(namespace string, labelsAndValues ...string) *Metrics {
	labels := []string{}
	for i := 0; i < len(labelsAndValues); i += 2 {
		labels = append(labels, labelsAndValues[i])
	}
	return &Metrics{
		Size: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "size",
			Help:      "Number of uncommitted transactions in the mempool.",
		}, labels).With(labelsAndValues...),
		SizeBytes: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "size_bytes",
			Help:      "Total size of the mempool in bytes.",
		}, labels).With(labelsAndValues...),
		TxSizeBytes: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "tx_size_bytes",
			Help:      "Histogram of transaction sizes in bytes.",

			Buckets: stdprometheus.ExponentialBuckets(1, 3, 7),
		}, labels).With(labelsAndValues...),
		FailedTxs: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "failed_txs",
			Help:      "FailedTxs defines the number of failed transactions. These are transactions that failed to make it into the mempool because they were deemed invalid. metrics:Number of failed transactions.",
		}, labels).With(labelsAndValues...),
		RejectedTxs: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "rejected_txs",
			Help:      "RejectedTxs defines the number of rejected transactions. These are transactions that failed to make it into the mempool due to resource limits, e.g. mempool is full. metrics:Number of rejected transactions.",
		}, labels).With(labelsAndValues...),
		EvictedTxs: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "evicted_txs",
			Help:      "EvictedTxs defines the number of evicted transactions. These are valid transactions that passed CheckTx and make it into the mempool but later became invalid. metrics:Number of evicted transactions.",
		}, labels).With(labelsAndValues...),
		RecheckTimes: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "recheck_times",
			Help:      "Number of times transactions are rechecked in the mempool.",
		}, labels).With(labelsAndValues...),
		ActiveOutboundConnections: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "active_outbound_connections",
			Help:      "Number of connections being actively used for gossiping transactions (experimental feature).",
		}, labels).With(labelsAndValues...),
		ExpiredTxs: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "expired_txs",
			Help:      "ExpiredTxs defines transactions that were removed from the mempool due to a TTL",
		}, labels).With(labelsAndValues...),
		SuccessfulTxs: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "successful_txs",
			Help:      "SuccessfulTxs defines the number of transactions that successfully made it into a block.",
		}, labels).With(labelsAndValues...),
		AlreadySeenTxs: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "already_seen_txs",
			Help:      "AlreadySeenTxs defines the number of transactions that entered the mempool which were already present in the mempool. This is a good indicator of the degree of duplication in message gossiping.",
		}, labels).With(labelsAndValues...),
		RequestedTxs: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "requested_txs",
			Help:      "RequestedTxs defines the number of times that the node requested a tx to a peer",
		}, labels).With(labelsAndValues...),
		RerequestedTxs: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "rerequested_txs",
			Help:      "RerequestedTxs defines the number of times that a requested tx never received a response in time and a new request was made.",
		}, labels).With(labelsAndValues...),
	}
}

func NopMetrics() *Metrics {
	return &Metrics{
		Size:                      discard.NewGauge(),
		SizeBytes:                 discard.NewGauge(),
		TxSizeBytes:               discard.NewHistogram(),
		FailedTxs:                 discard.NewCounter(),
		RejectedTxs:               discard.NewCounter(),
		EvictedTxs:                discard.NewCounter(),
		RecheckTimes:              discard.NewCounter(),
		ActiveOutboundConnections: discard.NewGauge(),
		ExpiredTxs:                discard.NewCounter(),
		SuccessfulTxs:             discard.NewCounter(),
		AlreadySeenTxs:            discard.NewCounter(),
		RequestedTxs:              discard.NewCounter(),
		RerequestedTxs:            discard.NewCounter(),
	}
}
