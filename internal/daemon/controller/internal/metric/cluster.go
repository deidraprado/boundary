package metric

import (
	"context"
	"net"

	"github.com/hashicorp/boundary/globals"
	"github.com/hashicorp/boundary/internal/daemon/metric"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/stats"
)

const (
	clusterSubsystem = "controller_cluster"
)

// Because we use grpc's stats.Handler to track close-to-the-wire server-side communication over grpc,
// there is no easy way to measure request and response size as we are recording latency. Thus we only
// track the request latency for server-side grpc connections.

// grpcRequestLatency collects measurements of how long it takes
// the boundary system to reply to a request to the controller cluster
// from the time that boundary received the request.
var grpcRequestLatency prometheus.ObserverVec = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Namespace: globals.MetricNamespace,
		Subsystem: clusterSubsystem,
		Name:      "grpc_request_duration_seconds",
		Help:      "Histogram of latencies for gRPC requests.",
		Buckets:   prometheus.DefBuckets,
	},
	metric.ListGrpcLabels,
)

// acceptedConnsTotal keeps a count of the total accepted connections to a controller.
var acceptedConnsTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Namespace: globals.MetricNamespace,
		Subsystem: clusterSubsystem,
		Name:      "accepted_connections_total",
		Help:      "Count of total accepted network connections to this controller.",
	},
	[]string{metric.LabelConnectionType},
)

// closedConnsTotal keeps a count of the total closed connections to a controller.
var closedConnsTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Namespace: globals.MetricNamespace,
		Subsystem: clusterSubsystem,
		Name:      "closed_connections_total",
		Help:      "Count of total closed network connections to this controller.",
	},
	[]string{metric.LabelConnectionType},
)

// All the codes expected to be returned by boundary or the grpc framework to
// requests to the cluster server.
var expectedGrpcCodes = []codes.Code{
	codes.OK, codes.InvalidArgument, codes.PermissionDenied,
	codes.FailedPrecondition,

	// Codes which can be generated by the gRPC framework
	codes.Canceled, codes.Unknown, codes.DeadlineExceeded,
	codes.ResourceExhausted, codes.Unimplemented, codes.Internal,
	codes.Unavailable, codes.Unauthenticated,
}

func InstrumentClusterTrackingListener(l net.Listener, purpose string) net.Listener {
	p := prometheus.Labels{metric.LabelConnectionType: purpose}
	return metric.NewConnectionTrackingListener(l, acceptedConnsTotal.With(p), closedConnsTotal.With(p))
}

// InstrumentClusterStatsHandler returns a gRPC stats.Handler which observes
// cluster specific metrics. Use with the cluster gRPC server.
func InstrumentClusterStatsHandler(ctx context.Context) (stats.Handler, error) {
	return metric.NewStatsHandler(ctx, grpcRequestLatency)
}

// InitializeClusterCollectors registers the cluster metrics to the default
// prometheus register and initializes them to 0 for all possible label
// combinations.
func InitializeClusterCollectors(r prometheus.Registerer, server *grpc.Server) {
	metric.InitializeGrpcCollectorsFromServer(r, grpcRequestLatency, server, expectedGrpcCodes)
}

func InitializeConnectionCounters(r prometheus.Registerer) {
	metric.InitializeConnectionCounters(r, []prometheus.CounterVec{*acceptedConnsTotal, *closedConnsTotal})
}
