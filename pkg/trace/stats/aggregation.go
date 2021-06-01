package stats

import (
	"strconv"
	"strings"

	"github.com/DataDog/datadog-agent/pkg/trace/pb"
	"github.com/DataDog/datadog-agent/pkg/trace/traceutil"
	"github.com/DataDog/datadog-agent/pkg/util/log"
)

const (
	tagHostname   = "_dd.hostname"
	tagStatusCode = "http.status_code"
	tagVersion    = "version"
	tagOrigin     = "_dd.origin"
	tagSynthetics = "synthetics"
)

// Aggregation contains all the dimension on which we aggregate statistics
// when adding or removing fields to Aggregation the methods ToTagSet, KeyLen,
// WriteKey and the structs payloadAggregationKey, bucketAggregationKey in the ClientStatsAggregator
// should always be updated accordingly.
type Aggregation struct {
	Env        string
	Resource   string
	Service    string
	Name       string
	Type       string
	Hostname   string
	EntityID   string
	StatusCode uint32
	Version    string
	Synthetics bool
}

func getStatusCode(s *pb.Span) uint32 {
	strC := traceutil.GetMetaDefault(s, tagStatusCode, "")
	if strC == "" {
		return 0
	}
	c, err := strconv.Atoi(strC)
	if err != nil {
		log.Debugf("Invalid status code %s. Using 0.", strC)
		return 0
	}
	return uint32(c)
}

// NewAggregationFromSpan creates a new aggregation from the provided span and env
func NewAggregationFromSpan(s *pb.Span, env string, agentHostname string, isServerless bool) Aggregation {
	synthetics := strings.HasPrefix(traceutil.GetMetaDefault(s, tagOrigin, ""), tagSynthetics)
	var hostname string
	if !isServerless {
		hostname = traceutil.GetMetaDefault(s, tagHostname, "")
		if hostname == "" {
			hostname = agentHostname
		}
	}
	entityID := getSpanEntity(s, isServerless)

	return Aggregation{
		Env:        env,
		Resource:   s.Resource,
		Service:    s.Service,
		Name:       s.Name,
		Type:       s.Type,
		Hostname:   hostname,
		EntityID:   entityID,
		StatusCode: getStatusCode(s),
		Version:    traceutil.GetMetaDefault(s, tagVersion, ""),
		Synthetics: synthetics,
	}
}

// NewAggregationFromGroup gets the Aggregation key of grouped stats.
func NewAggregationFromGroup(env, hostname, version string, g pb.ClientGroupedStats) Aggregation {
	return Aggregation{
		Env:        env,
		Hostname:   hostname,
		Version:    version,
		Resource:   g.Resource,
		Service:    g.Service,
		Name:       g.Name,
		StatusCode: g.HTTPStatusCode,
		Synthetics: g.Synthetics,
	}
}

func getSpanEntity(s *pb.Span, isServerless bool) string {
	if isServerless {
		// Try to find the Fargate task_arn in the span Metadata
		if containerTagsStr, ok := s.Meta["_dd.tags.container"]; ok {
			for _, t := range strings.Split(containerTagsStr, ",") {
				if k, v := splitTag(t); k == "task_arn" {
					return "task_arn://" + v
				}
			}
		}
	}
	return ""
}

func splitTag(tag string) (string, string) {
	if i := strings.Index(tag, ":"); i >= 0 {
		return tag[:i], tag[i+1:]
	}
	return tag, ""
}
