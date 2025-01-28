// Code generated by metricsgen. DO NOT EDIT.

package consensus

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
		Height: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "height",
			Help:      "Height of the chain.",
		}, labels).With(labelsAndValues...),
		ValidatorLastSignedHeight: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "validator_last_signed_height",
			Help:      "Last height signed by this validator if the node is a validator.",
		}, append(labels, "validator_address")).With(labelsAndValues...),
		Rounds: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "rounds",
			Help:      "Number of rounds.",
		}, labels).With(labelsAndValues...),
		RoundDurationSeconds: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "round_duration_seconds",
			Help:      "Histogram of round duration.",

			Buckets: stdprometheus.ExponentialBucketsRange(0.1, 100, 8),
		}, labels).With(labelsAndValues...),
		Validators: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "validators",
			Help:      "Number of validators.",
		}, labels).With(labelsAndValues...),
		ValidatorsPower: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "validators_power",
			Help:      "Total power of all validators.",
		}, labels).With(labelsAndValues...),
		ValidatorPower: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "validator_power",
			Help:      "Power of a validator.",
		}, append(labels, "validator_address")).With(labelsAndValues...),
		ValidatorMissedBlocks: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "validator_missed_blocks",
			Help:      "Amount of blocks missed per validator.",
		}, append(labels, "validator_address")).With(labelsAndValues...),
		MissingValidators: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "missing_validators",
			Help:      "Number of validators who did not sign.",
		}, labels).With(labelsAndValues...),
		MissingValidatorsPower: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "missing_validators_power",
			Help:      "Total power of the missing validators.",
		}, labels).With(labelsAndValues...),
		ByzantineValidators: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "byzantine_validators",
			Help:      "Number of validators who tried to double sign.",
		}, labels).With(labelsAndValues...),
		ByzantineValidatorsPower: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "byzantine_validators_power",
			Help:      "Total power of the byzantine validators.",
		}, labels).With(labelsAndValues...),
		BlockIntervalSeconds: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "block_interval_seconds",
			Help:      "Time between this and the last block.",
		}, labels).With(labelsAndValues...),
		NumTxs: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "num_txs",
			Help:      "Number of transactions.",
		}, labels).With(labelsAndValues...),
		BlockSizeBytes: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "block_size_bytes",
			Help:      "Size of the block.",
		}, labels).With(labelsAndValues...),
		ChainSizeBytes: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "chain_size_bytes",
			Help:      "Size of the chain in bytes.",
		}, labels).With(labelsAndValues...),
		TotalTxs: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "total_txs",
			Help:      "Total number of transactions.",
		}, labels).With(labelsAndValues...),
		CommittedHeight: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "latest_block_height",
			Help:      "The latest block height.",
		}, labels).With(labelsAndValues...),
		BlockParts: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "block_parts",
			Help:      "Number of block parts transmitted by each peer.",
		}, append(labels, "peer_id")).With(labelsAndValues...),
		DuplicateBlockPart: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "duplicate_block_part",
			Help:      "Number of times we received a duplicate block part",
		}, labels).With(labelsAndValues...),
		DuplicateVote: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "duplicate_vote",
			Help:      "Number of times we received a duplicate vote",
		}, labels).With(labelsAndValues...),
		StepDurationSeconds: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "step_duration_seconds",
			Help:      "Histogram of durations for each step in the consensus protocol.",

			Buckets: stdprometheus.ExponentialBucketsRange(0.1, 100, 8),
		}, append(labels, "step")).With(labelsAndValues...),
		BlockGossipPartsReceived: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "block_gossip_parts_received",
			Help:      "Number of block parts received by the node, separated by whether the part was relevant to the block the node is trying to gather or not.",
		}, append(labels, "matches_current")).With(labelsAndValues...),
		QuorumPrevoteDelay: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "quorum_prevote_delay",
			Help:      "Interval in seconds between the proposal timestamp and the timestamp of the earliest prevote that achieved a quorum.",
		}, append(labels, "proposer_address")).With(labelsAndValues...),
		FullPrevoteDelay: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "full_prevote_delay",
			Help:      "Interval in seconds between the proposal timestamp and the timestamp of the latest prevote in a round where all validators voted.",
		}, append(labels, "proposer_address")).With(labelsAndValues...),
		VoteExtensionReceiveCount: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "vote_extension_receive_count",
			Help:      "VoteExtensionReceiveCount is the number of vote extensions received by this node. The metric is annotated by the status of the vote extension from the application, either 'accepted' or 'rejected'.",
		}, append(labels, "status")).With(labelsAndValues...),
		ProposalReceiveCount: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "proposal_receive_count",
			Help:      "ProposalReceiveCount is the total number of proposals received by this node since process start. The metric is annotated by the status of the proposal from the application, either 'accepted' or 'rejected'.",
		}, append(labels, "status")).With(labelsAndValues...),
		ProposalCreateCount: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "proposal_create_count",
			Help:      "ProposalCreationCount is the total number of proposals created by this node since process start. The metric is annotated by the status of the proposal from the application, either 'accepted' or 'rejected'.",
		}, labels).With(labelsAndValues...),
		RoundVotingPowerPercent: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "round_voting_power_percent",
			Help:      "RoundVotingPowerPercent is the percentage of the total voting power received with a round. The value begins at 0 for each round and approaches 1.0 as additional voting power is observed. The metric is labeled by vote type.",
		}, append(labels, "vote_type")).With(labelsAndValues...),
		LateVotes: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "late_votes",
			Help:      "LateVotes stores the number of votes that were received by this node that correspond to earlier heights and rounds than this node is currently in.",
		}, append(labels, "vote_type")).With(labelsAndValues...),
		ProposalTimestampDifference: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "proposal_timestamp_difference",
			Help:      "Difference in seconds between the local time when a proposal message is received and the timestamp in the proposal message.",

			Buckets: []float64{-1.5, -1.0, -0.5, -0.2, 0, 0.2, 0.5, 1.0, 1.5, 2.0, 2.5, 4.0, 8.0},
		}, append(labels, "is_timely")).With(labelsAndValues...),
		StartHeight: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "start_height",
			Help:      "StartHeight is the height at which metrics began.",
		}, labels).With(labelsAndValues...),
		BlockTimeSeconds: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "block_time_seconds",
			Help:      "BlockTimeSeconds is the duration between this block and the preceding one.",
		}, labels).With(labelsAndValues...),
		ApplicationRejectedProposals: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "application_rejected_proposals",
			Help:      "ApplicationRejectedProposals is the number of proposals rejected by the application.",
		}, labels).With(labelsAndValues...),
		TimedOutProposals: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "timed_out_proposals",
			Help:      "TimedOutProposals is the number of proposals that failed to be received in time.",
		}, labels).With(labelsAndValues...),
	}
}

