package main

import (
	"github.com/henry0475/go-viewer-agent/src/options"
	"github.com/henry0475/go-viewer-agent/src/register"
	"github.com/henry0475/go-viewer-agent/src/reporter"
	"github.com/henry0475/go-viewer-agent/src/tools/ids"
	"github.com/henry0475/go-viewer-agent/src/tracker"
)

var opt *options.TrackerOption

// InitAgent is the init func for establishing the connections to the server
func InitAgent(opts ...*options.TrackerOption) {
	opt = options.MergeTrackerOptions(opts...)
	if opt.NodeID == nil {
		opt.NodeID = options.Int64(ids.GetNodeID())
	}

	if opt.GRPCOption != nil && reporter.GetRepoter() == nil {
		register.NewRegister(
			reporter.NewReporter(opt.GRPCOption),
		)
	}
}

// NewTracker should be used at very top level
func NewTracker(remarks string) *tracker.Tracker {
	if remarks == "" {
		remarks = "Entrance"
	}

	return tracker.NewTracker(
		reporter.GetRepoter(),
		*opt.NodeID, remarks,
	)
}
