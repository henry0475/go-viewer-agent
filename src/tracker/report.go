package tracker

import (
	"context"
	"runtime"

	collector "github.com/henry0475/go-viewer-agent/src/reporter/protos"
)

// Report will send the info to the server
func (t *Tracker) Report() (err error) {
	if v, ok := t.ctx.Value(TrackCalling).(*TrackingInfo); ok {
		var tracks = make([]*collector.TrackedInfo, 0, t.depth)

		for i := 0; i < t.depth && t.reporter != nil; i++ {
			trackInfo := v.Tracks[i]

			tracks = append(tracks, &collector.TrackedInfo{
				Bucket: &collector.BucketInfo{
					Name: trackInfo.Bucket.Name,
					Hash: trackInfo.Bucket.Hash,
				},
				FilePath:  trackInfo.FilePath,
				FuncName:  trackInfo.FuncName,
				Line:      int64(trackInfo.Line),
				Remarks:   trackInfo.Remarks,
				Timestamp: trackInfo.Timestamp,
				Duration:  trackInfo.Duration,
			})
		}

		if len(tracks) != 0 {
			t.reporter.Collector.RecordTracker(context.Background(), &collector.RecordTrackerRequest{
				Agent: &collector.AgentInfo{
					NodeID: t.nodeID,
					NumCPU: int64(runtime.NumCPU()),
				},
				UUID:  t.GetUUID(),
				Depth: int64(t.GetDepth()),
				Info:  tracks,
			})
		}
	}

	return
}
