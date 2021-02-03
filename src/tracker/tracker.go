package tracker

import (
	"context"
	"errors"
	"runtime"
	"strings"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/henry0475/go-viewer-agent/src/reporter"
	"github.com/henry0475/go-viewer-agent/src/tools/hash"
)

// Trackable is ...
type Trackable interface{}

// Tracker is the interface for ...
type Tracker struct {
	ctx context.Context

	reporter *reporter.Reporter

	nodeID int64
	uuid   int64

	depth int
}

// NewTracker is the constructor
func NewTracker(r *reporter.Reporter, nodeID int64, remarks string) *Tracker {
	var t = new(Tracker)
	t.reporter = r

	pc, file, line, _ := runtime.Caller(1)
	fileArr := strings.Split(file, "/")

	var trackInfo = new(TrackingInfo)
	var i = &info{
		FilePath:  file,
		Line:      line,
		Remarks:   remarks,
		FuncName:  runtime.FuncForPC(pc).Name(),
		Timestamp: time.Now().UnixNano(),
		Duration:  0,
	}
	i.Bucket.Name = fileArr[len(fileArr)-2]
	i.Bucket.Hash = hash.ToSha1(strings.Join(fileArr[:len(fileArr)-1], "/"))
	trackInfo.Tracks = append(trackInfo.Tracks, i)
	t.depth = 1
	t.nodeID = nodeID

	node, _ := snowflake.NewNode(t.nodeID)
	t.uuid = node.Generate().Int64()

	t.ctx = context.WithValue(context.Background(), TrackCalling, trackInfo)

	return t
}

// Track should be used frequently after the instance of `Entrance` has been initialized.
func Track(tracker *Tracker, remarks string) (*Tracker, error) {
	if v, ok := tracker.ctx.Value(TrackCalling).(*TrackingInfo); ok {
		pc, file, line, _ := runtime.Caller(1)

		fileArr := strings.Split(file, "/")
		var i = &info{
			FilePath:  file,
			FuncName:  runtime.FuncForPC(pc).Name(),
			Line:      line,
			Remarks:   remarks,
			Timestamp: time.Now().UnixNano(),
		}
		i.Duration = i.Timestamp - v.Tracks[len(v.Tracks)-1].Timestamp
		i.Bucket.Name = fileArr[len(fileArr)-2]
		i.Bucket.Hash = hash.ToSha1(strings.Join(fileArr[:len(fileArr)-1], "/"))
		v.Tracks = append(v.Tracks, i)

		tracker.ctx = context.WithValue(tracker.ctx, TrackCalling, v)
		tracker.depth = tracker.depth + 1

		return tracker, nil
	}

	return nil, errors.New("Parent context is not valid")
}

// EndTrack should be used as long as you want to end the tracking.
func EndTrack(tracker *Tracker) (*Tracker, error) {

	if v, ok := tracker.ctx.Value(TrackCalling).(*TrackingInfo); ok {
		pc, file, line, _ := runtime.Caller(1)

		fileArr := strings.Split(file, "/")
		var i = &info{
			FilePath:  file,
			FuncName:  runtime.FuncForPC(pc).Name(),
			Line:      line,
			Remarks:   "EndTrack",
			Timestamp: time.Now().UnixNano(),
		}
		i.Duration = i.Timestamp - v.Tracks[len(v.Tracks)-1].Timestamp
		i.Bucket.Name = fileArr[len(fileArr)-2]
		i.Bucket.Hash = hash.ToSha1(strings.Join(fileArr[:len(fileArr)-1], "/"))
		v.Tracks = append(v.Tracks, i)

		tracker.ctx = context.WithValue(tracker.ctx, TrackCalling, v)
		tracker.depth = tracker.depth + 1

		tracker.Report()
		return tracker, nil
	}

	return nil, errors.New("Parent context is not valid")
}
