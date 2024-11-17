package metadata

import (
	"go.opentelemetry.io/collector/component"
)

var (
	Type      = component.MustNewType("clickhouse")
	ScopeName = "github.com/zeelrupapara/clickhousemetrics"
)

const (
	MetricsStability = component.StabilityLevelAlpha
	TracesStability  = component.StabilityLevelBeta
	LogsStability    = component.StabilityLevelBeta
)