package main

import (
	"testing"
	"time"

	"github.com/henry0475/go-viewer-agent/src/options"
	"github.com/henry0475/go-viewer-agent/src/tracker"
)

func one(track *tracker.Tracker) {
	track, _ = tracker.Track(track, "one")
	two(track)
}

func two(track *tracker.Tracker) {
	track, _ = tracker.Track(track, "two")
	time.Sleep(time.Second * time.Duration(3))
	three(track)
}

func three(track *tracker.Tracker) {
	track, _ = tracker.Track(track, "three")
	end, _ := tracker.EndTrack(track)
	end.PrintTracker()
}

func TestMain(t *testing.T) {
	InitAgent(&options.TrackerOption{
		GRPCOption: &options.GRPCConfigs{
			Address: "127.0.0.1",
			Port:    3000,
		},
	})
	track := NewTracker("main")
	one(track)

	t.Errorf("Done")
}