func NopMetrics() *Metrics {
	return &Metrics{
		Height:                       discard.NewGauge(),
		ValidatorLastSignedHeight:    discard.NewGauge(),
		Rounds:                       discard.NewGauge(),
		RoundDurationSeconds:         discard.NewHistogram(),
		Validators:                   discard.NewGauge(),
		ValidatorsPower:              discard.NewGauge(),
		ValidatorPower:               discard.NewGauge(),
		ValidatorMissedBlocks:        discard.NewGauge(),
		MissingValidators:            discard.NewGauge(),
		MissingValidatorsPower:       discard.NewGauge(),
		ByzantineValidators:          discard.NewGauge(),
		ByzantineValidatorsPower:     discard.NewGauge(),
		BlockIntervalSeconds:         discard.NewHistogram(),
		NumTxs:                       discard.NewGauge(),
		BlockSizeBytes:               discard.NewGauge(),
		ChainSizeBytes:               discard.NewCounter(),
		TotalTxs:                     discard.NewGauge(),
		CommittedHeight:              discard.NewGauge(),
		BlockParts:                   discard.NewCounter(),
		DuplicateBlockPart:           discard.NewCounter(),
		DuplicateVote:                discard.NewCounter(),
		StepDurationSeconds:          discard.NewHistogram(),
		BlockGossipPartsReceived:     discard.NewCounter(),
		QuorumPrevoteDelay:           discard.NewGauge(),
		FullPrevoteDelay:             discard.NewGauge(),
		VoteExtensionReceiveCount:    discard.NewCounter(),
		ProposalReceiveCount:         discard.NewCounter(),
		ProposalCreateCount:          discard.NewCounter(),
		RoundVotingPowerPercent:      discard.NewGauge(),
		LateVotes:                    discard.NewCounter(),
		ProposalTimestampDifference:  discard.NewHistogram(),
		StartHeight:                  discard.NewGauge(),
		BlockTimeSeconds:             discard.NewGauge(),
		ApplicationRejectedProposals: discard.NewCounter(),
		TimedOutProposals:            discard.NewCounter(),
	}
}
