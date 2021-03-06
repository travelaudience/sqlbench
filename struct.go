package sqlbench

import (
	"sync"
)

// Config specifies the setup for benchmarks
type Config struct {
	Tags    []Tag   `json:"tags"`
	Queries []Query `json:"queries"`
	// PostgreSQL database DSN
	Db   string `json:"db"`
	Logs struct {
		// If set it will append the logs to this csv file
		Csv string `json:"csv"`
		// If set it will send the results to datadog
		Datadog string `json:"datadog"`
	} `json:"logs"`
}

// Tag defines tags which shows the condition which benchmark was running, like number of rows in table(s), database size.
type Tag struct {
	// Name can be any string
	Name string `json:"name"`
	// Value can be only a query which will result in a numeric value. Alternatively it can be `epoch` or `datetime`
	Value string `json:"value"`
}

// Query defines parameters to benchmark a query
type Query struct {
	// A name for the query
	Name string `json:"name"`
	// Running frequency in millisecond
	Frequency int `json:"frequency"`
	// Number of parallel runs for this query
	Parallel int `json:"parallel"`
	// Number if runs for this query
	Count int `json:"count"`
	// Query to run
	Query string `json:"query"`
}

type runner interface {
	run(string) error
	tag(string) (string, error)
	init() error
}

// Bench is the structure which encapsulated the required methods for running benchmark.
type Bench struct {
	config Config
	wait   chan bool
	runner
	runLog
	sync.Mutex
}

// Stats is collections data we collect for each query run
type Stats struct {
	// Minimum runtime
	Min, Avg, Max, Stdv, Pct95 float64
}

// Log of total exection and also queries benchmarks
type runLog struct {
	tags []Tag
	runs map[string]Stats
}
