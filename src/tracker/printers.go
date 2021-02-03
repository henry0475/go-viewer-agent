package tracker

import (
	"log"
)

// PrintTimeStamp prints ...
func (t *Tracker) PrintTimeStamp() {
	if v, ok := t.ctx.Value(TrackCalling).(*TrackingInfo); ok {
		for _, trackInfo := range v.Tracks {
			log.Println(trackInfo.Bucket.Hash, trackInfo.Timestamp)
		}
	}
}

// PrintUUID prints ....
func (t *Tracker) PrintUUID() {
	log.Println(t.uuid)
}

// PrintTracker is ....
func (t *Tracker) PrintTracker() {
	if v, ok := t.ctx.Value(TrackCalling).(*TrackingInfo); ok {
		for i, trackInfo := range v.Tracks {
			log.Printf(`
			Tracker Depth: %d
			Tracker UUID: %d
			Total Depth: %d
			Bucket Hash: %s
			Bucket Name: %s
			File: %s:%d
			FuncName: %s
			Timestamp: %d
			Duration Nano: %d
			Duration Millisecond: %d
			Remarks: %s
			=====================
			`, i, t.GetUUID(), t.GetDepth(), trackInfo.Bucket.Hash, trackInfo.Bucket.Name, trackInfo.FilePath, trackInfo.Line, trackInfo.FuncName, trackInfo.Timestamp, trackInfo.Duration, trackInfo.Duration/1e6, trackInfo.Remarks)
		}
	}
}
